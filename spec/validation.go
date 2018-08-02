package nitro

import (
	"fmt"
	"strings"
)

func validateResources(resources map[string]*Resource) error {
	for key, resource := range resources {
		if resource.Key == nil || resource.Key.Name == "" {
			return fmt.Errorf("Invalid resource spec, no key name defined : %v", key)
		}

		if resource.Key.Fields != nil {
			for _, field := range resource.Key.Fields {
				_, ok := resource.Fields[field]
				if !ok {
					return fmt.Errorf("Invalid resource spec, key field unknown : %v.%v", key, field)
				}
			}
		}

		if resource.Update != nil {
			for _, field := range resource.Update {
				_, ok := resource.Fields[field]
				if !ok && field != resource.State {
					return fmt.Errorf("Invalid resource spec, update field unknown : %v.%v", key, field)
				}
			}
		}

		for field, value := range resource.Fields {
			ok := false

			if strings.HasSuffix(value, "[]") {
				value = strings.TrimSuffix(value, "[]")
			}

			if value == "ip" || value == "ip_mask"  || value == "int"  || value == "string"  || value == "bool"  || value == "double" {
				ok = true
			} else if strings.HasPrefix(value, "(") && strings.HasSuffix(value, ")")  {
				ok = true
			} else if strings.Contains(value, ".") {
				parts := strings.Split(value, ".")

				if len(parts) == 2 {
					target, found := resources[parts[0]]

					if found && parts[1] != target.Key.Name && parts[1] != target.State  {
						_, found = target.Fields[parts[1]]
					}

					if found {
						ok = true
					}
				}
			}

			if !ok {
				return fmt.Errorf("Invalid resource spec, invalid field type : %v.%v (%v)", key, field, value)
			}
		}
	}

	return nil
}

func validateBindings(resources map[string]*Resource, bindings map[string]*Binding) error {
	return nil
}

func validateSpec(resources map[string]*Resource, bindings map[string]*Binding) error {
	err := validateResources(resources)

	if err != nil {
		return err
	}

	err = validateBindings(resources, bindings)

	if err != nil {
		return err
	}

	return nil
}
