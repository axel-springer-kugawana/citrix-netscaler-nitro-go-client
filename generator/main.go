package main

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/spec"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func is_array(t string) bool {
	return strings.HasSuffix(t, "[]")
}

func is_enum(t string) bool {
	return strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")")
}

func go_base_type(t string) string {
	t = strings.TrimSuffix(t, "[]")

	isEnum := is_enum(t)

	t = strings.TrimPrefix(t, "(")
	t = strings.TrimSuffix(t, ")")

	base := "string"

	switch t {
	case "ip", "ip_mask", "string":
		base = "string"
	case "double", "int":
		base = "int"
	case "bool":
		base = "bool"
	}

	if isEnum {
		base = "string"
	}

	return base
}

func go_type(t string) string {
	isArray := is_array(t)

	base := go_base_type(t)

	if isArray {
		base = "[]" + base
	}

	return base
}

func main() {
	spec, err := nitro.ReadSpec("C:/Users/ce.breteche/go/src/github.com/doubret/citrix-netscaler-nitro-yaml-specs/yml")

	if err != nil {
		log.Println("Failed to read spec : ", err)
	} else {
		funcMap := template.FuncMap{
			"title":   strings.Title,
			"go_type": go_type,
		}

		templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/resource.tmpl", "templates/binding.tmpl"))

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
		}

		for key, _ := range spec.Bindings {
			log.Print("Binding : " + key)
		}
	}
}
