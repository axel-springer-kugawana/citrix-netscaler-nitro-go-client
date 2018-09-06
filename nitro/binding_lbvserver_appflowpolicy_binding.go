package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppflowpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource LbvserverAppflowpolicyBinding) ToKey() LbvserverAppflowpolicyBindingKey {
	key := LbvserverAppflowpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key LbvserverAppflowpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverAppflowpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverAppflowpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_appflowpolicy_binding_payload struct {
	Resource LbvserverAppflowpolicyBinding `json:"lbvserver_appflowpolicy_binding"`
}

func (c *NitroClient) AddLbvserverAppflowpolicyBinding(resource LbvserverAppflowpolicyBinding) error {
	payload := add_lbvserver_appflowpolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_appflowpolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_appflowpolicy_binding_result struct {
	Results []LbvserverAppflowpolicyBinding `json:"lbvserver_appflowpolicy_binding"`
}

func (c *NitroClient) ListLbvserverAppflowpolicyBinding() ([]LbvserverAppflowpolicyBinding, error) {
	results := list_lbvserver_appflowpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_appflowpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_appflowpolicy_binding_result struct {
	Results []Count `json:"lbvserver_appflowpolicy_binding"`
}

func (c *NitroClient) CountLbvserverAppflowpolicyBinding() (int, error) {
	results := count_lbvserver_appflowpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_appflowpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_lbvserver_appflowpolicy_binding_result struct {
	Results []LbvserverAppflowpolicyBinding `json:"lbvserver_appflowpolicy_binding"`
}

func (c *NitroClient) GetLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) (*LbvserverAppflowpolicyBinding, error) {
	var results get_lbvserver_appflowpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_appflowpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_appflowpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) (bool, error) {
	var results count_lbvserver_appflowpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_appflowpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_appflowpolicy_binding", id, qs)
}
