package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAuthorizationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuthorizationpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource LbvserverAuthorizationpolicyBinding) ToKey() LbvserverAuthorizationpolicyBindingKey {
	key := LbvserverAuthorizationpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key LbvserverAuthorizationpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key LbvserverAuthorizationpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverAuthorizationpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_authorizationpolicy_binding_payload struct {
	Resource LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

func (c *NitroClient) AddLbvserverAuthorizationpolicyBinding(resource LbvserverAuthorizationpolicyBinding) error {
	payload := add_lbvserver_authorizationpolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_authorizationpolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_authorizationpolicy_binding_result struct {
	Results []LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

func (c *NitroClient) ListLbvserverAuthorizationpolicyBinding() ([]LbvserverAuthorizationpolicyBinding, error) {
	results := list_lbvserver_authorizationpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_authorizationpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_authorizationpolicy_binding_result struct {
	Results []Count `json:"lbvserver_authorizationpolicy_binding"`
}

func (c *NitroClient) CountLbvserverAuthorizationpolicyBinding() (int, error) {
	results := count_lbvserver_authorizationpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_authorizationpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_lbvserver_authorizationpolicy_binding_result struct {
	Results []LbvserverAuthorizationpolicyBinding `json:"lbvserver_authorizationpolicy_binding"`
}

func (c *NitroClient) GetLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) (*LbvserverAuthorizationpolicyBinding, error) {
	var results get_lbvserver_authorizationpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_authorizationpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_authorizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_authorizationpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) (bool, error) {
	var results count_lbvserver_authorizationpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_authorizationpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverAuthorizationpolicyBinding(key LbvserverAuthorizationpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_authorizationpolicy_binding", id, qs)
}
