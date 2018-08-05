package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverScpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverScpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type add_lbvserver_scpolicy_binding_payload struct {
	Resource LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

type get_lbvserver_scpolicy_binding_result struct {
	Results []LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

type count_lbvserver_scpolicy_binding_result struct {
	Results []Count `json:"lbvserver_scpolicy_binding"`
}

func lbvserver_scpolicy_binding_key_to_id_args(key LbvserverScpolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddLbvserverScpolicyBinding(binding LbvserverScpolicyBinding) error {
	payload := add_lbvserver_scpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_scpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverScpolicyBinding() (int, error) {
	var results count_lbvserver_scpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_scpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverScpolicyBinding(id string) (int, error) {
	var results count_lbvserver_scpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_scpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (bool, error) {
	// TODO : wrong implementation
	return false, nil
	//        if count, err := c.CountLbvserverScpolicyBinding(id); err != nil {
	//                return false, err
	//        } else {
	//                return count == 1, nil
	//        }
}

func (c *NitroClient) BulkListLbvserverScpolicyBinding() ([]LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_scpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverScpolicyBinding(id string) ([]LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding_result

	if err := c.get("lbvserver_scpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (*LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding_result

	id, qs := lbvserver_scpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_scpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_scpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_scpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) error {
	id, qs := lbvserver_scpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_scpolicy_binding", id, qs)
}
