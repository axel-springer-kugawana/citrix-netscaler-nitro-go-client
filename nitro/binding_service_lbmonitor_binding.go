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

type get_service_lbmonitor_binding struct {
	Results []ServiceLbmonitorBinding `json:"service_lbmonitor_binding"`
}

type add_service_lbmonitor_binding_payload struct {
	service_lbmonitor_binding ServiceLbmonitorBinding
}

func service_lbmonitor_binding_key_to_id_args(key ServiceLbmonitorBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddServiceLbmonitorBinding(binding ServiceLbmonitorBinding) error {
	payload := add_service_lbmonitor_binding_payload{
		binding,
	}

	return c.put("service_lbmonitor_binding", "", "", "", payload)
}

func (c *NitroClient) ListServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) ([]ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding

	id, args := service_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("service_lbmonitor_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListServiceLbmonitorBinding() ([]ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding

	if err := c.get("service_lbmonitor_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) (*ServiceLbmonitorBinding, error) {
	var results get_service_lbmonitor_binding

	id, args := service_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("service_lbmonitor_binding", id, "", args, &results); err != nil {
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
	id, args := service_lbmonitor_binding_key_to_id_args(key)

	return c.delete("service_lbmonitor_binding", id, "", args)
}
