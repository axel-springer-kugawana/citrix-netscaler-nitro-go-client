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

type add_lbvserver_cachepolicy_binding_payload struct {
	Resources LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

type get_lbvserver_cachepolicy_binding_result struct {
	Results []LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

type count_lbvserver_cachepolicy_binding_result struct {
	Results []Count `json:"lbvserver_cachepolicy_binding"`
}

func lbvserver_cachepolicy_binding_key_to_id_args(key LbvserverCachepolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddLbvserverCachepolicyBinding(binding LbvserverCachepolicyBinding) error {
	payload := add_lbvserver_cachepolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_cachepolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverCachepolicyBinding() (int, error) {
	var results count_lbvserver_cachepolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_cachepolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverCachepolicyBinding(id string) (int, error) {
	var results count_lbvserver_cachepolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_cachepolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverCachepolicyBinding(id string) (bool, error) {
	if count, err := c.CountLbvserverCachepolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverCachepolicyBinding() ([]LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_cachepolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverCachepolicyBinding(id string) ([]LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding_result

	if err := c.get("lbvserver_cachepolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) (*LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding_result

	id, qs := lbvserver_cachepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cachepolicy_binding", id, qs, &results); err != nil {
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
	id, qs := lbvserver_cachepolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_cachepolicy_binding", id, qs)
}
