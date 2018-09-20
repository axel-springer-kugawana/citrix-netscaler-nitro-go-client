package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type AuthorizationpolicylabelAuthorizationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invoke_labelname       string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type AuthorizationpolicylabelAuthorizationpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

func (resource AuthorizationpolicylabelAuthorizationpolicyBinding) ToKey() AuthorizationpolicylabelAuthorizationpolicyBindingKey {
	key := AuthorizationpolicylabelAuthorizationpolicyBindingKey{
		resource.Labelname,
		resource.Policyname,
	}

	return key
}

func (key AuthorizationpolicylabelAuthorizationpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Labelname)
	args = append(args, "policyname:"+url.QueryEscape(key.Policyname))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AuthorizationpolicylabelAuthorizationpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuthorizationpolicylabelAuthorizationpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_authorizationpolicylabel_authorizationpolicy_binding_payload struct {
	Resource AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

func (c *NitroClient) AddAuthorizationpolicylabelAuthorizationpolicyBinding(resource AuthorizationpolicylabelAuthorizationpolicyBinding) error {
	payload := add_authorizationpolicylabel_authorizationpolicy_binding_payload{
		resource,
	}

	return c.put("authorizationpolicylabel_authorizationpolicy_binding", "", nil, payload)
}

//      LIST

type list_authorizationpolicylabel_authorizationpolicy_binding_result struct {
	Results []AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

func (c *NitroClient) ListAuthorizationpolicylabelAuthorizationpolicyBinding() ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	results := list_authorizationpolicylabel_authorizationpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_authorizationpolicylabel_authorizationpolicy_binding_result struct {
	Results []Count `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

func (c *NitroClient) CountAuthorizationpolicylabelAuthorizationpolicyBinding() (int, error) {
	results := count_authorizationpolicylabel_authorizationpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_authorizationpolicylabel_authorizationpolicy_binding_result struct {
	Results []AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

func (c *NitroClient) GetAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (*AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicylabel_authorizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("authorizationpolicylabel_authorizationpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (bool, error) {
	var results count_authorizationpolicylabel_authorizationpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("authorizationpolicylabel_authorizationpolicy_binding", id, qs)
}
