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

type add_service_dospolicy_binding_payload struct {
	Resource ServiceDospolicyBinding `json:"service_dospolicy_binding"`
}

type get_service_dospolicy_binding_result struct {
	Results []ServiceDospolicyBinding `json:"service_dospolicy_binding"`
}

type count_service_dospolicy_binding_result struct {
	Results []Count `json:"service_dospolicy_binding"`
}

func service_dospolicy_binding_key_to_id_qs(key ServiceDospolicyBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func service_dospolicy_binding_key_to_id_args(key ServiceDospolicyBindingKey) (string, map[string]string) {
	return service_dospolicy_binding_key_to_id_qs(key, "args")
}

func service_dospolicy_binding_key_to_id_filter(key ServiceDospolicyBindingKey) (string, map[string]string) {
	return service_dospolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddServiceDospolicyBinding(binding ServiceDospolicyBinding) error {
	payload := add_service_dospolicy_binding_payload{
		binding,
	}

	return c.put("service_dospolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountServiceDospolicyBinding() (int, error) {
	var results count_service_dospolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("service_dospolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountServiceDospolicyBinding(id string) (int, error) {
	var results count_service_dospolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("service_dospolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServiceDospolicyBinding(key ServiceDospolicyBindingKey) (bool, error) {
	var results count_service_dospolicy_binding_result

	id, qs := service_dospolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("service_dospolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one service_dospolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListServiceDospolicyBinding() ([]ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("service_dospolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListServiceDospolicyBinding(id string) ([]ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding_result

	if err := c.get("service_dospolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceDospolicyBinding(key ServiceDospolicyBindingKey) (*ServiceDospolicyBinding, error) {
	var results get_service_dospolicy_binding_result

	id, qs := service_dospolicy_binding_key_to_id_args(key)

	if err := c.get("service_dospolicy_binding", id, qs, &results); err != nil {
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
	id, qs := service_dospolicy_binding_key_to_id_args(key)

	return c.delete("service_dospolicy_binding", id, qs)
}
