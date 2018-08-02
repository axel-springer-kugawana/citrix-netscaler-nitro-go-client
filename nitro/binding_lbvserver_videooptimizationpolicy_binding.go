package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverVideooptimizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverVideooptimizationpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_videooptimizationpolicy_binding struct {
	Results []LbvserverVideooptimizationpolicyBinding `json:"lbvserver_videooptimizationpolicy_binding"`
}

type add_lbvserver_videooptimizationpolicy_binding_payload struct {
	lbvserver_videooptimizationpolicy_binding LbvserverVideooptimizationpolicyBinding
}

func lbvserver_videooptimizationpolicy_binding_key_to_id_args(key LbvserverVideooptimizationpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverVideooptimizationpolicyBinding(binding LbvserverVideooptimizationpolicyBinding) error {
	payload := add_lbvserver_videooptimizationpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_videooptimizationpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) ([]LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding

	id, args := lbvserver_videooptimizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_videooptimizationpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverVideooptimizationpolicyBinding() ([]LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding

	if err := c.get("lbvserver_videooptimizationpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) (*LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding

	id, args := lbvserver_videooptimizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_videooptimizationpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_videooptimizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_videooptimizationpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) error {
	id, args := lbvserver_videooptimizationpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_videooptimizationpolicy_binding", id, "", args)
}
