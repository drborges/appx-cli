package appxcli

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func Parse(pkg string) map[string]*ast.File {
	excludeTestFiles := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go")
	}

	fset := token.NewFileSet() // positions are relative to fset
	asts, err := parser.ParseDir(fset, pkg, excludeTestFiles, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	return asts[pkg].Files
}
