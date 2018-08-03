package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Authorizationpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuthorizationpolicyKey struct {
	Name string `json:"name"`
}

type AuthorizationpolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_authorizationpolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_authorizationpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_authorizationpolicy_payload struct {
	Resource Authorizationpolicy `json:"authorizationpolicy"`
}

type rename_authorizationpolicy_payload struct {
	Rename rename_authorizationpolicy `json:"authorizationpolicy"`
}

type unset_authorizationpolicy_payload struct {
	Unset AuthorizationpolicyUnset `json:"authorizationpolicy"`
}

type update_authorizationpolicy_payload struct {
	Update update_authorizationpolicy `json:"authorizationpolicy"`
}

type get_authorizationpolicy_result struct {
	Results []Authorizationpolicy `json:"authorizationpolicy"`
}

type count_authorizationpolicy_result struct {
	Results []Count `json:"authorizationpolicy"`
}

func authorizationpolicy_key_to_id_args(key AuthorizationpolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddAuthorizationpolicy(resource Authorizationpolicy) error {
	payload := add_authorizationpolicy_payload{
		resource,
	}

	return c.post("authorizationpolicy", "", nil, payload)
}

func (c *NitroClient) RenameAuthorizationpolicy(name string, newName string) error {
	payload := rename_authorizationpolicy_payload{
		rename_authorizationpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("authorizationpolicy", "", qs, payload)
}

func (c *NitroClient) CountAuthorizationpolicy() (int, error) {
	var results count_authorizationpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("authorizationpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuthorizationpolicy(key AuthorizationpolicyKey) (bool, error) {
	var results count_authorizationpolicy_result

	id, qs := authorizationpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("authorizationpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuthorizationpolicy() ([]Authorizationpolicy, error) {
	var results get_authorizationpolicy_result

	if err := c.get("authorizationpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuthorizationpolicy(key AuthorizationpolicyKey) (*Authorizationpolicy, error) {
	var results get_authorizationpolicy_result

	id, qs := authorizationpolicy_key_to_id_args(key)

	if err := c.get("authorizationpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("authorizationpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuthorizationpolicy(key AuthorizationpolicyKey) error {
	id, qs := authorizationpolicy_key_to_id_args(key)

	return c.delete("authorizationpolicy", id, qs)
}

func (c *NitroClient) UnsetAuthorizationpolicy(unset AuthorizationpolicyUnset) error {
	payload := unset_authorizationpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("authorizationpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAuthorizationpolicy(resource Authorizationpolicy) error {
	payload := update_authorizationpolicy_payload{
		update_authorizationpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("authorizationpolicy", "", nil, payload)
}
