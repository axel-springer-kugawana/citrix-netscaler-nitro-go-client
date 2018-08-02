package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCachepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_cachepolicy_binding struct {
	Results []LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

type add_lbvserver_cachepolicy_binding_payload struct {
	lbvserver_cachepolicy_binding LbvserverCachepolicyBinding
}

func lbvserver_cachepolicy_binding_key_to_id_args(key LbvserverCachepolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverCachepolicyBinding(binding LbvserverCachepolicyBinding) error {
	payload := add_lbvserver_cachepolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_cachepolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) ([]LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding

	id, args := lbvserver_cachepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cachepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverCachepolicyBinding() ([]LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding

	if err := c.get("lbvserver_cachepolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) (*LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding

	id, args := lbvserver_cachepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cachepolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_cachepolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_cachepolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) error {
	id, args := lbvserver_cachepolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_cachepolicy_binding", id, "", args)
}
