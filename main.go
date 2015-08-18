package main

import (
	"github.com/drborges/go-ast/appx"
	"log"
	"sync"
)

func main() {
	pkg := "models"
	var wg sync.WaitGroup
	for _, ast := range appx.Parse(pkg) {
		for model := range appx.Traverse(pkg, ast) {
			wg.Add(1)
			go func(m *appx.ModelDescriptor) {
				if err := appx.Generate(m); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(model)
		}
	}

	wg.Wait()
}
