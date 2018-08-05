package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCmppolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_cmppolicy_binding_payload struct {
	Resource LbvserverCmppolicyBinding `json:"lbvserver_cmppolicy_binding"`
}

type get_lbvserver_cmppolicy_binding_result struct {
	Results []LbvserverCmppolicyBinding `json:"lbvserver_cmppolicy_binding"`
}

type count_lbvserver_cmppolicy_binding_result struct {
	Results []Count `json:"lbvserver_cmppolicy_binding"`
}

func lbvserver_cmppolicy_binding_key_to_id_qs(key LbvserverCmppolicyBindingKey, arg string) (string, map[string]string) {
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

func lbvserver_cmppolicy_binding_key_to_id_args(key LbvserverCmppolicyBindingKey) (string, map[string]string) {
	return lbvserver_cmppolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_cmppolicy_binding_key_to_id_filter(key LbvserverCmppolicyBindingKey) (string, map[string]string) {
	return lbvserver_cmppolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverCmppolicyBinding(binding LbvserverCmppolicyBinding) error {
	payload := add_lbvserver_cmppolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_cmppolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverCmppolicyBinding() (int, error) {
	var results count_lbvserver_cmppolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_cmppolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverCmppolicyBinding(id string) (int, error) {
	var results count_lbvserver_cmppolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_cmppolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) (bool, error) {
	var results count_lbvserver_cmppolicy_binding_result

	id, qs := lbvserver_cmppolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_cmppolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_cmppolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverCmppolicyBinding() ([]LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_cmppolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverCmppolicyBinding(id string) ([]LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding_result

	if err := c.get("lbvserver_cmppolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) (*LbvserverCmppolicyBinding, error) {
	var results get_lbvserver_cmppolicy_binding_result

	id, qs := lbvserver_cmppolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_cmppolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_cmppolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_cmppolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) error {
	id, qs := lbvserver_cmppolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_cmppolicy_binding", id, qs)
}
