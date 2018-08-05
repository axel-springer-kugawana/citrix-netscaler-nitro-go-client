package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmsessionpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func tmsessionpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type TmsessionpolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_tmsessionpolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_tmsessionpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_tmsessionpolicy_payload struct {
	Resource Tmsessionpolicy `json:"tmsessionpolicy"`
}

type rename_tmsessionpolicy_payload struct {
	Rename rename_tmsessionpolicy `json:"tmsessionpolicy"`
}

type unset_tmsessionpolicy_payload struct {
	Unset TmsessionpolicyUnset `json:"tmsessionpolicy"`
}

type update_tmsessionpolicy_payload struct {
	Update update_tmsessionpolicy `json:"tmsessionpolicy"`
}

type get_tmsessionpolicy_result struct {
	Results []Tmsessionpolicy `json:"tmsessionpolicy"`
}

type count_tmsessionpolicy_result struct {
	Results []Count `json:"tmsessionpolicy"`
}

func (c *NitroClient) AddTmsessionpolicy(resource Tmsessionpolicy) error {
	payload := add_tmsessionpolicy_payload{
		resource,
	}

	return c.post("tmsessionpolicy", "", nil, payload)
}

func (c *NitroClient) RenameTmsessionpolicy(name string, newName string) error {
	payload := rename_tmsessionpolicy_payload{
		rename_tmsessionpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("tmsessionpolicy", "", qs, payload)
}

func (c *NitroClient) CountTmsessionpolicy() (int, error) {
	var results count_tmsessionpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmsessionpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTmsessionpolicy(key string) (bool, error) {
	var results count_tmsessionpolicy_result

	id, qs := tmsessionpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("tmsessionpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTmsessionpolicy() ([]Tmsessionpolicy, error) {
	results := get_tmsessionpolicy_result{}

	if err := c.get("tmsessionpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTmsessionpolicy(key string) (*Tmsessionpolicy, error) {
	var results get_tmsessionpolicy_result

	id, qs := tmsessionpolicy_key_to_id_args(key)

	if err := c.get("tmsessionpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmsessionpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmsessionpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTmsessionpolicy(key string) error {
	id, qs := tmsessionpolicy_key_to_id_args(key)

	return c.delete("tmsessionpolicy", id, qs)
}

func (c *NitroClient) UnsetTmsessionpolicy(unset TmsessionpolicyUnset) error {
	payload := unset_tmsessionpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("tmsessionpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateTmsessionpolicy(resource Tmsessionpolicy) error {
	payload := update_tmsessionpolicy_payload{
		update_tmsessionpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("tmsessionpolicy", "", nil, payload)
}
