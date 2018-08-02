package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverPqpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverPqpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_pqpolicy_binding struct {
	Results []LbvserverPqpolicyBinding `json:"lbvserver_pqpolicy_binding"`
}

type add_lbvserver_pqpolicy_binding_payload struct {
	lbvserver_pqpolicy_binding LbvserverPqpolicyBinding
}

func lbvserver_pqpolicy_binding_key_to_id_args(key LbvserverPqpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverPqpolicyBinding(binding LbvserverPqpolicyBinding) error {
	payload := add_lbvserver_pqpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_pqpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) ([]LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding

	id, args := lbvserver_pqpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_pqpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverPqpolicyBinding() ([]LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding

	if err := c.get("lbvserver_pqpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) (*LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding

	id, args := lbvserver_pqpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_pqpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_pqpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_pqpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) error {
	id, args := lbvserver_pqpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_pqpolicy_binding", id, "", args)
}
