package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppqoepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_appqoepolicy_binding_payload struct {
	Resource LbvserverAppqoepolicyBinding `json:"lbvserver_appqoepolicy_binding"`
}

type get_lbvserver_appqoepolicy_binding_result struct {
	Results []LbvserverAppqoepolicyBinding `json:"lbvserver_appqoepolicy_binding"`
}

type count_lbvserver_appqoepolicy_binding_result struct {
	Results []Count `json:"lbvserver_appqoepolicy_binding"`
}

func lbvserver_appqoepolicy_binding_key_to_id_args(key LbvserverAppqoepolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddLbvserverAppqoepolicyBinding(binding LbvserverAppqoepolicyBinding) error {
	payload := add_lbvserver_appqoepolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_appqoepolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverAppqoepolicyBinding() (int, error) {
	var results count_lbvserver_appqoepolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_appqoepolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverAppqoepolicyBinding(id string) (int, error) {
	var results count_lbvserver_appqoepolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_appqoepolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverAppqoepolicyBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountLbvserverAppqoepolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverAppqoepolicyBinding() ([]LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_appqoepolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverAppqoepolicyBinding(id string) ([]LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding_result

	if err := c.get("lbvserver_appqoepolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) (*LbvserverAppqoepolicyBinding, error) {
	var results get_lbvserver_appqoepolicy_binding_result

	id, qs := lbvserver_appqoepolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appqoepolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appqoepolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_appqoepolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) error {
	id, qs := lbvserver_appqoepolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_appqoepolicy_binding", id, qs)
}
