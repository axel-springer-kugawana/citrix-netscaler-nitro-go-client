package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Spilloveraction struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
}

func spilloveraction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type rename_spilloveraction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_spilloveraction_payload struct {
	Resource Spilloveraction `json:"spilloveraction"`
}

type rename_spilloveraction_payload struct {
	Rename rename_spilloveraction `json:"spilloveraction"`
}

type get_spilloveraction_result struct {
	Results []Spilloveraction `json:"spilloveraction"`
}

type count_spilloveraction_result struct {
	Results []Count `json:"spilloveraction"`
}

func (c *NitroClient) AddSpilloveraction(resource Spilloveraction) error {
	payload := add_spilloveraction_payload{
		resource,
	}

	return c.post("spilloveraction", "", nil, payload)
}

func (c *NitroClient) RenameSpilloveraction(name string, newName string) error {
	payload := rename_spilloveraction_payload{
		rename_spilloveraction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("spilloveraction", "", qs, payload)
}

func (c *NitroClient) CountSpilloveraction() (int, error) {
	var results count_spilloveraction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("spilloveraction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsSpilloveraction(key string) (bool, error) {
	var results count_spilloveraction_result

	id, qs := spilloveraction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("spilloveraction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListSpilloveraction() ([]Spilloveraction, error) {
	results := get_spilloveraction_result{}

	if err := c.get("spilloveraction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetSpilloveraction(key string) (*Spilloveraction, error) {
	var results get_spilloveraction_result

	id, qs := spilloveraction_key_to_id_args(key)

	if err := c.get("spilloveraction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one spilloveraction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("spilloveraction element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteSpilloveraction(key string) error {
	id, qs := spilloveraction_key_to_id_args(key)

	return c.delete("spilloveraction", id, qs)
}
