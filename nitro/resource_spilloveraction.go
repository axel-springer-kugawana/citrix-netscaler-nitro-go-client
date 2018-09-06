package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Spilloveraction struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
}

type SpilloveractionKey struct {
	Name string
}

func (resource Spilloveraction) ToKey() SpilloveractionKey {
	key := SpilloveractionKey{
		resource.Name,
	}

	return key
}

func (key SpilloveractionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key SpilloveractionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key SpilloveractionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_spilloveraction_payload struct {
	Resource Spilloveraction `json:"spilloveraction"`
}

func (c *NitroClient) AddSpilloveraction(resource Spilloveraction) error {
	payload := add_spilloveraction_payload{
		resource,
	}

	return c.post("spilloveraction", "", nil, payload)
}

//      LIST

type list_spilloveraction_result struct {
	Results []Spilloveraction `json:"spilloveraction"`
}

func (c *NitroClient) ListSpilloveraction() ([]Spilloveraction, error) {
	results := list_spilloveraction_result{}

	if err := c.get("spilloveraction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_spilloveraction_result struct {
	Results []Spilloveraction `json:"spilloveraction"`
}

func (c *NitroClient) GetSpilloveraction(key SpilloveractionKey) (*Spilloveraction, error) {
	var results get_spilloveraction_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_spilloveraction_result struct {
	Results []Count `json:"spilloveraction"`
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

//      EXISTS

func (c *NitroClient) ExistsSpilloveraction(key SpilloveractionKey) (bool, error) {
	var results count_spilloveraction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("spilloveraction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteSpilloveraction(key SpilloveractionKey) error {
	id, qs := key.to_id_args()

	return c.delete("spilloveraction", id, qs)
}

//      RENAME
//      TODO
