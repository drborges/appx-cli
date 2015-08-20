package appxcli

import (
	"github.com/drborges/go-ast/appxcli/templates"
	"log"
	"os"
	"text/template"
	"strings"
)

func GenerateApi(model *ModelDescriptor) error {
	t := template.Must(template.New("model").Parse(templates.AppxModel))
	output := model.Pkg + "/appx_model_" + model.Name + ".go"
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := t.Execute(f, model); err != nil {
		return err
	}

	t = template.Must(template.New("repo").Funcs(templates.FuncMap).Parse(templates.AppxRepository))
	output = model.Pkg + "/appx_model_" + model.Name + "_repository.go"
	f, err = os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := t.Execute(f, model); err != nil {
		return err
	}

	log.Println("[API] generated for model", model.Name)
	return nil
}

func GenerateModel(outPath string, model *Model) error {
	t := template.Must(template.New("entity").Funcs(templates.FuncMap).Parse(templates.AppxEntity))
	output := outPath + "/" + strings.ToLower(model.Name) + ".go"

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, model)
}