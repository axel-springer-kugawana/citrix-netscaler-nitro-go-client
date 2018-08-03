package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Rewritepolicylabel struct {
	Labelname string `json:"labelname"`
	Comment   string `json:"comment,omitempty"`
	Transform string `json:"transform,omitempty"`
}

type RewritepolicylabelKey struct {
	Labelname string `json:"labelname"`
}

type rename_rewritepolicylabel struct {
	Name    string `json:"labelname"`
	Newname string `json:"newname"`
}

type add_rewritepolicylabel_payload struct {
	Resource Rewritepolicylabel `json:"rewritepolicylabel"`
}

type rename_rewritepolicylabel_payload struct {
	Rename rename_rewritepolicylabel `json:"rewritepolicylabel"`
}

type get_rewritepolicylabel_result struct {
	Results []Rewritepolicylabel `json:"rewritepolicylabel"`
}

type count_rewritepolicylabel_result struct {
	Results []Count `json:"rewritepolicylabel"`
}

func rewritepolicylabel_key_to_id_args(key RewritepolicylabelKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Labelname, qs
}

func (c *NitroClient) AddRewritepolicylabel(resource Rewritepolicylabel) error {
	payload := add_rewritepolicylabel_payload{
		resource,
	}

	return c.post("rewritepolicylabel", "", nil, payload)
}

func (c *NitroClient) RenameRewritepolicylabel(name string, newName string) error {
	payload := rename_rewritepolicylabel_payload{
		rename_rewritepolicylabel{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("rewritepolicylabel", "", qs, payload)
}

func (c *NitroClient) CountRewritepolicylabel() (int, error) {
	var results count_rewritepolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("rewritepolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsRewritepolicylabel(key RewritepolicylabelKey) (bool, error) {
	var results count_rewritepolicylabel_result

	id, qs := rewritepolicylabel_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("rewritepolicylabel", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListRewritepolicylabel() ([]Rewritepolicylabel, error) {
	var results get_rewritepolicylabel_result

	if err := c.get("rewritepolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetRewritepolicylabel(key RewritepolicylabelKey) (*Rewritepolicylabel, error) {
	var results get_rewritepolicylabel_result

	id, qs := rewritepolicylabel_key_to_id_args(key)

	if err := c.get("rewritepolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one rewritepolicylabel element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("rewritepolicylabel element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteRewritepolicylabel(key RewritepolicylabelKey) error {
	id, qs := rewritepolicylabel_key_to_id_args(key)

	return c.delete("rewritepolicylabel", id, qs)
}
