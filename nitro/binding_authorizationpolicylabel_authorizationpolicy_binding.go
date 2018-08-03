package nitro

import (
	"fmt"
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

type add_authorizationpolicylabel_authorizationpolicy_binding_payload struct {
	Resource AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

type get_authorizationpolicylabel_authorizationpolicy_binding_result struct {
	Results []AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

type count_authorizationpolicylabel_authorizationpolicy_binding_result struct {
	Results []Count `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

func authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "labelname:"+key.Labelname)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddAuthorizationpolicylabelAuthorizationpolicyBinding(binding AuthorizationpolicylabelAuthorizationpolicyBinding) error {
	payload := add_authorizationpolicylabel_authorizationpolicy_binding_payload{
		binding,
	}

	return c.put("authorizationpolicylabel_authorizationpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountAuthorizationpolicylabelAuthorizationpolicyBinding() (int, error) {
	var results count_authorizationpolicylabel_authorizationpolicy_binding_result

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

func (c *NitroClient) CountAuthorizationpolicylabelAuthorizationpolicyBinding(id string) (int, error) {
	var results count_authorizationpolicylabel_authorizationpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuthorizationpolicylabelAuthorizationpolicyBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountAuthorizationpolicylabelAuthorizationpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListAuthorizationpolicylabelAuthorizationpolicyBinding() ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListAuthorizationpolicylabelAuthorizationpolicyBinding(id string) ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding_result

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (*AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding_result

	id, qs := authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicylabel_authorizationpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("authorizationpolicylabel_authorizationpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) error {
	id, qs := authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key)

	return c.delete("authorizationpolicylabel_authorizationpolicy_binding", id, qs)
}
