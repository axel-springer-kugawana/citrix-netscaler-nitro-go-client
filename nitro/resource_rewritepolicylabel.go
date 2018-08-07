package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Rewritepolicylabel struct {
	Comment   string `json:"comment,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Transform string `json:"transform,omitempty"`
}

type RewritepolicylabelKey struct {
	Labelname string
}

func (resource Rewritepolicylabel) ToKey() RewritepolicylabelKey {
	key := RewritepolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key RewritepolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key RewritepolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key RewritepolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_rewritepolicylabel_payload struct {
	Resource Rewritepolicylabel `json:"rewritepolicylabel"`
}

func (c *NitroClient) AddRewritepolicylabel(resource Rewritepolicylabel) error {
	payload := add_rewritepolicylabel_payload{
		resource,
	}

	return c.post("rewritepolicylabel", "", nil, payload)
}

//      LIST

type list_rewritepolicylabel_result struct {
	Results []Rewritepolicylabel `json:"rewritepolicylabel"`
}

func (c *NitroClient) ListRewritepolicylabel() ([]Rewritepolicylabel, error) {
	results := list_rewritepolicylabel_result{}

	if err := c.get("rewritepolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_rewritepolicylabel_result struct {
	Results []Rewritepolicylabel `json:"rewritepolicylabel"`
}

func (c *NitroClient) GetRewritepolicylabel(key RewritepolicylabelKey) (*Rewritepolicylabel, error) {
	var results get_rewritepolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("rewritepolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one rewritepolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("rewritepolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_rewritepolicylabel_result struct {
	Results []Count `json:"rewritepolicylabel"`
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

//      EXISTS

func (c *NitroClient) ExistsRewritepolicylabel(key RewritepolicylabelKey) (bool, error) {
	var results count_rewritepolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("rewritepolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteRewritepolicylabel(key RewritepolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("rewritepolicylabel", id, qs)
}

//      RENAME
//      TODO
