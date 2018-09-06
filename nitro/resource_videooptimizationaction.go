package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationaction struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Type    string `json:"type,omitempty"`
}

type VideooptimizationactionKey struct {
	Name string
}

func (resource Videooptimizationaction) ToKey() VideooptimizationactionKey {
	key := VideooptimizationactionKey{
		resource.Name,
	}

	return key
}

func (key VideooptimizationactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key VideooptimizationactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key VideooptimizationactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_videooptimizationaction_payload struct {
	Resource Videooptimizationaction `json:"videooptimizationaction"`
}

func (c *NitroClient) AddVideooptimizationaction(resource Videooptimizationaction) error {
	payload := add_videooptimizationaction_payload{
		resource,
	}

	return c.post("videooptimizationaction", "", nil, payload)
}

//      LIST

type list_videooptimizationaction_result struct {
	Results []Videooptimizationaction `json:"videooptimizationaction"`
}

func (c *NitroClient) ListVideooptimizationaction() ([]Videooptimizationaction, error) {
	results := list_videooptimizationaction_result{}

	if err := c.get("videooptimizationaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_videooptimizationaction_result struct {
	Results []Videooptimizationaction `json:"videooptimizationaction"`
}

func (c *NitroClient) GetVideooptimizationaction(key VideooptimizationactionKey) (*Videooptimizationaction, error) {
	var results get_videooptimizationaction_result

	id, qs := key.to_id_args()

	if err := c.get("videooptimizationaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one videooptimizationaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("videooptimizationaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_videooptimizationaction_result struct {
	Results []Count `json:"videooptimizationaction"`
}

func (c *NitroClient) CountVideooptimizationaction() (int, error) {
	var results count_videooptimizationaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("videooptimizationaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsVideooptimizationaction(key VideooptimizationactionKey) (bool, error) {
	var results count_videooptimizationaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("videooptimizationaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteVideooptimizationaction(key VideooptimizationactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("videooptimizationaction", id, qs)
}

//      UPDATE

type VideooptimizationactionUpdate struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (resource Videooptimizationaction) ToUpdate() VideooptimizationactionUpdate {
	update := VideooptimizationactionUpdate{
		resource.Name,
		resource.Type,
		resource.Rate,
		resource.Comment,
	}

	return update
}

type update_videooptimizationaction_payload struct {
	Update VideooptimizationactionUpdate `json:"videooptimizationaction"`
}

func (c *NitroClient) UpdateVideooptimizationaction(update VideooptimizationactionUpdate) error {
	payload := update_videooptimizationaction_payload{
		update,
	}

	return c.put("videooptimizationaction", "", nil, payload)
}

//      UNSET

type VideooptimizationactionUnset struct {
	Name    string `json:"name,omitempty"`
	Type    bool   `json:"type,omitempty"`
	Rate    bool   `json:"rate,omitempty"`
	Comment bool   `json:"comment,omitempty"`
}

func (resource Videooptimizationaction) ToUnset() VideooptimizationactionUnset {
	unset := VideooptimizationactionUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_videooptimizationaction_payload struct {
	Unset VideooptimizationactionUnset `json:"videooptimizationaction"`
}

func (c *NitroClient) UnsetVideooptimizationaction(unset VideooptimizationactionUnset) error {
	payload := unset_videooptimizationaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("videooptimizationaction", "", qs, payload)
}

//      RENAME
//      TODO
