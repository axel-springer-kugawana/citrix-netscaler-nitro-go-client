package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Authorizationpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuthorizationpolicyKey struct {
	Name string
}

func (resource Authorizationpolicy) ToKey() AuthorizationpolicyKey {
	key := AuthorizationpolicyKey{
		resource.Name,
	}

	return key
}

func (key AuthorizationpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AuthorizationpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuthorizationpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_authorizationpolicy_payload struct {
	Resource Authorizationpolicy `json:"authorizationpolicy"`
}

func (c *NitroClient) AddAuthorizationpolicy(resource Authorizationpolicy) error {
	payload := add_authorizationpolicy_payload{
		resource,
	}

	return c.post("authorizationpolicy", "", nil, payload)
}

//      LIST

type list_authorizationpolicy_result struct {
	Results []Authorizationpolicy `json:"authorizationpolicy"`
}

func (c *NitroClient) ListAuthorizationpolicy() ([]Authorizationpolicy, error) {
	results := list_authorizationpolicy_result{}

	if err := c.get("authorizationpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_authorizationpolicy_result struct {
	Results []Authorizationpolicy `json:"authorizationpolicy"`
}

func (c *NitroClient) GetAuthorizationpolicy(key AuthorizationpolicyKey) (*Authorizationpolicy, error) {
	var results get_authorizationpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("authorizationpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("authorizationpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_authorizationpolicy_result struct {
	Results []Count `json:"authorizationpolicy"`
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

//      EXISTS

func (c *NitroClient) ExistsAuthorizationpolicy(key AuthorizationpolicyKey) (bool, error) {
	var results count_authorizationpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("authorizationpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuthorizationpolicy(key AuthorizationpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("authorizationpolicy", id, qs)
}

//      UPDATE

type AuthorizationpolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Authorizationpolicy) ToUpdate() AuthorizationpolicyUpdate {
	update := AuthorizationpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_authorizationpolicy_payload struct {
	Update AuthorizationpolicyUpdate `json:"authorizationpolicy"`
}

func (c *NitroClient) UpdateAuthorizationpolicy(update AuthorizationpolicyUpdate) error {
	payload := update_authorizationpolicy_payload{
		update,
	}

	return c.put("authorizationpolicy", "", nil, payload)
}

//      UNSET

type AuthorizationpolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Authorizationpolicy) ToUnset() AuthorizationpolicyUnset {
	unset := AuthorizationpolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_authorizationpolicy_payload struct {
	Unset AuthorizationpolicyUnset `json:"authorizationpolicy"`
}

func (c *NitroClient) UnsetAuthorizationpolicy(unset AuthorizationpolicyUnset) error {
	payload := unset_authorizationpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("authorizationpolicy", "", qs, payload)
}

//      RENAME
//      TODO
