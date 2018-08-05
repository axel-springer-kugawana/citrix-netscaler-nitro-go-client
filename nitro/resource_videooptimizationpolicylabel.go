package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationpolicylabel struct {
	Labelname       string `json:"labelname"`
	Comment         string `json:"comment,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

func videooptimizationpolicylabel_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type rename_videooptimizationpolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_videooptimizationpolicylabel_payload struct {
	Resource Videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

type rename_videooptimizationpolicylabel_payload struct {
	Rename rename_videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

type get_videooptimizationpolicylabel_result struct {
	Results []Videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

type count_videooptimizationpolicylabel_result struct {
	Results []Count `json:"videooptimizationpolicylabel"`
}

func (c *NitroClient) AddVideooptimizationpolicylabel(resource Videooptimizationpolicylabel) error {
	payload := add_videooptimizationpolicylabel_payload{
		resource,
	}

	return c.post("videooptimizationpolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameVideooptimizationpolicylabel(name string, newName string) error {
	payload := rename_videooptimizationpolicylabel_payload{
		rename_videooptimizationpolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("videooptimizationpolicylabel", "", qs, payload)
}

func (c *NitroClient) CountVideooptimizationpolicylabel() (int, error) {
	var results count_videooptimizationpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("videooptimizationpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsVideooptimizationpolicylabel(key string) (bool, error) {
	var results count_videooptimizationpolicylabel_result

	id, qs := videooptimizationpolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("videooptimizationpolicylabel", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListVideooptimizationpolicylabel() ([]Videooptimizationpolicylabel, error) {
	results := get_videooptimizationpolicylabel_result{}

	if err := c.get("videooptimizationpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetVideooptimizationpolicylabel(key string) (*Videooptimizationpolicylabel, error) {
	var results get_videooptimizationpolicylabel_result

	id, qs := videooptimizationpolicylabel_key_to_id_args(key)

	if err := c.get("videooptimizationpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one videooptimizationpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("videooptimizationpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteVideooptimizationpolicylabel(key string) error {
	id, qs := videooptimizationpolicylabel_key_to_id_args(key)

	return c.delete("videooptimizationpolicylabel", id, qs)
}
