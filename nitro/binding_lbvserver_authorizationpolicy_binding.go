package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverAuthorizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuthorizationpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_authorizationpolicy_binding struct {
	Results []LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

type add_lbvserver_authorizationpolicy_binding_payload struct {
	lbvserver_authorizationpolicy_binding LbvserverAuthorizationpolicyBinding
}

func lbvserver_authorizationpolicy_binding_key_to_id_args(key LbvserverAuthorizationpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverAuthorizationpolicyBinding(binding LbvserverAuthorizationpolicyBinding) error {
	payload := add_lbvserver_authorizationpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_authorizationpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) ([]LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding

	id, args := lbvserver_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_authorizationpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAuthorizationpolicyBinding() ([]LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding

	if err := c.get("lbvserver_authorizationpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) (*LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding

	id, args := lbvserver_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_authorizationpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_authorizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_authorizationpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) error {
	id, args := lbvserver_authorizationpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_authorizationpolicy_binding", id, "", args)
}
