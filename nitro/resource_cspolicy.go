package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cspolicy struct {
	Policyname string `json:"policyname"`
	Action     string `json:"action,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Logaction  string `json:"logaction,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Url        string `json:"url,omitempty"`
}

func cspolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type CspolicyUnset struct {
	Policyname string `json:"policyname"`
	Url        bool   `json:"url,string,omitempty"`
	Rule       bool   `json:"rule,string,omitempty"`
	Domain     bool   `json:"domain,string,omitempty"`
	Action     bool   `json:"action,string,omitempty"`
	Logaction  bool   `json:"logaction,string,omitempty"`
}

type update_cspolicy struct {
	Policyname string `json:"policyname"`
	Url        string `json:"url,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Action     string `json:"action,omitempty"`
	Logaction  string `json:"logaction,omitempty"`
}

type rename_cspolicy struct {
	Name    string `json:"policyname"`
	Newname string `json:"newname"`
}

type add_cspolicy_payload struct {
	Resource Cspolicy `json:"cspolicy"`
}

type rename_cspolicy_payload struct {
	Rename rename_cspolicy `json:"cspolicy"`
}

type unset_cspolicy_payload struct {
	Unset CspolicyUnset `json:"cspolicy"`
}

type update_cspolicy_payload struct {
	Update update_cspolicy `json:"cspolicy"`
}

type get_cspolicy_result struct {
	Results []Cspolicy `json:"cspolicy"`
}

type count_cspolicy_result struct {
	Results []Count `json:"cspolicy"`
}

func (c *NitroClient) AddCspolicy(resource Cspolicy) error {
	payload := add_cspolicy_payload{
		resource,
	}

	return c.post("cspolicy", "", nil, payload)
}

func (c *NitroClient) RenameCspolicy(name string, newName string) error {
	payload := rename_cspolicy_payload{
		rename_cspolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cspolicy", "", qs, payload)
}

func (c *NitroClient) CountCspolicy() (int, error) {
	var results count_cspolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cspolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCspolicy(key string) (bool, error) {
	var results count_cspolicy_result

	id, qs := cspolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cspolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCspolicy() ([]Cspolicy, error) {
	results := get_cspolicy_result{}

	if err := c.get("cspolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCspolicy(key string) (*Cspolicy, error) {
	var results get_cspolicy_result

	id, qs := cspolicy_key_to_id_args(key)

	if err := c.get("cspolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cspolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cspolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCspolicy(key string) error {
	id, qs := cspolicy_key_to_id_args(key)

	return c.delete("cspolicy", id, qs)
}

func (c *NitroClient) UnsetCspolicy(unset CspolicyUnset) error {
	payload := unset_cspolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("cspolicy", "", qs, payload)
}

func (c *NitroClient) UpdateCspolicy(resource Cspolicy) error {
	payload := update_cspolicy_payload{
		update_cspolicy{
			resource.Policyname,
			resource.Url,
			resource.Rule,
			resource.Domain,
			resource.Action,
			resource.Logaction,
		},
	}

	return c.put("cspolicy", "", nil, payload)
}
