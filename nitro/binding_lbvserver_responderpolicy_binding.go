package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverResponderpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_responderpolicy_binding_payload struct {
	Resource LbvserverResponderpolicyBinding `json:"lbvserver_responderpolicy_binding"`
}

type get_lbvserver_responderpolicy_binding_result struct {
	Results []LbvserverResponderpolicyBinding `json:"lbvserver_responderpolicy_binding"`
}

type count_lbvserver_responderpolicy_binding_result struct {
	Results []Count `json:"lbvserver_responderpolicy_binding"`
}

func lbvserver_responderpolicy_binding_key_to_id_args(key LbvserverResponderpolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddLbvserverResponderpolicyBinding(binding LbvserverResponderpolicyBinding) error {
	payload := add_lbvserver_responderpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_responderpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverResponderpolicyBinding() (int, error) {
	var results count_lbvserver_responderpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_responderpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverResponderpolicyBinding(id string) (int, error) {
	var results count_lbvserver_responderpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_responderpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverResponderpolicyBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountLbvserverResponderpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverResponderpolicyBinding() ([]LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_responderpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverResponderpolicyBinding(id string) ([]LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding_result

	if err := c.get("lbvserver_responderpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) (*LbvserverResponderpolicyBinding, error) {
	var results get_lbvserver_responderpolicy_binding_result

	id, qs := lbvserver_responderpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_responderpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_responderpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_responderpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) error {
	id, qs := lbvserver_responderpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_responderpolicy_binding", id, qs)
}
