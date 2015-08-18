package appx

import (
	"go/ast"
	"log"
	"strings"
)

func fieldType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident: // Trivial
		return t.Name
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + fieldType(t.X)
	case *ast.ArrayType:
		// TODO handle slices of primitive types
		// TODO handle embedded types
		log.Println("Ignoring array field type:", expr)
		return ""//"[]" + fieldType(t.Elt)
	case *ast.MapType:
		log.Println("Ignoring map field type:", expr)
		return ""//"map[" + fieldType(t.Key) + "]" + fieldType(t.Value)
	default:
		log.Println("Ignoring unkownfield type:", expr)
		return ""
	}
}

func Traverse(pkg string, f *ast.File) <-chan *ModelDescriptor {
	ch := make(chan *ModelDescriptor, 10)

	go func() {
		defer close(ch)
		for _, decl := range f.Decls {
			switch gendecl := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range gendecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok {
						continue
					}

					if len(structType.Fields.List) == 0 {
						continue
					}

					isAppxModel := false
					fields := make(map[string]string)
					for _, field := range structType.Fields.List {
						if len(field.Names) > 0 {
							if ftype := fieldType(field.Type); ftype != "" {
								fields[field.Names[0].Name] = ftype
							}
						}

						// A struct definition is an appx model if it has a field 'f' where:
						//
						// 1. f.Names is empty
						// 2. f.Type is *ast.SelectorExpr
						// 3. f.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name == "appx"
						// 4. f.Type.(*ast.SelectorExpr).Sel.Name == "Model"
						if expr, ok := field.Type.(*ast.SelectorExpr); ok {
							if expr.X.(*ast.Ident).Name == "appx" || expr.Sel.Name == "Model" {
								isAppxModel = true
							}
						}
					}

					if isAppxModel {
						modelName := typeSpec.Name.Name
						ch <- &ModelDescriptor{
							Fields: fields,
							Pkg:    pkg,
							Name:   modelName,
							Path:   pkg + "/appx_model_" + strings.ToLower(modelName) + ".go",
						}
					}
				}
			}
		}
	}()

	return ch
}
