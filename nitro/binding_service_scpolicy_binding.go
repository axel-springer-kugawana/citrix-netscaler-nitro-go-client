package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type ServiceScpolicyBinding struct {
	Name       string `json:"name,omitempty"`
	Policyname string `json:"policyname,omitempty"`
}

type ServiceScpolicyBindingKey struct {
	Name       string
	Policyname string
}

type get_service_scpolicy_binding struct {
	Results []ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

type add_service_scpolicy_binding_payload struct {
	service_scpolicy_binding ServiceScpolicyBinding
}

func service_scpolicy_binding_key_to_id_args(key ServiceScpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddServiceScpolicyBinding(binding ServiceScpolicyBinding) error {
	payload := add_service_scpolicy_binding_payload{
		binding,
	}

	return c.put("service_scpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListServiceScpolicyBinding(key ServiceScpolicyBindingKey) ([]ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding

	id, args := service_scpolicy_binding_key_to_id_args(key)

	if err := c.get("service_scpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListServiceScpolicyBinding() ([]ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding

	if err := c.get("service_scpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceScpolicyBinding(key ServiceScpolicyBindingKey) (*ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding

	id, args := service_scpolicy_binding_key_to_id_args(key)

	if err := c.get("service_scpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service_scpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("service_scpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServiceScpolicyBinding(key ServiceScpolicyBindingKey) error {
	id, args := service_scpolicy_binding_key_to_id_args(key)

	return c.delete("service_scpolicy_binding", id, "", args)
}
