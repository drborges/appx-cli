package appxcli

import (
	"log"
	"sync"
)

func Process(path string) {
	var wg sync.WaitGroup
	generated := false

	for _, ast := range Parse(path) {
		for model := range Traverse(path, ast) {
			wg.Add(1)
			go func(m *ModelDescriptor) {
				if err := GenerateApi(m); err != nil {
					log.Fatal(err)
				}
				wg.Done()
				generated = true
			}(model)
		}
	}

	wg.Wait()

	if !generated {
		log.Fatalln("No Appx model found at", path)
	}
}
