package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverCapolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCapolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_capolicy_binding struct {
	Results []LbvserverCapolicyBinding `json:"lbvserver_capolicy_binding"`
}

type add_lbvserver_capolicy_binding_payload struct {
	lbvserver_capolicy_binding LbvserverCapolicyBinding
}

func lbvserver_capolicy_binding_key_to_id_args(key LbvserverCapolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverCapolicyBinding(binding LbvserverCapolicyBinding) error {
	payload := add_lbvserver_capolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_capolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverCapolicyBinding(key LbvserverCapolicyBindingKey) ([]LbvserverCapolicyBinding, error) {
	var results get_lbvserver_capolicy_binding

	id, args := lbvserver_capolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_capolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverCapolicyBinding() ([]LbvserverCapolicyBinding, error) {
	var results get_lbvserver_capolicy_binding

	if err := c.get("lbvserver_capolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverCapolicyBinding(key LbvserverCapolicyBindingKey) (*LbvserverCapolicyBinding, error) {
	var results get_lbvserver_capolicy_binding

	id, args := lbvserver_capolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_capolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_capolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_capolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverCapolicyBinding(key LbvserverCapolicyBindingKey) error {
	id, args := lbvserver_capolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_capolicy_binding", id, "", args)
}
