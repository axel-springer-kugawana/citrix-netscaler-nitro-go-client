package main

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/spec"
	"log"
)

func main() {
	spec, err := nitro.ReadSpec("C:/Users/ce.breteche/go/src/github.com/doubret/citrix-netscaler-nitro-yaml-specs/yml")

	if err != nil {
		log.Println("Failed to read spec : ", err)
	} else {
		for key, _ := range spec.Resources {
			log.Print("Resource : " + key)
		}
		for key, _ := range spec.Bindings {
			log.Print("Binding : " + key)
		}
	}
}
