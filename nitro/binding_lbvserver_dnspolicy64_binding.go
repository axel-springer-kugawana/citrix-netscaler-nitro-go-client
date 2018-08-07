package nitro

import (
	"fmt"
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

func (key LbvserverDnspolicy64BindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
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

	if err := c.get("lbvserver_dnspolicy64_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbvserver_dnspolicy64_binding_result struct {
	Results []LbvserverDnspolicy64Binding `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) GetLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (*LbvserverDnspolicy64Binding, error) {
	var results get_lbvserver_dnspolicy64_binding_result

	id, qs := key.to_id_args()

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

type count_lbvserver_dnspolicy64_binding_result struct {
	Results []Count `json:"lbvserver_dnspolicy64_binding"`
}

func (c *NitroClient) ExistsLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (bool, error) {
	var results count_lbvserver_dnspolicy64_binding_result

	id, qs := key.to_id_args()

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
