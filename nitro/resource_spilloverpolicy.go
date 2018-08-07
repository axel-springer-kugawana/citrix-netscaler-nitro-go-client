package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Spilloverpolicy struct {
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
	Rule    string `json:"rule,omitempty"`
}

type SpilloverpolicyKey struct {
	Name string
}

func (resource Spilloverpolicy) ToKey() SpilloverpolicyKey {
	key := SpilloverpolicyKey{
		resource.Name,
	}

	return key
}

func (key SpilloverpolicyKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_spilloverpolicy_payload struct {
	Resource Spilloverpolicy `json:"spilloverpolicy"`
}

func (c *NitroClient) AddSpilloverpolicy(resource Spilloverpolicy) error {
	payload := add_spilloverpolicy_payload{
		resource,
	}

	return c.post("spilloverpolicy", "", nil, payload)
}

//      LIST

type list_spilloverpolicy_result struct {
	Results []Spilloverpolicy `json:"spilloverpolicy"`
}

func (c *NitroClient) ListSpilloverpolicy() ([]Spilloverpolicy, error) {
	results := list_spilloverpolicy_result{}

	if err := c.get("spilloverpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_spilloverpolicy_result struct {
	Results []Spilloverpolicy `json:"spilloverpolicy"`
}

func (c *NitroClient) GetSpilloverpolicy(key SpilloverpolicyKey) (*Spilloverpolicy, error) {
	var results get_spilloverpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("spilloverpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one spilloverpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("spilloverpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_spilloverpolicy_result struct {
	Results []Count `json:"spilloverpolicy"`
}

func (c *NitroClient) CountSpilloverpolicy() (int, error) {
	var results count_spilloverpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("spilloverpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsSpilloverpolicy(key SpilloverpolicyKey) (bool, error) {
	var results count_spilloverpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("spilloverpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteSpilloverpolicy(key SpilloverpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("spilloverpolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
