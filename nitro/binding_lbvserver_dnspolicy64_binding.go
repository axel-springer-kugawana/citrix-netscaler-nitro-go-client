package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverDnspolicy64Binding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverDnspolicy64BindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_dnspolicy64_binding_payload struct {
	Resource LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

type get_lbvserver_dnspolicy64_binding_result struct {
	Results []LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

type count_lbvserver_dnspolicy64_binding_result struct {
	Results []Count `json:"lbvserver_dnspolicy64_binding"`
}

func lbvserver_dnspolicy64_binding_key_to_id_qs(key LbvserverDnspolicy64BindingKey, arg string) (string, map[string]string) {
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

func lbvserver_dnspolicy64_binding_key_to_id_args(key LbvserverDnspolicy64BindingKey) (string, map[string]string) {
	return lbvserver_dnspolicy64_binding_key_to_id_qs(key, "args")
}

func lbvserver_dnspolicy64_binding_key_to_id_filter(key LbvserverDnspolicy64BindingKey) (string, map[string]string) {
	return lbvserver_dnspolicy64_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverDnspolicy64Binding(binding LbvserverDnspolicy64Binding) error {
	payload := add_lbvserver_dnspolicy64_binding_payload{
		binding,
	}

	return c.put("lbvserver_dnspolicy64_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverDnspolicy64Binding() (int, error) {
	var results count_lbvserver_dnspolicy64_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_dnspolicy64_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverDnspolicy64Binding(id string) (int, error) {
	var results count_lbvserver_dnspolicy64_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_dnspolicy64_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (bool, error) {
	var results count_lbvserver_dnspolicy64_binding_result

	id, qs := lbvserver_dnspolicy64_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_dnspolicy64_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_dnspolicy64_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverDnspolicy64Binding() ([]LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_dnspolicy64_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverDnspolicy64Binding(id string) ([]LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding_result

	if err := c.get("lbvserver_dnspolicy64_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (*LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding_result

	id, qs := lbvserver_dnspolicy64_binding_key_to_id_args(key)

	if err := c.get("lbvserver_dnspolicy64_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_dnspolicy64_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_dnspolicy64_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) error {
	id, qs := lbvserver_dnspolicy64_binding_key_to_id_args(key)

	return c.delete("lbvserver_dnspolicy64_binding", id, qs)
}
