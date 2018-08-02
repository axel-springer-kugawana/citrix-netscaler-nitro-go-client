package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppqoepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_appqoepolicy_binding struct {
	Results []LbvserverAppqoepolicyBinding `json:"lbvserver_appqoepolicy_binding"`
}

type add_lbvserver_appqoepolicy_binding_payload struct {
	lbvserver_appqoepolicy_binding LbvserverAppqoepolicyBinding
}

func lbvserver_appqoepolicy_binding_key_to_id_args(key LbvserverAppqoepolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverAppqoepolicyBinding(binding LbvserverAppqoepolicyBinding) error {
	payload := add_lbvserver_appqoepolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_appqoepolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) ([]LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding

	id, args := lbvserver_appqoepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appqoepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAppqoepolicyBinding() ([]LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding

	if err := c.get("lbvserver_appqoepolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) (*LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding

	id, args := lbvserver_appqoepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appqoepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appqoepolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_appqoepolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) error {
	id, args := lbvserver_appqoepolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_appqoepolicy_binding", id, "", args)
}
