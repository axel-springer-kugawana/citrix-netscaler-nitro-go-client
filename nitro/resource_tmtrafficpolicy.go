package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmtrafficpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func tmtrafficpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type TmtrafficpolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_tmtrafficpolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_tmtrafficpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_tmtrafficpolicy_payload struct {
	Resource Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

type rename_tmtrafficpolicy_payload struct {
	Rename rename_tmtrafficpolicy `json:"tmtrafficpolicy"`
}

type unset_tmtrafficpolicy_payload struct {
	Unset TmtrafficpolicyUnset `json:"tmtrafficpolicy"`
}

type update_tmtrafficpolicy_payload struct {
	Update update_tmtrafficpolicy `json:"tmtrafficpolicy"`
}

type get_tmtrafficpolicy_result struct {
	Results []Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

type count_tmtrafficpolicy_result struct {
	Results []Count `json:"tmtrafficpolicy"`
}

func (c *NitroClient) AddTmtrafficpolicy(resource Tmtrafficpolicy) error {
	payload := add_tmtrafficpolicy_payload{
		resource,
	}

	return c.post("tmtrafficpolicy", "", nil, payload)
}

func (c *NitroClient) RenameTmtrafficpolicy(name string, newName string) error {
	payload := rename_tmtrafficpolicy_payload{
		rename_tmtrafficpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("tmtrafficpolicy", "", qs, payload)
}

func (c *NitroClient) CountTmtrafficpolicy() (int, error) {
	var results count_tmtrafficpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmtrafficpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTmtrafficpolicy(key string) (bool, error) {
	var results count_tmtrafficpolicy_result

	id, qs := tmtrafficpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTmtrafficpolicy() ([]Tmtrafficpolicy, error) {
	results := get_tmtrafficpolicy_result{}

	if err := c.get("tmtrafficpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTmtrafficpolicy(key string) (*Tmtrafficpolicy, error) {
	var results get_tmtrafficpolicy_result

	id, qs := tmtrafficpolicy_key_to_id_args(key)

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmtrafficpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmtrafficpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTmtrafficpolicy(key string) error {
	id, qs := tmtrafficpolicy_key_to_id_args(key)

	return c.delete("tmtrafficpolicy", id, qs)
}

func (c *NitroClient) UnsetTmtrafficpolicy(unset TmtrafficpolicyUnset) error {
	payload := unset_tmtrafficpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("tmtrafficpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateTmtrafficpolicy(resource Tmtrafficpolicy) error {
	payload := update_tmtrafficpolicy_payload{
		update_tmtrafficpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("tmtrafficpolicy", "", nil, payload)
}
