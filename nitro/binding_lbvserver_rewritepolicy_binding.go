package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverRewritepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_rewritepolicy_binding struct {
	Results []LbvserverRewritepolicyBinding `json:"lbvserver_rewritepolicy_binding"`
}

type add_lbvserver_rewritepolicy_binding_payload struct {
	lbvserver_rewritepolicy_binding LbvserverRewritepolicyBinding
}

func lbvserver_rewritepolicy_binding_key_to_id_args(key LbvserverRewritepolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverRewritepolicyBinding(binding LbvserverRewritepolicyBinding) error {
	payload := add_lbvserver_rewritepolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_rewritepolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) ([]LbvserverRewritepolicyBinding, error) {
	var results get_lbvserver_rewritepolicy_binding

	id, args := lbvserver_rewritepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_rewritepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverRewritepolicyBinding() ([]LbvserverRewritepolicyBinding, error) {
	var results get_lbvserver_rewritepolicy_binding

	if err := c.get("lbvserver_rewritepolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) (*LbvserverRewritepolicyBinding, error) {
	var results get_lbvserver_rewritepolicy_binding

	id, args := lbvserver_rewritepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_rewritepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_rewritepolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_rewritepolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) error {
	id, args := lbvserver_rewritepolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_rewritepolicy_binding", id, "", args)
}
