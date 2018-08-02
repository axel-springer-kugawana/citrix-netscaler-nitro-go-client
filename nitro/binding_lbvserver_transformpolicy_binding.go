package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverTransformpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverTransformpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_transformpolicy_binding struct {
	Results []LbvserverTransformpolicyBinding `json:"lbvserver_transformpolicy_binding"`
}

type add_lbvserver_transformpolicy_binding_payload struct {
	lbvserver_transformpolicy_binding LbvserverTransformpolicyBinding
}

func lbvserver_transformpolicy_binding_key_to_id_args(key LbvserverTransformpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverTransformpolicyBinding(binding LbvserverTransformpolicyBinding) error {
	payload := add_lbvserver_transformpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_transformpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverTransformpolicyBinding(key LbvserverTransformpolicyBindingKey) ([]LbvserverTransformpolicyBinding, error) {
	var results get_lbvserver_transformpolicy_binding

	id, args := lbvserver_transformpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_transformpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverTransformpolicyBinding() ([]LbvserverTransformpolicyBinding, error) {
	var results get_lbvserver_transformpolicy_binding

	if err := c.get("lbvserver_transformpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverTransformpolicyBinding(key LbvserverTransformpolicyBindingKey) (*LbvserverTransformpolicyBinding, error) {
	var results get_lbvserver_transformpolicy_binding

	id, args := lbvserver_transformpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_transformpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_transformpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_transformpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverTransformpolicyBinding(key LbvserverTransformpolicyBindingKey) error {
	id, args := lbvserver_transformpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_transformpolicy_binding", id, "", args)
}
