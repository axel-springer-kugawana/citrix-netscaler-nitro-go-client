package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationpolicylabel struct {
	Comment         string `json:"comment,omitempty"`
	Labelname       string `json:"labelname,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type VideooptimizationpolicylabelKey struct {
	Labelname string
}

func (resource Videooptimizationpolicylabel) ToKey() VideooptimizationpolicylabelKey {
	key := VideooptimizationpolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key VideooptimizationpolicylabelKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_videooptimizationpolicylabel_payload struct {
	Resource Videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

func (c *NitroClient) AddVideooptimizationpolicylabel(resource Videooptimizationpolicylabel) error {
	payload := add_videooptimizationpolicylabel_payload{
		resource,
	}

	return c.post("videooptimizationpolicylabel", "", nil, payload)
}

//      LIST

type list_videooptimizationpolicylabel_result struct {
	Results []Videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

func (c *NitroClient) ListVideooptimizationpolicylabel() ([]Videooptimizationpolicylabel, error) {
	results := list_videooptimizationpolicylabel_result{}

	if err := c.get("videooptimizationpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_videooptimizationpolicylabel_result struct {
	Results []Videooptimizationpolicylabel `json:"videooptimizationpolicylabel"`
}

func (c *NitroClient) GetVideooptimizationpolicylabel(key VideooptimizationpolicylabelKey) (*Videooptimizationpolicylabel, error) {
	var results get_videooptimizationpolicylabel_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_videooptimizationpolicylabel_result struct {
	Results []Count `json:"videooptimizationpolicylabel"`
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

//      EXISTS

func (c *NitroClient) ExistsVideooptimizationpolicylabel(key VideooptimizationpolicylabelKey) (bool, error) {
	var results count_videooptimizationpolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("videooptimizationpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteVideooptimizationpolicylabel(key VideooptimizationpolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("videooptimizationpolicylabel", id, qs)
}

//      RENAME
//      TODO
