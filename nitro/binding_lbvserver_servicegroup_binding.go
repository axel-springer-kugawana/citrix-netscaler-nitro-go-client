package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverServicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type LbvserverServicegroupBindingKey struct {
	Name             string
	Servicegroupname string
}

type get_lbvserver_servicegroup_binding struct {
	Results []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
}

type add_lbvserver_servicegroup_binding_payload struct {
	lbvserver_servicegroup_binding LbvserverServicegroupBinding
}

func lbvserver_servicegroup_binding_key_to_id_args(key LbvserverServicegroupBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",servicegroupname:" + key.Servicegroupname
	return "", result
}

func (c *NitroClient) AddLbvserverServicegroupBinding(binding LbvserverServicegroupBinding) error {
	payload := add_lbvserver_servicegroup_binding_payload{
		binding,
	}

	return c.put("lbvserver_servicegroup_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) ([]LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding

	id, args := lbvserver_servicegroup_binding_key_to_id_args(key)

	if err := c.get("lbvserver_servicegroup_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverServicegroupBinding() ([]LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding

	if err := c.get("lbvserver_servicegroup_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (*LbvserverServicegroupBinding, error) {
	var results get_lbvserver_servicegroup_binding

	id, args := lbvserver_servicegroup_binding_key_to_id_args(key)

	if err := c.get("lbvserver_servicegroup_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_servicegroup_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_servicegroup_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) error {
	id, args := lbvserver_servicegroup_binding_key_to_id_args(key)

	return c.delete("lbvserver_servicegroup_binding", id, "", args)
}
