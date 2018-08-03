package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type ServiceLbmonitorBinding struct {
	Monitor_name string `json:"monitor_name,omitempty"`
	Monstate     string `json:"monstate,omitempty"`
	Name         string `json:"name,omitempty"`
	Passive      bool   `json:"passive,omitempty"`
	Weight       int    `json:"weight,string,omitempty"`
}

type ServiceLbmonitorBindingKey struct {
	Name string
}

type add_service_lbmonitor_binding_payload struct {
	Resource ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

type get_service_lbmonitor_binding_result struct {
	Results []ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

type count_service_lbmonitor_binding_result struct {
	Results []Count `json:"service_lbmonitor_binding"`
}

func service_lbmonitor_binding_key_to_id_args(key ServiceLbmonitorBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddServiceLbmonitorBinding(binding ServiceLbmonitorBinding) error {
	payload := add_service_lbmonitor_binding_payload{
		binding,
	}

	return c.put("service_lbmonitor_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountServiceLbmonitorBinding() (int, error) {
	var results count_service_lbmonitor_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("service_lbmonitor_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountServiceLbmonitorBinding(id string) (int, error) {
	var results count_service_lbmonitor_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("service_lbmonitor_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServiceLbmonitorBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountServiceLbmonitorBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListServiceLbmonitorBinding() ([]ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("service_lbmonitor_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListServiceLbmonitorBinding(id string) ([]ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding_result

	if err := c.get("service_lbmonitor_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) (*ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding_result

	id, qs := service_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("service_lbmonitor_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service_lbmonitor_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("service_lbmonitor_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) error {
	id, qs := service_lbmonitor_binding_key_to_id_args(key)

	return c.delete("service_lbmonitor_binding", id, qs)
}
