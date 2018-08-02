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

type get_authorizationpolicylabel_authorizationpolicy_binding struct {
	Results []AuthorizationpolicylabelAuthorizationpolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
}

type add_authorizationpolicylabel_authorizationpolicy_binding_payload struct {
	authorizationpolicylabel_authorizationpolicy_binding AuthorizationpolicylabelAuthorizationpolicyBinding
}

func authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "labelname:"+key.Labelname)
	args = append(args, "policyname:"+key.Policyname)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddAuthorizationpolicylabelAuthorizationpolicyBinding(binding AuthorizationpolicylabelAuthorizationpolicyBinding) error {
	payload := add_authorizationpolicylabel_authorizationpolicy_binding_payload{
		binding,
	}

	return c.put("authorizationpolicylabel_authorizationpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding

	id, args := authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListAuthorizationpolicylabelAuthorizationpolicyBinding() ([]AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuthorizationpolicylabelAuthorizationpolicyBinding(key AuthorizationpolicylabelAuthorizationpolicyBindingKey) (*AuthorizationpolicylabelAuthorizationpolicyBinding, error) {
	var results get_authorizationpolicylabel_authorizationpolicy_binding

	id, args := authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key)

	if err := c.get("authorizationpolicylabel_authorizationpolicy_binding", id, "", args, &results); err != nil {
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
	id, args := authorizationpolicylabel_authorizationpolicy_binding_key_to_id_args(key)

	return c.delete("authorizationpolicylabel_authorizationpolicy_binding", id, "", args)
}
