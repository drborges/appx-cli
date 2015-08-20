package appxcli

import (
	"log"
	"sync"
)

func Process(pkg string) {
	var wg sync.WaitGroup
	for _, ast := range Parse(pkg) {
		for model := range Traverse(pkg, ast) {
			wg.Add(1)
			go func(m *ModelDescriptor) {
				if err := Generate(m); err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(model)
		}
	}

	wg.Wait()
}
