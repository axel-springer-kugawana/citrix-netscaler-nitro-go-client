package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnspolicy64 struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func dnspolicy64_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type Dnspolicy64Unset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_dnspolicy64 struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_dnspolicy64 struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_dnspolicy64_payload struct {
	Resource Dnspolicy64 `json:"dnspolicy64"`
}

type rename_dnspolicy64_payload struct {
	Rename rename_dnspolicy64 `json:"dnspolicy64"`
}

type unset_dnspolicy64_payload struct {
	Unset Dnspolicy64Unset `json:"dnspolicy64"`
}

type update_dnspolicy64_payload struct {
	Update update_dnspolicy64 `json:"dnspolicy64"`
}

type get_dnspolicy64_result struct {
	Results []Dnspolicy64 `json:"dnspolicy64"`
}

type count_dnspolicy64_result struct {
	Results []Count `json:"dnspolicy64"`
}

func (c *NitroClient) AddDnspolicy64(resource Dnspolicy64) error {
	payload := add_dnspolicy64_payload{
		resource,
	}

	return c.post("dnspolicy64", "", nil, payload)
}

func (c *NitroClient) RenameDnspolicy64(name string, newName string) error {
	payload := rename_dnspolicy64_payload{
		rename_dnspolicy64{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("dnspolicy64", "", qs, payload)
}

func (c *NitroClient) CountDnspolicy64() (int, error) {
	var results count_dnspolicy64_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dnspolicy64", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsDnspolicy64(key string) (bool, error) {
	var results count_dnspolicy64_result

	id, qs := dnspolicy64_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("dnspolicy64", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListDnspolicy64() ([]Dnspolicy64, error) {
	results := get_dnspolicy64_result{}

	if err := c.get("dnspolicy64", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetDnspolicy64(key string) (*Dnspolicy64, error) {
	var results get_dnspolicy64_result

	id, qs := dnspolicy64_key_to_id_args(key)

	if err := c.get("dnspolicy64", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dnspolicy64 element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dnspolicy64 element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteDnspolicy64(key string) error {
	id, qs := dnspolicy64_key_to_id_args(key)

	return c.delete("dnspolicy64", id, qs)
}

func (c *NitroClient) UnsetDnspolicy64(unset Dnspolicy64Unset) error {
	payload := unset_dnspolicy64_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("dnspolicy64", "", qs, payload)
}

func (c *NitroClient) UpdateDnspolicy64(resource Dnspolicy64) error {
	payload := update_dnspolicy64_payload{
		update_dnspolicy64{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("dnspolicy64", "", nil, payload)
}
