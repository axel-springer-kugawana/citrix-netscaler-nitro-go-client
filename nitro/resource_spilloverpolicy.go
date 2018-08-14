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

func (key SpilloverpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key SpilloverpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key SpilloverpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type SpilloverpolicyUpdate struct {
	Name    string `json:"name,omitempty"`
	Rule    string `json:"rule,omitempty"`
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (resource Spilloverpolicy) ToUpdate() SpilloverpolicyUpdate {
	update := SpilloverpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Comment,
	}

	return update
}

type update_spilloverpolicy_payload struct {
	Update SpilloverpolicyUpdate `json:"spilloverpolicy"`
}

func (c *NitroClient) UpdateSpilloverpolicy(update SpilloverpolicyUpdate) error {
	payload := update_spilloverpolicy_payload{
		update,
	}

	return c.put("spilloverpolicy", "", nil, payload)
}

//      UNSET

type SpilloverpolicyUnset struct {
	Name    string `json:"name,omitempty"`
	Rule    bool   `json:"rule,omitempty"`
	Action  bool   `json:"action,omitempty"`
	Comment bool   `json:"comment,omitempty"`
}

func (resource Spilloverpolicy) ToUnset() SpilloverpolicyUnset {
	unset := SpilloverpolicyUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_spilloverpolicy_payload struct {
	Unset SpilloverpolicyUnset `json:"spilloverpolicy"`
}

func (c *NitroClient) UnsetSpilloverpolicy(unset SpilloverpolicyUnset) error {
	payload := unset_spilloverpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("spilloverpolicy", "", qs, payload)
}

//      RENAME
//      TODO
