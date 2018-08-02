package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCmppolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_cmppolicy_binding struct {
	Results []LbvserverCmppolicyBinding `json:"lbvserver_cmppolicy_binding"`
}

type add_lbvserver_cmppolicy_binding_payload struct {
	lbvserver_cmppolicy_binding LbvserverCmppolicyBinding
}

func lbvserver_cmppolicy_binding_key_to_id_args(key LbvserverCmppolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverCmppolicyBinding(binding LbvserverCmppolicyBinding) error {
	payload := add_lbvserver_cmppolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_cmppolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) ([]LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding

	id, args := lbvserver_cmppolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cmppolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverCmppolicyBinding() ([]LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding

	if err := c.get("lbvserver_cmppolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) (*LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding

	id, args := lbvserver_cmppolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cmppolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_cmppolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_cmppolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) error {
	id, args := lbvserver_cmppolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_cmppolicy_binding", id, "", args)
}
