package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppflowpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_appflowpolicy_binding struct {
	Results []LbvserverAppflowpolicyBinding `json:"lbvserver_appflowpolicy_binding"`
}

type add_lbvserver_appflowpolicy_binding_payload struct {
	lbvserver_appflowpolicy_binding LbvserverAppflowpolicyBinding
}

func lbvserver_appflowpolicy_binding_key_to_id_args(key LbvserverAppflowpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverAppflowpolicyBinding(binding LbvserverAppflowpolicyBinding) error {
	payload := add_lbvserver_appflowpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_appflowpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) ([]LbvserverAppflowpolicyBinding, error) {
	var results get_lbvserver_appflowpolicy_binding

	id, args := lbvserver_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAppflowpolicyBinding() ([]LbvserverAppflowpolicyBinding, error) {
	var results get_lbvserver_appflowpolicy_binding

	if err := c.get("lbvserver_appflowpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) (*LbvserverAppflowpolicyBinding, error) {
	var results get_lbvserver_appflowpolicy_binding

	id, args := lbvserver_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_appflowpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) error {
	id, args := lbvserver_appflowpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_appflowpolicy_binding", id, "", args)
}
