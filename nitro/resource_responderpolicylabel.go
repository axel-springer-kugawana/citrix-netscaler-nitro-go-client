package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Responderpolicylabel struct {
	Labelname       string `json:"labelname"`
	Comment         string `json:"comment,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

func responderpolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type rename_responderpolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_responderpolicylabel_payload struct {
	Resource Responderpolicylabel `json:"responderpolicylabel"`
}

type rename_responderpolicylabel_payload struct {
	Rename rename_responderpolicylabel `json:"responderpolicylabel"`
}

type get_responderpolicylabel_result struct {
	Results []Responderpolicylabel `json:"responderpolicylabel"`
}

type count_responderpolicylabel_result struct {
	Results []Count `json:"responderpolicylabel"`
}

func (c *NitroClient) AddResponderpolicylabel(resource Responderpolicylabel) error {
	payload := add_responderpolicylabel_payload{
		resource,
	}

	return c.post("responderpolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameResponderpolicylabel(name string, newName string) error {
	payload := rename_responderpolicylabel_payload{
		rename_responderpolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("responderpolicylabel", "", qs, payload)
}

func (c *NitroClient) CountResponderpolicylabel() (int, error) {
	var results count_responderpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("responderpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsResponderpolicylabel(key string) (bool, error) {
	var results count_responderpolicylabel_result

	id, qs := responderpolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("responderpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListResponderpolicylabel() ([]Responderpolicylabel, error) {
	results := get_responderpolicylabel_result{}

	if err := c.get("responderpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetResponderpolicylabel(key string) (*Responderpolicylabel, error) {
	var results get_responderpolicylabel_result

	id, qs := responderpolicylabel_key_to_id_args(key)

	if err := c.get("responderpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one responderpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("responderpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteResponderpolicylabel(key string) error {
	id, qs := responderpolicylabel_key_to_id_args(key)

	return c.delete("responderpolicylabel", id, qs)
}
