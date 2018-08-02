package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppfwpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_appfwpolicy_binding struct {
	Results []LbvserverAppfwpolicyBinding `json:"lbvserver_appfwpolicy_binding"`
}

type add_lbvserver_appfwpolicy_binding_payload struct {
	lbvserver_appfwpolicy_binding LbvserverAppfwpolicyBinding
}

func lbvserver_appfwpolicy_binding_key_to_id_args(key LbvserverAppfwpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverAppfwpolicyBinding(binding LbvserverAppfwpolicyBinding) error {
	payload := add_lbvserver_appfwpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_appfwpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) ([]LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding

	id, args := lbvserver_appfwpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appfwpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAppfwpolicyBinding() ([]LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding

	if err := c.get("lbvserver_appfwpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) (*LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding

	id, args := lbvserver_appfwpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appfwpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appfwpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_appfwpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) error {
	id, args := lbvserver_appfwpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_appfwpolicy_binding", id, "", args)
}
