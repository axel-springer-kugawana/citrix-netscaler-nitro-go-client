package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Authorizationpolicylabel struct {
	Labelname string `json:"labelname"`
}

func authorizationpolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type rename_authorizationpolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_authorizationpolicylabel_payload struct {
	Resource Authorizationpolicylabel `json:"authorizationpolicylabel"`
}

type rename_authorizationpolicylabel_payload struct {
	Rename rename_authorizationpolicylabel `json:"authorizationpolicylabel"`
}

type get_authorizationpolicylabel_result struct {
	Results []Authorizationpolicylabel `json:"authorizationpolicylabel"`
}

type count_authorizationpolicylabel_result struct {
	Results []Count `json:"authorizationpolicylabel"`
}

func (c *NitroClient) AddAuthorizationpolicylabel(resource Authorizationpolicylabel) error {
	payload := add_authorizationpolicylabel_payload{
		resource,
	}

	return c.post("authorizationpolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameAuthorizationpolicylabel(name string, newName string) error {
	payload := rename_authorizationpolicylabel_payload{
		rename_authorizationpolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("authorizationpolicylabel", "", qs, payload)
}

func (c *NitroClient) CountAuthorizationpolicylabel() (int, error) {
	var results count_authorizationpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("authorizationpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuthorizationpolicylabel(key string) (bool, error) {
	var results count_authorizationpolicylabel_result

	id, qs := authorizationpolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("authorizationpolicylabel", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuthorizationpolicylabel() ([]Authorizationpolicylabel, error) {
	results := get_authorizationpolicylabel_result{}

	if err := c.get("authorizationpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuthorizationpolicylabel(key string) (*Authorizationpolicylabel, error) {
	var results get_authorizationpolicylabel_result

	id, qs := authorizationpolicylabel_key_to_id_args(key)

	if err := c.get("authorizationpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicylabel element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("authorizationpolicylabel element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuthorizationpolicylabel(key string) error {
	id, qs := authorizationpolicylabel_key_to_id_args(key)

	return c.delete("authorizationpolicylabel", id, qs)
}
