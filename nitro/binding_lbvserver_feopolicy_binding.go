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

func (resource LbvserverFeopolicyBinding) ToKey() LbvserverFeopolicyBindingKey {
	key := LbvserverFeopolicyBindingKey{
		resource.Name,
		resource.Policyname,
		resource.Bindpoint,
	}

	return key
}

func (key LbvserverFeopolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverFeopolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverFeopolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_feopolicy_binding_payload struct {
	Resource LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

func (c *NitroClient) AddLbvserverFeopolicyBinding(resource LbvserverFeopolicyBinding) error {
	payload := add_lbvserver_feopolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_feopolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_feopolicy_binding_result struct {
	Results []LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

func (c *NitroClient) ListLbvserverFeopolicyBinding() ([]LbvserverFeopolicyBinding, error) {
	results := list_lbvserver_feopolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_feopolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_feopolicy_binding_result struct {
	Results []Count `json:"lbvserver_feopolicy_binding"`
}

func (c *NitroClient) CountLbvserverFeopolicyBinding() (int, error) {
	results := count_lbvserver_feopolicy_binding_result{}

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

//      READ

type get_lbvserver_feopolicy_binding_result struct {
	Results []LbvserverFeopolicyBinding `json:"lbvserver_feopolicy_binding"`
}

func (c *NitroClient) GetLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (*LbvserverFeopolicyBinding, error) {
	var results get_lbvserver_feopolicy_binding_result

	id, qs := key.to_id_filter()

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

//      EXISTS

func (c *NitroClient) ExistsLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) (bool, error) {
	var results count_lbvserver_feopolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_feopolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverFeopolicyBinding(key LbvserverFeopolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_feopolicy_binding", id, qs)
}
