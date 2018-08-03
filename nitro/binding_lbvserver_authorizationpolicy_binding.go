package nitro

import (
	"fmt"
	"strconv"
	"strings"
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

type add_lbvserver_authorizationpolicy_binding_payload struct {
	Resource LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

type get_lbvserver_authorizationpolicy_binding_result struct {
	Results []LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

type count_lbvserver_authorizationpolicy_binding_result struct {
	Results []Count `json:"lbvserver_authorizationpolicy_binding"`
}

func lbvserver_authorizationpolicy_binding_key_to_id_args(key LbvserverAuthorizationpolicyBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddLbvserverAuthorizationpolicyBinding(binding LbvserverAuthorizationpolicyBinding) error {
	payload := add_lbvserver_authorizationpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_authorizationpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverAuthorizationpolicyBinding() (int, error) {
	var results count_lbvserver_authorizationpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_authorizationpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverAuthorizationpolicyBinding(id string) (int, error) {
	var results count_lbvserver_authorizationpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_authorizationpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverAuthorizationpolicyBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountLbvserverAuthorizationpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverAuthorizationpolicyBinding() ([]LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_authorizationpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverAuthorizationpolicyBinding(id string) ([]LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding_result

	if err := c.get("lbvserver_authorizationpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) (*LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding_result

	id, qs := lbvserver_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_authorizationpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_authorizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_authorizationpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) error {
	id, qs := lbvserver_authorizationpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_authorizationpolicy_binding", id, qs)
}
