package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Feopolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type FeopolicyKey struct {
	Name string `json:"name"`
}

type FeopolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_feopolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_feopolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_feopolicy_payload struct {
	Resource Feopolicy `json:"feopolicy"`
}

type rename_feopolicy_payload struct {
	Rename rename_feopolicy `json:"feopolicy"`
}

type unset_feopolicy_payload struct {
	Unset FeopolicyUnset `json:"feopolicy"`
}

type update_feopolicy_payload struct {
	Update update_feopolicy `json:"feopolicy"`
}

type get_feopolicy_result struct {
	Results []Feopolicy `json:"feopolicy"`
}

type count_feopolicy_result struct {
	Results []Count `json:"feopolicy"`
}

func feopolicy_key_to_id_args(key FeopolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddFeopolicy(resource Feopolicy) error {
	payload := add_feopolicy_payload{
		resource,
	}

	return c.post("feopolicy", "", nil, payload)
}

func (c *NitroClient) RenameFeopolicy(name string, newName string) error {
	payload := rename_feopolicy_payload{
		rename_feopolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("feopolicy", "", qs, payload)
}

func (c *NitroClient) CountFeopolicy() (int, error) {
	var results count_feopolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("feopolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsFeopolicy(key FeopolicyKey) (bool, error) {
	var results count_feopolicy_result

	id, qs := feopolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("feopolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListFeopolicy() ([]Feopolicy, error) {
	var results get_feopolicy_result

	if err := c.get("feopolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetFeopolicy(key FeopolicyKey) (*Feopolicy, error) {
	var results get_feopolicy_result

	id, qs := feopolicy_key_to_id_args(key)

	if err := c.get("feopolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one feopolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("feopolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteFeopolicy(key FeopolicyKey) error {
	id, qs := feopolicy_key_to_id_args(key)

	return c.delete("feopolicy", id, qs)
}

func (c *NitroClient) UnsetFeopolicy(unset FeopolicyUnset) error {
	payload := unset_feopolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("feopolicy", "", qs, payload)
}

func (c *NitroClient) UpdateFeopolicy(resource Feopolicy) error {
	payload := update_feopolicy_payload{
		update_feopolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("feopolicy", "", nil, payload)
}
