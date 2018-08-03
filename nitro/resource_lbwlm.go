package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbwlm struct {
	Wlmname   string `json:"wlmname"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Katimeout int    `json:"katimeout,string,omitempty"`
	Lbuid     string `json:"lbuid,omitempty"`
	Port      int    `json:"port,omitempty"`
}

func lbwlm_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type LbwlmUnset struct {
	Wlmname   string `json:"wlmname"`
	Katimeout bool   `json:"katimeout,string,omitempty"`
}

type update_lbwlm struct {
	Wlmname   string `json:"wlmname"`
	Katimeout int    `json:"katimeout,string,omitempty"`
}

type rename_lbwlm struct {
	Name    string `json:"wlmname"`
	Newname string `json:"newname"`
}

type add_lbwlm_payload struct {
	Resource Lbwlm `json:"lbwlm"`
}

type rename_lbwlm_payload struct {
	Rename rename_lbwlm `json:"lbwlm"`
}

type unset_lbwlm_payload struct {
	Unset LbwlmUnset `json:"lbwlm"`
}

type update_lbwlm_payload struct {
	Update update_lbwlm `json:"lbwlm"`
}

type get_lbwlm_result struct {
	Results []Lbwlm `json:"lbwlm"`
}

type count_lbwlm_result struct {
	Results []Count `json:"lbwlm"`
}

func (c *NitroClient) AddLbwlm(resource Lbwlm) error {
	payload := add_lbwlm_payload{
		resource,
	}

	return c.post("lbwlm", "", nil, payload)
}

func (c *NitroClient) RenameLbwlm(name string, newName string) error {
	payload := rename_lbwlm_payload{
		rename_lbwlm{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbwlm", "", qs, payload)
}

func (c *NitroClient) CountLbwlm() (int, error) {
	var results count_lbwlm_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbwlm", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbwlm(key string) (bool, error) {
	var results count_lbwlm_result

	id, qs := lbwlm_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbwlm", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbwlm() ([]Lbwlm, error) {
	results := get_lbwlm_result{}

	if err := c.get("lbwlm", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbwlm(key string) (*Lbwlm, error) {
	var results get_lbwlm_result

	id, qs := lbwlm_key_to_id_args(key)

	if err := c.get("lbwlm", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbwlm element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbwlm element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbwlm(key string) error {
	id, qs := lbwlm_key_to_id_args(key)

	return c.delete("lbwlm", id, qs)
}

func (c *NitroClient) UnsetLbwlm(unset LbwlmUnset) error {
	payload := unset_lbwlm_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("lbwlm", "", qs, payload)
}

func (c *NitroClient) UpdateLbwlm(resource Lbwlm) error {
	payload := update_lbwlm_payload{
		update_lbwlm{
			resource.Wlmname,
			resource.Katimeout,
		},
	}

	return c.put("lbwlm", "", nil, payload)
}
