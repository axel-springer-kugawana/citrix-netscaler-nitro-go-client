package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCachepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (resource LbvserverCachepolicyBinding) ToKey() LbvserverCachepolicyBindingKey {
	key := LbvserverCachepolicyBindingKey{
		resource.Name,
		resource.Policyname,
		resource.Bindpoint,
	}

	return key
}

func (key LbvserverCachepolicyBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_lbvserver_cachepolicy_binding_payload struct {
	Resource LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

func (c *NitroClient) AddLbvserverCachepolicyBinding(resource LbvserverCachepolicyBinding) error {
	payload := add_lbvserver_cachepolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_cachepolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_cachepolicy_binding_result struct {
	Results []LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

func (c *NitroClient) ListLbvserverCachepolicyBinding() ([]LbvserverCachepolicyBinding, error) {
	results := list_lbvserver_cachepolicy_binding_result{}

	if err := c.get("lbvserver_cachepolicy_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbvserver_cachepolicy_binding_result struct {
	Results []LbvserverCachepolicyBinding `json:"lbvserver_cachepolicy_binding"`
}

func (c *NitroClient) GetLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) (*LbvserverCachepolicyBinding, error) {
	var results get_lbvserver_cachepolicy_binding_result

	id, qs := key.to_id_args()

	if err := c.get("lbvserver_cachepolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_cachepolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_cachepolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_lbvserver_cachepolicy_binding_result struct {
	Results []Count `json:"lbvserver_cachepolicy_binding"`
}

func (c *NitroClient) ExistsLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) (bool, error) {
	var results count_lbvserver_cachepolicy_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbvserver_cachepolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverCachepolicyBinding(key LbvserverCachepolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_cachepolicy_binding", id, qs)
}
