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

type add_service_scpolicy_binding_payload struct {
	Resources ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

type get_service_scpolicy_binding_result struct {
	Results []ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

type count_service_scpolicy_binding_result struct {
	Results []Count `json:"service_scpolicy_binding"`
}

func service_scpolicy_binding_key_to_id_args(key ServiceScpolicyBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddServiceScpolicyBinding(binding ServiceScpolicyBinding) error {
	payload := add_service_scpolicy_binding_payload{
		binding,
	}

	return c.put("service_scpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountServiceScpolicyBinding() (int, error) {
	var results count_service_scpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("service_scpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountServiceScpolicyBinding(id string) (int, error) {
	var results count_service_scpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("service_scpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServiceScpolicyBinding(id string) (bool, error) {
	if count, err := c.CountServiceScpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListServiceScpolicyBinding() ([]ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("service_scpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListServiceScpolicyBinding(id string) ([]ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding_result

	if err := c.get("service_scpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceScpolicyBinding(key ServiceScpolicyBindingKey) (*ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding_result

	id, qs := service_scpolicy_binding_key_to_id_args(key)

	if err := c.get("service_scpolicy_binding", id, qs, &results); err != nil {
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
	id, qs := service_scpolicy_binding_key_to_id_args(key)

	return c.delete("service_scpolicy_binding", id, qs)
}
