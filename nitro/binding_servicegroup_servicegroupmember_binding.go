package nitro

import (
	"fmt"
	"strconv"
)

type ServicegroupServicegroupmemberBinding struct {
	Port             int    `json:"port,omitempty"`
	Servername       string `json:"servername,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupServicegroupmemberBindingKey struct {
	Servicegroupname string
	Servername       string
	Port             int
}

type get_servicegroup_servicegroupmember_binding struct {
	Results []ServicegroupServicegroupmemberBinding `json:"servicegroup_servicegroupmember_binding"`
}

type add_servicegroup_servicegroupmember_binding_payload struct {
	servicegroup_servicegroupmember_binding ServicegroupServicegroupmemberBinding
}

func servicegroup_servicegroupmember_binding_key_to_id_args(key ServicegroupServicegroupmemberBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",servicegroupname:" + key.Servicegroupname
	result = result + ",servername:" + key.Servername
	result = result + ",port:" + strconv.Itoa(key.Port)
	return "", result
}

func (c *NitroClient) AddServicegroupServicegroupmemberBinding(binding ServicegroupServicegroupmemberBinding) error {
	payload := add_servicegroup_servicegroupmember_binding_payload{
		binding,
	}

	return c.put("servicegroup_servicegroupmember_binding", "", "", "", payload)
}

func (c *NitroClient) ListServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) ([]ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding

	id, args := servicegroup_servicegroupmember_binding_key_to_id_args(key)

	if err := c.get("servicegroup_servicegroupmember_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListServicegroupServicegroupmemberBinding() ([]ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding

	if err := c.get("servicegroup_servicegroupmember_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (*ServicegroupServicegroupmemberBinding, error) {
	var results get_servicegroup_servicegroupmember_binding

	id, args := servicegroup_servicegroupmember_binding_key_to_id_args(key)

	if err := c.get("servicegroup_servicegroupmember_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup_servicegroupmember_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("servicegroup_servicegroupmember_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) error {
	id, args := servicegroup_servicegroupmember_binding_key_to_id_args(key)

	return c.delete("servicegroup_servicegroupmember_binding", id, "", args)
}
