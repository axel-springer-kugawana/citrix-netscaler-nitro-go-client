package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policystringmap struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

type PolicystringmapKey struct {
	Name string `json:"name"`
}

type PolicystringmapUnset struct {
	Name    string `json:"name"`
	Comment bool   `json:"comment,string,omitempty"`
}

type update_policystringmap struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
}

type rename_policystringmap struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_policystringmap_payload struct {
	Resource Policystringmap `json:"policystringmap"`
}

type rename_policystringmap_payload struct {
	Rename rename_policystringmap `json:"policystringmap"`
}

type unset_policystringmap_payload struct {
	Unset PolicystringmapUnset `json:"policystringmap"`
}

type update_policystringmap_payload struct {
	Update update_policystringmap `json:"policystringmap"`
}

type get_policystringmap_result struct {
	Results []Policystringmap `json:"policystringmap"`
}

type count_policystringmap_result struct {
	Results []Count `json:"policystringmap"`
}

func policystringmap_key_to_id_args(key PolicystringmapKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddPolicystringmap(resource Policystringmap) error {
	payload := add_policystringmap_payload{
		resource,
	}

	return c.post("policystringmap", "", nil, payload)
}

func (c *NitroClient) RenamePolicystringmap(name string, newName string) error {
	payload := rename_policystringmap_payload{
		rename_policystringmap{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("policystringmap", "", qs, payload)
}

func (c *NitroClient) CountPolicystringmap() (int, error) {
	var results count_policystringmap_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policystringmap", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicystringmap(key PolicystringmapKey) (bool, error) {
	var results count_policystringmap_result

	id, qs := policystringmap_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("policystringmap", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListPolicystringmap() ([]Policystringmap, error) {
	var results get_policystringmap_result

	if err := c.get("policystringmap", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicystringmap(key PolicystringmapKey) (*Policystringmap, error) {
	var results get_policystringmap_result

	id, qs := policystringmap_key_to_id_args(key)

	if err := c.get("policystringmap", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policystringmap element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policystringmap element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicystringmap(key PolicystringmapKey) error {
	id, qs := policystringmap_key_to_id_args(key)

	return c.delete("policystringmap", id, qs)
}

func (c *NitroClient) UnsetPolicystringmap(unset PolicystringmapUnset) error {
	payload := unset_policystringmap_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("policystringmap", "", qs, payload)
}

func (c *NitroClient) UpdatePolicystringmap(resource Policystringmap) error {
	payload := update_policystringmap_payload{
		update_policystringmap{
			resource.Name,
			resource.Comment,
		},
	}

	return c.put("policystringmap", "", nil, payload)
}
