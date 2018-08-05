package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppfwpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_appfwpolicy_binding_payload struct {
	Resource LbvserverAppfwpolicyBinding `json:"lbvserver_appfwpolicy_binding"`
}

type get_lbvserver_appfwpolicy_binding_result struct {
	Results []LbvserverAppfwpolicyBinding `json:"lbvserver_appfwpolicy_binding"`
}

type count_lbvserver_appfwpolicy_binding_result struct {
	Results []Count `json:"lbvserver_appfwpolicy_binding"`
}

func lbvserver_appfwpolicy_binding_key_to_id_qs(key LbvserverAppfwpolicyBindingKey, arg string) (string, map[string]string) {
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

func lbvserver_appfwpolicy_binding_key_to_id_args(key LbvserverAppfwpolicyBindingKey) (string, map[string]string) {
	return lbvserver_appfwpolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_appfwpolicy_binding_key_to_id_filter(key LbvserverAppfwpolicyBindingKey) (string, map[string]string) {
	return lbvserver_appfwpolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverAppfwpolicyBinding(binding LbvserverAppfwpolicyBinding) error {
	payload := add_lbvserver_appfwpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_appfwpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverAppfwpolicyBinding() (int, error) {
	var results count_lbvserver_appfwpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_appfwpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverAppfwpolicyBinding(id string) (int, error) {
	var results count_lbvserver_appfwpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_appfwpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) (bool, error) {
	var results count_lbvserver_appfwpolicy_binding_result

	id, qs := lbvserver_appfwpolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_appfwpolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_appfwpolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverAppfwpolicyBinding() ([]LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_appfwpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverAppfwpolicyBinding(id string) ([]LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding_result

	if err := c.get("lbvserver_appfwpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) (*LbvserverAppfwpolicyBinding, error) {
	var results get_lbvserver_appfwpolicy_binding_result

	id, qs := lbvserver_appfwpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_appfwpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appfwpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_appfwpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) error {
	id, qs := lbvserver_appfwpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_appfwpolicy_binding", id, qs)
}
