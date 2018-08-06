package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverTmtrafficpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverTmtrafficpolicyBindingKey struct {
	Name       string
	Policyname string
}

type add_lbvserver_tmtrafficpolicy_binding_payload struct {
	Resource LbvserverTmtrafficpolicyBinding `json:"lbvserver_tmtrafficpolicy_binding"`
}

type get_lbvserver_tmtrafficpolicy_binding_result struct {
	Results []LbvserverTmtrafficpolicyBinding `json:"lbvserver_tmtrafficpolicy_binding"`
}

type count_lbvserver_tmtrafficpolicy_binding_result struct {
	Results []Count `json:"lbvserver_tmtrafficpolicy_binding"`
}

func lbvserver_tmtrafficpolicy_binding_key_to_id_qs(key LbvserverTmtrafficpolicyBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbvserver_tmtrafficpolicy_binding_key_to_id_args(key LbvserverTmtrafficpolicyBindingKey) (string, map[string]string) {
	return lbvserver_tmtrafficpolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_tmtrafficpolicy_binding_key_to_id_filter(key LbvserverTmtrafficpolicyBindingKey) (string, map[string]string) {
	return lbvserver_tmtrafficpolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverTmtrafficpolicyBinding(binding LbvserverTmtrafficpolicyBinding) error {
	payload := add_lbvserver_tmtrafficpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_tmtrafficpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverTmtrafficpolicyBinding() (int, error) {
	var results count_lbvserver_tmtrafficpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_tmtrafficpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverTmtrafficpolicyBinding(id string) (int, error) {
	var results count_lbvserver_tmtrafficpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) (bool, error) {
	var results count_lbvserver_tmtrafficpolicy_binding_result

	id, qs := lbvserver_tmtrafficpolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_tmtrafficpolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverTmtrafficpolicyBinding() ([]LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_tmtrafficpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverTmtrafficpolicyBinding(id string) ([]LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding_result

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) (*LbvserverTmtrafficpolicyBinding, error) {
	var results get_lbvserver_tmtrafficpolicy_binding_result

	id, qs := lbvserver_tmtrafficpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_tmtrafficpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_tmtrafficpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_tmtrafficpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverTmtrafficpolicyBinding(key LbvserverTmtrafficpolicyBindingKey) error {
	id, qs := lbvserver_tmtrafficpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_tmtrafficpolicy_binding", id, qs)
}
