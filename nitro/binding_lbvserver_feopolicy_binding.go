package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverFeopolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_feopolicy_binding struct {
	Results []LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

type add_lbvserver_feopolicy_binding_payload struct {
	lbvserver_feopolicy_binding LbvserverFeopolicyBinding
}

func lbvserver_feopolicy_binding_key_to_id_args(key LbvserverFeopolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverFeopolicyBinding(binding LbvserverFeopolicyBinding) error {
	payload := add_lbvserver_feopolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_feopolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) ([]LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding

	id, args := lbvserver_feopolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_feopolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverFeopolicyBinding() ([]LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding

	if err := c.get("lbvserver_feopolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (*LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding

	id, args := lbvserver_feopolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_feopolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_feopolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_feopolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) error {
	id, args := lbvserver_feopolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_feopolicy_binding", id, "", args)
}
