package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverScpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverScpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_scpolicy_binding struct {
	Results []LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

type add_lbvserver_scpolicy_binding_payload struct {
	lbvserver_scpolicy_binding LbvserverScpolicyBinding
}

func lbvserver_scpolicy_binding_key_to_id_args(key LbvserverScpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverScpolicyBinding(binding LbvserverScpolicyBinding) error {
	payload := add_lbvserver_scpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_scpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) ([]LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding

	id, args := lbvserver_scpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_scpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverScpolicyBinding() ([]LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding

	if err := c.get("lbvserver_scpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (*LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding

	id, args := lbvserver_scpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_scpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_scpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_scpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) error {
	id, args := lbvserver_scpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_scpolicy_binding", id, "", args)
}
