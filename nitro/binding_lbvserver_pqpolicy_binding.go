package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverPqpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverPqpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_pqpolicy_binding_payload struct {
	Resource LbvserverPqpolicyBinding `json:"lbvserver_pqpolicy_binding"`
}

type get_lbvserver_pqpolicy_binding_result struct {
	Results []LbvserverPqpolicyBinding `json:"lbvserver_pqpolicy_binding"`
}

type count_lbvserver_pqpolicy_binding_result struct {
	Results []Count `json:"lbvserver_pqpolicy_binding"`
}

func lbvserver_pqpolicy_binding_key_to_id_qs(key LbvserverPqpolicyBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbvserver_pqpolicy_binding_key_to_id_args(key LbvserverPqpolicyBindingKey) (string, map[string]string) {
	return lbvserver_pqpolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_pqpolicy_binding_key_to_id_filter(key LbvserverPqpolicyBindingKey) (string, map[string]string) {
	return lbvserver_pqpolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverPqpolicyBinding(binding LbvserverPqpolicyBinding) error {
	payload := add_lbvserver_pqpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_pqpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverPqpolicyBinding() (int, error) {
	var results count_lbvserver_pqpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_pqpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverPqpolicyBinding(id string) (int, error) {
	var results count_lbvserver_pqpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_pqpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) (bool, error) {
	var results count_lbvserver_pqpolicy_binding_result

	id, qs := lbvserver_pqpolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_pqpolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_pqpolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverPqpolicyBinding() ([]LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_pqpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverPqpolicyBinding(id string) ([]LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding_result

	if err := c.get("lbvserver_pqpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) (*LbvserverPqpolicyBinding, error) {
	var results get_lbvserver_pqpolicy_binding_result

	id, qs := lbvserver_pqpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_pqpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_pqpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_pqpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverPqpolicyBinding(key LbvserverPqpolicyBindingKey) error {
	id, qs := lbvserver_pqpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_pqpolicy_binding", id, qs)
}
