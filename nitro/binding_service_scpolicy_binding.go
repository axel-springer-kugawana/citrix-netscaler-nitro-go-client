package nitro

import (
	"fmt"
	"net/url"
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

func (resource ServiceScpolicyBinding) ToKey() ServiceScpolicyBindingKey {
	key := ServiceScpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key ServiceScpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)
	args = append(args, "policyname:"+url.QueryEscape(key.Policyname))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ServiceScpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServiceScpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_service_scpolicy_binding_payload struct {
	Resource ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

func (c *NitroClient) AddServiceScpolicyBinding(resource ServiceScpolicyBinding) error {
	payload := add_service_scpolicy_binding_payload{
		resource,
	}

	return c.put("service_scpolicy_binding", "", nil, payload)
}

//      LIST

type list_service_scpolicy_binding_result struct {
	Results []ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

func (c *NitroClient) ListServiceScpolicyBinding() ([]ServiceScpolicyBinding, error) {
	results := list_service_scpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("service_scpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_service_scpolicy_binding_result struct {
	Results []Count `json:"service_scpolicy_binding"`
}

func (c *NitroClient) CountServiceScpolicyBinding() (int, error) {
	results := count_service_scpolicy_binding_result{}

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

//      READ

type get_service_scpolicy_binding_result struct {
	Results []ServiceScpolicyBinding `json:"service_scpolicy_binding"`
}

func (c *NitroClient) GetServiceScpolicyBinding(key ServiceScpolicyBindingKey) (*ServiceScpolicyBinding, error) {
	var results get_service_scpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("service_scpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service_scpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("service_scpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsServiceScpolicyBinding(key ServiceScpolicyBindingKey) (bool, error) {
	var results count_service_scpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("service_scpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServiceScpolicyBinding(key ServiceScpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("service_scpolicy_binding", id, qs)
}
