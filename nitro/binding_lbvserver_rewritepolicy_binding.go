package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type LbvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverRewritepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (resource LbvserverRewritepolicyBinding) ToKey() LbvserverRewritepolicyBindingKey {
	key := LbvserverRewritepolicyBindingKey{
		resource.Name,
		resource.Policyname,
		resource.Bindpoint,
	}

	return key
}

func (key LbvserverRewritepolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)
	args = append(args, "policyname:"+url.QueryEscape(key.Policyname))
	args = append(args, "bindpoint:"+url.QueryEscape(key.Bindpoint))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverRewritepolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverRewritepolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_rewritepolicy_binding_payload struct {
	Resource LbvserverRewritepolicyBinding `json:"lbvserver_rewritepolicy_binding"`
}

func (c *NitroClient) AddLbvserverRewritepolicyBinding(resource LbvserverRewritepolicyBinding) error {
	payload := add_lbvserver_rewritepolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_rewritepolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_rewritepolicy_binding_result struct {
	Results []LbvserverRewritepolicyBinding `json:"lbvserver_rewritepolicy_binding"`
}

func (c *NitroClient) ListLbvserverRewritepolicyBinding() ([]LbvserverRewritepolicyBinding, error) {
	results := list_lbvserver_rewritepolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_rewritepolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_rewritepolicy_binding_result struct {
	Results []Count `json:"lbvserver_rewritepolicy_binding"`
}

func (c *NitroClient) CountLbvserverRewritepolicyBinding() (int, error) {
	results := count_lbvserver_rewritepolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_rewritepolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_lbvserver_rewritepolicy_binding_result struct {
	Results []LbvserverRewritepolicyBinding `json:"lbvserver_rewritepolicy_binding"`
}

func (c *NitroClient) GetLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) (*LbvserverRewritepolicyBinding, error) {
	var results get_lbvserver_rewritepolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_rewritepolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_rewritepolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_rewritepolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) (bool, error) {
	var results count_lbvserver_rewritepolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_rewritepolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverRewritepolicyBinding(key LbvserverRewritepolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_rewritepolicy_binding", id, qs)
}
