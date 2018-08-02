package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverResponderpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_responderpolicy_binding struct {
	Results []LbvserverResponderpolicyBinding `json:"lbvserver_responderpolicy_binding"`
}

type add_lbvserver_responderpolicy_binding_payload struct {
	lbvserver_responderpolicy_binding LbvserverResponderpolicyBinding
}

func lbvserver_responderpolicy_binding_key_to_id_args(key LbvserverResponderpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverResponderpolicyBinding(binding LbvserverResponderpolicyBinding) error {
	payload := add_lbvserver_responderpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_responderpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) ([]LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding

	id, args := lbvserver_responderpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_responderpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverResponderpolicyBinding() ([]LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding

	if err := c.get("lbvserver_responderpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) (*LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding

	id, args := lbvserver_responderpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_responderpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_responderpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_responderpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) error {
	id, args := lbvserver_responderpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_responderpolicy_binding", id, "", args)
}
