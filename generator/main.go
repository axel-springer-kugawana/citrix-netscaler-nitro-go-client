package main

import (
	"flag"
	"github.com/doubret/citrix-netscaler-nitro-go-specs/specs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var inputFolder = flag.String("i", "", "input folder to read specs from")

func generateResource(templates *template.Template, name string, resource *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		resource,
	}

	writer, err := os.Create(filepath.Join("nitro", "resource_"+name+".go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "resource.tmpl", context)

	return err
}

func generateBinding(templates *template.Template, name string, binding *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		binding,
	}

	writer, err := os.Create(filepath.Join("nitro", "binding_"+name+".go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "binding.tmpl", context)

	return err
}

func generateResourceTestFactory(templates *template.Template, name string, resource *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		resource,
	}

	writer, err := os.Create(filepath.Join("tests", "resources", name+"_factory.go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "test_factory.tmpl", context)

	return err
}

func generateBindingTestFactory(templates *template.Template, name string, binding *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		binding,
	}

	writer, err := os.Create(filepath.Join("tests", "bindings", name+"_factory.go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "test_factory.tmpl", context)

	return err
}

func generateResourceTest(templates *template.Template, name string, resource *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		resource,
	}

	writer, err := os.Create(filepath.Join("tests", "resources", name+"_test.go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "resource_test.tmpl", context)

	return err
}

func generateBindingTest(templates *template.Template, name string, binding *nitro.SpecFile) error {
	context := struct {
		Name   string
		Schema *nitro.SpecFile
	}{
		name,
		binding,
	}

	writer, err := os.Create(filepath.Join("tests", "bindings", name+"_test.go"))

	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(writer, "binding_test.tmpl", context)

	return err
}

func main() {
	flag.Parse()

	if inputFolder == nil {
		log.Print("No input spec folder specified")

		return
	} else {
		spec, err := nitro.ReadSpec(*inputFolder)

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
				"templates/resource.tmpl",
				"templates/binding.tmpl",
				"templates/resource_test.tmpl", "templates/binding_test.tmpl", "templates/test_factory.tmpl",
				"templates/utils.tmpl",
			))

			for key, value := range spec.Resources {
				err := generateResource(templates, key, value)

				if err != nil {
					log.Println("Failed to generate resource : ", err)
				}

				err = generateResourceTest(templates, key, value)

				if err != nil {
					log.Println("Failed to generate resource test : ", err)
				}

				if _, err := os.Stat(filepath.Join("tests", "resources", key+"_factory.go")); os.IsNotExist(err) {
					err := generateResourceTestFactory(templates, key, value)

					if err != nil {
						log.Println("Failed to generate resource test factory : ", err)
					}
				}
			}

			for key, value := range spec.Bindings {
				err := generateBinding(templates, key, value)

				if err != nil {
					log.Println("Failed to generate binding : ", err)
				}

				err = generateBindingTest(templates, key, value)

				if err != nil {
					log.Println("Failed to generate binding test : ", err)
				}

				if _, err := os.Stat(filepath.Join("tests", "bindings", key+"_factory.go")); os.IsNotExist(err) {
					err := generateBindingTestFactory(templates, key, value)

					if err != nil {
						log.Println("Failed to generate binding test factory : ", err)
					}
				}
			}
		}
	}
}
