package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverVideooptimizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverVideooptimizationpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_videooptimizationpolicy_binding_payload struct {
	Resources LbvserverVideooptimizationpolicyBinding `json:"lbvserver_videooptimizationpolicy_binding"`
}

type get_lbvserver_videooptimizationpolicy_binding_result struct {
	Results []LbvserverVideooptimizationpolicyBinding `json:"lbvserver_videooptimizationpolicy_binding"`
}

type count_lbvserver_videooptimizationpolicy_binding_result struct {
	Results []Count `json:"lbvserver_videooptimizationpolicy_binding"`
}

func lbvserver_videooptimizationpolicy_binding_key_to_id_args(key LbvserverVideooptimizationpolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddLbvserverVideooptimizationpolicyBinding(binding LbvserverVideooptimizationpolicyBinding) error {
	payload := add_lbvserver_videooptimizationpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_videooptimizationpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverVideooptimizationpolicyBinding() (int, error) {
	var results count_lbvserver_videooptimizationpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_videooptimizationpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverVideooptimizationpolicyBinding(id string) (int, error) {
	var results count_lbvserver_videooptimizationpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_videooptimizationpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverVideooptimizationpolicyBinding(id string) (bool, error) {
	if count, err := c.CountLbvserverVideooptimizationpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverVideooptimizationpolicyBinding() ([]LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_videooptimizationpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverVideooptimizationpolicyBinding(id string) ([]LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding_result

	if err := c.get("lbvserver_videooptimizationpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) (*LbvserverVideooptimizationpolicyBinding, error) {
	var results get_lbvserver_videooptimizationpolicy_binding_result

	id, qs := lbvserver_videooptimizationpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_videooptimizationpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_videooptimizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_videooptimizationpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverVideooptimizationpolicyBinding(key LbvserverVideooptimizationpolicyBindingKey) error {
	id, qs := lbvserver_videooptimizationpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_videooptimizationpolicy_binding", id, qs)
}
