package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbmetrictable struct {
	Metrictable string `json:"metrictable"`
}

func lbmetrictable_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type rename_lbmetrictable struct {
	Name    string `json:"metrictable"`
	Newname string `json:"newname"`
}

type add_lbmetrictable_payload struct {
	Resource Lbmetrictable `json:"lbmetrictable"`
}

type rename_lbmetrictable_payload struct {
	Rename rename_lbmetrictable `json:"lbmetrictable"`
}

type get_lbmetrictable_result struct {
	Results []Lbmetrictable `json:"lbmetrictable"`
}

type count_lbmetrictable_result struct {
	Results []Count `json:"lbmetrictable"`
}

func (c *NitroClient) AddLbmetrictable(resource Lbmetrictable) error {
	payload := add_lbmetrictable_payload{
		resource,
	}

	return c.post("lbmetrictable", "", nil, payload)
}

func (c *NitroClient) RenameLbmetrictable(name string, newName string) error {
	payload := rename_lbmetrictable_payload{
		rename_lbmetrictable{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbmetrictable", "", qs, payload)
}

func (c *NitroClient) CountLbmetrictable() (int, error) {
	var results count_lbmetrictable_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmetrictable", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbmetrictable(key string) (bool, error) {
	var results count_lbmetrictable_result

	id, qs := lbmetrictable_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbmetrictable", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbmetrictable() ([]Lbmetrictable, error) {
	results := get_lbmetrictable_result{}

	if err := c.get("lbmetrictable", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmetrictable(key string) (*Lbmetrictable, error) {
	var results get_lbmetrictable_result

	id, qs := lbmetrictable_key_to_id_args(key)

	if err := c.get("lbmetrictable", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmetrictable element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbmetrictable element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmetrictable(key string) error {
	id, qs := lbmetrictable_key_to_id_args(key)

	return c.delete("lbmetrictable", id, qs)
}
