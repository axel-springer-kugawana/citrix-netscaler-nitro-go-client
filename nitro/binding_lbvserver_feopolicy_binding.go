package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverFeopolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_feopolicy_binding_payload struct {
	Resource LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

type get_lbvserver_feopolicy_binding_result struct {
	Results []LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

type count_lbvserver_feopolicy_binding_result struct {
	Results []Count `json:"lbvserver_feopolicy_binding"`
}

func lbvserver_feopolicy_binding_key_to_id_qs(key LbvserverFeopolicyBindingKey, arg string) (string, map[string]string) {
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

func lbvserver_feopolicy_binding_key_to_id_args(key LbvserverFeopolicyBindingKey) (string, map[string]string) {
	return lbvserver_feopolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_feopolicy_binding_key_to_id_filter(key LbvserverFeopolicyBindingKey) (string, map[string]string) {
	return lbvserver_feopolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverFeopolicyBinding(binding LbvserverFeopolicyBinding) error {
	payload := add_lbvserver_feopolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_feopolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverFeopolicyBinding() (int, error) {
	var results count_lbvserver_feopolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_feopolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverFeopolicyBinding(id string) (int, error) {
	var results count_lbvserver_feopolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_feopolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (bool, error) {
	var results count_lbvserver_feopolicy_binding_result

	id, qs := lbvserver_feopolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_feopolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_feopolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverFeopolicyBinding() ([]LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_feopolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverFeopolicyBinding(id string) ([]LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding_result

	if err := c.get("lbvserver_feopolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (*LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding_result

	id, qs := lbvserver_feopolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_feopolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_feopolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_feopolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) error {
	id, qs := lbvserver_feopolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_feopolicy_binding", id, qs)
}
