package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type ServiceDospolicyBinding struct {
	Name       string `json:"name,omitempty"`
	Policyname string `json:"policyname,omitempty"`
}

type ServiceDospolicyBindingKey struct {
	Name       string
	Policyname string
}

type get_service_dospolicy_binding struct {
	Results []ServiceDospolicyBinding `json:"service_dospolicy_binding"`
}

type add_service_dospolicy_binding_payload struct {
	service_dospolicy_binding ServiceDospolicyBinding
}

func service_dospolicy_binding_key_to_id_args(key ServiceDospolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddServiceDospolicyBinding(binding ServiceDospolicyBinding) error {
	payload := add_service_dospolicy_binding_payload{
		binding,
	}

	return c.put("service_dospolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListServiceDospolicyBinding(key ServiceDospolicyBindingKey) ([]ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding

	id, args := service_dospolicy_binding_key_to_id_args(key)

	if err := c.get("service_dospolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListServiceDospolicyBinding() ([]ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding

	if err := c.get("service_dospolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceDospolicyBinding(key ServiceDospolicyBindingKey) (*ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding

	id, args := service_dospolicy_binding_key_to_id_args(key)

	if err := c.get("service_dospolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service_dospolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("service_dospolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServiceDospolicyBinding(key ServiceDospolicyBindingKey) error {
	id, args := service_dospolicy_binding_key_to_id_args(key)

	return c.delete("service_dospolicy_binding", id, "", args)
}
