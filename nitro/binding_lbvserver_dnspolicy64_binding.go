package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type LbvserverDnspolicy64Binding struct {
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
}

func (resource LbvserverDnspolicy64Binding) ToKey() LbvserverDnspolicy64BindingKey {
	key := LbvserverDnspolicy64BindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key LbvserverDnspolicy64BindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)
	args = append(args, "policyname:"+url.QueryEscape(key.Policyname))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverDnspolicy64BindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverDnspolicy64BindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_dnspolicy64_binding_payload struct {
	Resource LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) AddLbvserverDnspolicy64Binding(resource LbvserverDnspolicy64Binding) error {
	payload := add_lbvserver_dnspolicy64_binding_payload{
		resource,
	}

	return c.put("lbvserver_dnspolicy64_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_dnspolicy64_binding_result struct {
	Results []LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) ListLbvserverDnspolicy64Binding() ([]LbvserverDnspolicy64Binding, error) {
	results := list_lbvserver_dnspolicy64_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_dnspolicy64_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_dnspolicy64_binding_result struct {
	Results []Count `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) CountLbvserverDnspolicy64Binding() (int, error) {
	results := count_lbvserver_dnspolicy64_binding_result{}

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

//      READ

type get_lbvserver_dnspolicy64_binding_result struct {
	Results []LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) GetLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (*LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_dnspolicy64_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_dnspolicy64_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_dnspolicy64_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (bool, error) {
	var results count_lbvserver_dnspolicy64_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_dnspolicy64_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_dnspolicy64_binding", id, qs)
}
