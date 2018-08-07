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

func (resource ServiceLbmonitorBinding) ToKey() ServiceLbmonitorBindingKey {
	key := ServiceLbmonitorBindingKey{
		resource.Name,
	}

	return key
}

func (key ServiceLbmonitorBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_service_lbmonitor_binding_payload struct {
	Resource ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

func (c *NitroClient) AddServiceLbmonitorBinding(resource ServiceLbmonitorBinding) error {
	payload := add_service_lbmonitor_binding_payload{
		resource,
	}

	return c.put("service_lbmonitor_binding", "", nil, payload)
}

//      LIST

type list_service_lbmonitor_binding_result struct {
	Results []ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

func (c *NitroClient) ListServiceLbmonitorBinding() ([]ServiceLbmonitorBinding, error) {
	results := list_service_lbmonitor_binding_result{}

	if err := c.get("service_lbmonitor_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_service_lbmonitor_binding_result struct {
	Results []ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

func (c *NitroClient) GetServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) (*ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding_result

	id, qs := key.to_id_args()

	if err := c.get("service_lbmonitor_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service_lbmonitor_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("service_lbmonitor_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_service_lbmonitor_binding_result struct {
	Results []Count `json:"service_lbmonitor_binding"`
}

func (c *NitroClient) ExistsServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) (bool, error) {
	var results count_service_lbmonitor_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("service_lbmonitor_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("service_lbmonitor_binding", id, qs)
}
