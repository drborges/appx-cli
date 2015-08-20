package appxcli

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"log"
)

func Parse(path string) []*ast.File {
	excludeTestFiles := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go")
	}

	fset := token.NewFileSet() // positions are relative to fset
	asts, err := parser.ParseDir(fset, path, excludeTestFiles, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	files := []*ast.File{}
	for _, pkgAsts := range asts {
		for _, pkgFile := range pkgAsts.Files {
			files = append(files, pkgFile)
		}
	}

	return files
}
