package appxcli

import (
	"github.com/drborges/go-ast/appxcli/templates"
	"log"
	"os"
	"text/template"
)

func Generate(model *ModelDescriptor) error {
	log.Println("Generating:", model.Name)

	t := template.Must(template.New("model").Parse(templates.AppxModel))
	output := model.Pkg + "/appx_model_" + model.Name + ".go"
	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Could not create file %v: %v", output, err)
	}
	defer f.Close()

	if err := t.Execute(f, model); err != nil {
		log.Fatal(err)
	}

	t = template.Must(template.New("repo").Funcs(templates.FuncMap).Parse(templates.AppxRepository))
	output = model.Pkg + "/appx_model_" + model.Name + "_repository.go"
	f, err = os.Create(output)
	if err != nil {
		log.Fatalf("Could not create file %v: %v", output, err)
	}
	defer f.Close()

	if err := t.Execute(f, model); err != nil {
		log.Fatal(err)
	}

	return nil
}
