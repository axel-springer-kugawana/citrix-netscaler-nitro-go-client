package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmppolicylabel struct {
	Labelname string `json:"labelname"`
	Type      string `json:"type,omitempty"`
}

func cmppolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type rename_cmppolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_cmppolicylabel_payload struct {
	Resource Cmppolicylabel `json:"cmppolicylabel"`
}

type rename_cmppolicylabel_payload struct {
	Rename rename_cmppolicylabel `json:"cmppolicylabel"`
}

type get_cmppolicylabel_result struct {
	Results []Cmppolicylabel `json:"cmppolicylabel"`
}

type count_cmppolicylabel_result struct {
	Results []Count `json:"cmppolicylabel"`
}

func (c *NitroClient) AddCmppolicylabel(resource Cmppolicylabel) error {
	payload := add_cmppolicylabel_payload{
		resource,
	}

	return c.post("cmppolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameCmppolicylabel(name string, newName string) error {
	payload := rename_cmppolicylabel_payload{
		rename_cmppolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cmppolicylabel", "", qs, payload)
}

func (c *NitroClient) CountCmppolicylabel() (int, error) {
	var results count_cmppolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cmppolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCmppolicylabel(key string) (bool, error) {
	var results count_cmppolicylabel_result

	id, qs := cmppolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cmppolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCmppolicylabel() ([]Cmppolicylabel, error) {
	results := get_cmppolicylabel_result{}

	if err := c.get("cmppolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCmppolicylabel(key string) (*Cmppolicylabel, error) {
	var results get_cmppolicylabel_result

	id, qs := cmppolicylabel_key_to_id_args(key)

	if err := c.get("cmppolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmppolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cmppolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCmppolicylabel(key string) error {
	id, qs := cmppolicylabel_key_to_id_args(key)

	return c.delete("cmppolicylabel", id, qs)
}
