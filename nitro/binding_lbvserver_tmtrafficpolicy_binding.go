package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverTmtrafficpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverTmtrafficpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_tmtrafficpolicy_binding struct {
	Results []LbvserverTmtrafficpolicyBinding `json:"lbvserver_tmtrafficpolicy_binding"`
}

type add_lbvserver_tmtrafficpolicy_binding_payload struct {
	lbvserver_tmtrafficpolicy_binding LbvserverTmtrafficpolicyBinding
}

func lbvserver_tmtrafficpolicy_binding_key_to_id_args(key LbvserverTmtrafficpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverTmtrafficpolicyBinding(binding LbvserverTmtrafficpolicyBinding) error {
	payload := add_lbvserver_tmtrafficpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_tmtrafficpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) ([]LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding

	id, args := lbvserver_tmtrafficpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverTmtrafficpolicyBinding() ([]LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding

	if err := c.get("lbvserver_tmtrafficpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) (*LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding

	id, args := lbvserver_tmtrafficpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_tmtrafficpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_tmtrafficpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) error {
	id, args := lbvserver_tmtrafficpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_tmtrafficpolicy_binding", id, "", args)
}
