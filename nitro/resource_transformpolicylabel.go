package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformpolicylabel struct {
	Labelname       string `json:"labelname"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type TransformpolicylabelKey struct {
	Labelname string `json:"labelname"`
}

type rename_transformpolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_transformpolicylabel_payload struct {
	Resource Transformpolicylabel `json:"transformpolicylabel"`
}

type rename_transformpolicylabel_payload struct {
	Rename rename_transformpolicylabel `json:"transformpolicylabel"`
}

type get_transformpolicylabel_result struct {
	Results []Transformpolicylabel `json:"transformpolicylabel"`
}

type count_transformpolicylabel_result struct {
	Results []Count `json:"transformpolicylabel"`
}

func transformpolicylabel_key_to_id_args(key TransformpolicylabelKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Labelname, qs
}

func (c *NitroClient) AddTransformpolicylabel(resource Transformpolicylabel) error {
	payload := add_transformpolicylabel_payload{
		resource,
	}

	return c.post("transformpolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameTransformpolicylabel(name string, newName string) error {
	payload := rename_transformpolicylabel_payload{
		rename_transformpolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("transformpolicylabel", "", qs, payload)
}

func (c *NitroClient) CountTransformpolicylabel() (int, error) {
	var results count_transformpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTransformpolicylabel(key TransformpolicylabelKey) (bool, error) {
	var results count_transformpolicylabel_result

	id, qs := transformpolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("transformpolicylabel", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTransformpolicylabel() ([]Transformpolicylabel, error) {
	var results get_transformpolicylabel_result

	if err := c.get("transformpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTransformpolicylabel(key TransformpolicylabelKey) (*Transformpolicylabel, error) {
	var results get_transformpolicylabel_result

	id, qs := transformpolicylabel_key_to_id_args(key)

	if err := c.get("transformpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformpolicylabel element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("transformpolicylabel element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTransformpolicylabel(key TransformpolicylabelKey) error {
	id, qs := transformpolicylabel_key_to_id_args(key)

	return c.delete("transformpolicylabel", id, qs)
}
