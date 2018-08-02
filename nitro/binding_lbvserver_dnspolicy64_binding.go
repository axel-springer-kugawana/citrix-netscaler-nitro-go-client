package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverDnspolicy64Binding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverDnspolicy64BindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_dnspolicy64_binding struct {
	Results []LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

type add_lbvserver_dnspolicy64_binding_payload struct {
	lbvserver_dnspolicy64_binding LbvserverDnspolicy64Binding
}

func lbvserver_dnspolicy64_binding_key_to_id_args(key LbvserverDnspolicy64BindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverDnspolicy64Binding(binding LbvserverDnspolicy64Binding) error {
	payload := add_lbvserver_dnspolicy64_binding_payload{
		binding,
	}

	return c.put("lbvserver_dnspolicy64_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) ([]LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding

	id, args := lbvserver_dnspolicy64_binding_key_to_id_args(key)

	if err := c.get("lbvserver_dnspolicy64_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverDnspolicy64Binding() ([]LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding

	if err := c.get("lbvserver_dnspolicy64_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (*LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding

	id, args := lbvserver_dnspolicy64_binding_key_to_id_args(key)

	if err := c.get("lbvserver_dnspolicy64_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_dnspolicy64_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_dnspolicy64_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) error {
	id, args := lbvserver_dnspolicy64_binding_key_to_id_args(key)

	return c.delete("lbvserver_dnspolicy64_binding", id, "", args)
}
