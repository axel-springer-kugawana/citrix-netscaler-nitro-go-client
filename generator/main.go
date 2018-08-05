package main

import (
	"github.com/doubret/citrix-netscaler-nitro-go-specs/specs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	spec, err := nitro.ReadSpec("C:/Users/ce.breteche/go/src/github.com/doubret/citrix-netscaler-nitro-yaml-specs/yml")

	if err != nil {
		log.Println("Failed to read spec : ", err)
	} else {
		funcMap := template.FuncMap{
			"title":   strings.Title,
			"go_type": nitro.GoType,
			"name":    nitro.Name,
			"is_in":   nitro.IsIn,
		}

		templates := template.Must(template.New("").Funcs(funcMap).ParseFiles(
			"templates/resource.tmpl", "templates/resource_test.tmpl", "templates/resource_factory.tmpl",
			"templates/binding.tmpl",
		))

		for key, value := range spec.Resources {
			context := struct {
				Name   string
				Schema *nitro.Resource
			}{
				key,
				value,
			}

			writer, err := os.Create(filepath.Join("nitro", "resource_"+key+".go"))

			if err != nil {
				log.Println("Failed to create file : ", err)
			}

			err = templates.ExecuteTemplate(writer, "resource.tmpl", context)

			if err != nil {
				log.Println("Failed to execute template : ", err)
			}

			if _, err := os.Stat(filepath.Join("tests", "resources", key+"_factory.go")); os.IsNotExist(err) {
				//{
				writer, err := os.Create(filepath.Join("tests", "resources", key+"_factory.go"))

				if err != nil {
					log.Println("Failed to create file : ", err)
				}

				err = templates.ExecuteTemplate(writer, "resource_factory.tmpl", context)

				if err != nil {
					log.Println("Failed to execute template : ", err)
				}
			}

			writer, err = os.Create(filepath.Join("tests", "resources", key+"_test.go"))

			if err != nil {
				log.Println("Failed to create file : ", err)
			}

			err = templates.ExecuteTemplate(writer, "resource_test.tmpl", context)

			if err != nil {
				log.Println("Failed to execute template : ", err)
			}
		}

		for key, value := range spec.Bindings {
			context := struct {
				Name   string
				Schema *nitro.Binding
			}{
				key,
				value,
			}

			writer, err := os.Create(filepath.Join("nitro", "binding_"+key+".go"))

			if err != nil {
				log.Println("Failed to create file : ", err)
			}

			err = templates.ExecuteTemplate(writer, "binding.tmpl", context)

			if err != nil {
				log.Println("Failed to execute template : ", err)
			}
		}
	}
}
