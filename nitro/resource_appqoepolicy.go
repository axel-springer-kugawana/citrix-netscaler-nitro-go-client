package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appqoepolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func appqoepolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppqoepolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

type update_appqoepolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_appqoepolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appqoepolicy_payload struct {
	Resource Appqoepolicy `json:"appqoepolicy"`
}

type rename_appqoepolicy_payload struct {
	Rename rename_appqoepolicy `json:"appqoepolicy"`
}

type unset_appqoepolicy_payload struct {
	Unset AppqoepolicyUnset `json:"appqoepolicy"`
}

type update_appqoepolicy_payload struct {
	Update update_appqoepolicy `json:"appqoepolicy"`
}

type get_appqoepolicy_result struct {
	Results []Appqoepolicy `json:"appqoepolicy"`
}

type count_appqoepolicy_result struct {
	Results []Count `json:"appqoepolicy"`
}

func (c *NitroClient) AddAppqoepolicy(resource Appqoepolicy) error {
	payload := add_appqoepolicy_payload{
		resource,
	}

	return c.post("appqoepolicy", "", nil, payload)
}

func (c *NitroClient) RenameAppqoepolicy(name string, newName string) error {
	payload := rename_appqoepolicy_payload{
		rename_appqoepolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appqoepolicy", "", qs, payload)
}

func (c *NitroClient) CountAppqoepolicy() (int, error) {
	var results count_appqoepolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appqoepolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppqoepolicy(key string) (bool, error) {
	var results count_appqoepolicy_result

	id, qs := appqoepolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appqoepolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppqoepolicy() ([]Appqoepolicy, error) {
	results := get_appqoepolicy_result{}

	if err := c.get("appqoepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppqoepolicy(key string) (*Appqoepolicy, error) {
	var results get_appqoepolicy_result

	id, qs := appqoepolicy_key_to_id_args(key)

	if err := c.get("appqoepolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appqoepolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appqoepolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppqoepolicy(key string) error {
	id, qs := appqoepolicy_key_to_id_args(key)

	return c.delete("appqoepolicy", id, qs)
}

func (c *NitroClient) UnsetAppqoepolicy(unset AppqoepolicyUnset) error {
	payload := unset_appqoepolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appqoepolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAppqoepolicy(resource Appqoepolicy) error {
	payload := update_appqoepolicy_payload{
		update_appqoepolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("appqoepolicy", "", nil, payload)
}
