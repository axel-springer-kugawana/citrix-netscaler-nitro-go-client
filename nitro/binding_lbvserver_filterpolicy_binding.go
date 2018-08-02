package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverFilterpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverFilterpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_filterpolicy_binding struct {
	Results []LbvserverFilterpolicyBinding `json:"lbvserver_filterpolicy_binding"`
}

type add_lbvserver_filterpolicy_binding_payload struct {
	lbvserver_filterpolicy_binding LbvserverFilterpolicyBinding
}

func lbvserver_filterpolicy_binding_key_to_id_args(key LbvserverFilterpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverFilterpolicyBinding(binding LbvserverFilterpolicyBinding) error {
	payload := add_lbvserver_filterpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_filterpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverFilterpolicyBinding(key LbvserverFilterpolicyBindingKey) ([]LbvserverFilterpolicyBinding, error) {
	var results get_lbvserver_filterpolicy_binding

	id, args := lbvserver_filterpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_filterpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverFilterpolicyBinding() ([]LbvserverFilterpolicyBinding, error) {
	var results get_lbvserver_filterpolicy_binding

	if err := c.get("lbvserver_filterpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverFilterpolicyBinding(key LbvserverFilterpolicyBindingKey) (*LbvserverFilterpolicyBinding, error) {
	var results get_lbvserver_filterpolicy_binding

	id, args := lbvserver_filterpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_filterpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_filterpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_filterpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverFilterpolicyBinding(key LbvserverFilterpolicyBindingKey) error {
	id, args := lbvserver_filterpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_filterpolicy_binding", id, "", args)
}
