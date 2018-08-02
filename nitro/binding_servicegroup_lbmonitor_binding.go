package nitro

import (
	"fmt"
	"strconv"
)

type ServicegroupLbmonitorBinding struct {
	Monitor_name     string `json:"monitor_name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupLbmonitorBindingKey struct {
	Servicegroupname string
	Monitor_name     string
}

type get_servicegroup_lbmonitor_binding struct {
	Results []ServicegroupLbmonitorBinding `json:"servicegroup_lbmonitor_binding"`
}

type add_servicegroup_lbmonitor_binding_payload struct {
	servicegroup_lbmonitor_binding ServicegroupLbmonitorBinding
}

func servicegroup_lbmonitor_binding_key_to_id_args(key ServicegroupLbmonitorBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",servicegroupname:" + key.Servicegroupname
	result = result + ",monitor_name:" + key.Monitor_name
	return "", result
}

func (c *NitroClient) AddServicegroupLbmonitorBinding(binding ServicegroupLbmonitorBinding) error {
	payload := add_servicegroup_lbmonitor_binding_payload{
		binding,
	}

	return c.put("servicegroup_lbmonitor_binding", "", "", "", payload)
}

func (c *NitroClient) ListServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) ([]ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding

	id, args := servicegroup_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("servicegroup_lbmonitor_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListServicegroupLbmonitorBinding() ([]ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding

	if err := c.get("servicegroup_lbmonitor_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (*ServicegroupLbmonitorBinding, error) {
	var results get_servicegroup_lbmonitor_binding

	id, args := servicegroup_lbmonitor_binding_key_to_id_args(key)

	if err := c.get("servicegroup_lbmonitor_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_lbmonitor_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("servicegroup_lbmonitor_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) error {
	id, args := servicegroup_lbmonitor_binding_key_to_id_args(key)

	return c.delete("servicegroup_lbmonitor_binding", id, "", args)
}
