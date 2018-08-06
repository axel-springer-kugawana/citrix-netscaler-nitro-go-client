package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cspolicylabel struct {
	Labelname         string `json:"labelname"`
	Cspolicylabeltype string `json:"cspolicylabeltype,omitempty"`
}

func cspolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type rename_cspolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_cspolicylabel_payload struct {
	Resource Cspolicylabel `json:"cspolicylabel"`
}

type rename_cspolicylabel_payload struct {
	Rename rename_cspolicylabel `json:"cspolicylabel"`
}

type get_cspolicylabel_result struct {
	Results []Cspolicylabel `json:"cspolicylabel"`
}

type count_cspolicylabel_result struct {
	Results []Count `json:"cspolicylabel"`
}

func (c *NitroClient) AddCspolicylabel(resource Cspolicylabel) error {
	payload := add_cspolicylabel_payload{
		resource,
	}

	return c.post("cspolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameCspolicylabel(name string, newName string) error {
	payload := rename_cspolicylabel_payload{
		rename_cspolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cspolicylabel", "", qs, payload)
}

func (c *NitroClient) CountCspolicylabel() (int, error) {
	var results count_cspolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cspolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCspolicylabel(key string) (bool, error) {
	var results count_cspolicylabel_result

	id, qs := cspolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cspolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCspolicylabel() ([]Cspolicylabel, error) {
	results := get_cspolicylabel_result{}

	if err := c.get("cspolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCspolicylabel(key string) (*Cspolicylabel, error) {
	var results get_cspolicylabel_result

	id, qs := cspolicylabel_key_to_id_args(key)

	if err := c.get("cspolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cspolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cspolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCspolicylabel(key string) error {
	id, qs := cspolicylabel_key_to_id_args(key)

	return c.delete("cspolicylabel", id, qs)
}
