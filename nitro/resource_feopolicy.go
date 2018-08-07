package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Feopolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type FeopolicyKey struct {
	Name string
}

func (resource Feopolicy) ToKey() FeopolicyKey {
	key := FeopolicyKey{
		resource.Name,
	}

	return key
}

func (key FeopolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key FeopolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key FeopolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_feopolicy_payload struct {
	Resource Feopolicy `json:"feopolicy"`
}

func (c *NitroClient) AddFeopolicy(resource Feopolicy) error {
	payload := add_feopolicy_payload{
		resource,
	}

	return c.post("feopolicy", "", nil, payload)
}

//      LIST

type list_feopolicy_result struct {
	Results []Feopolicy `json:"feopolicy"`
}

func (c *NitroClient) ListFeopolicy() ([]Feopolicy, error) {
	results := list_feopolicy_result{}

	if err := c.get("feopolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_feopolicy_result struct {
	Results []Feopolicy `json:"feopolicy"`
}

func (c *NitroClient) GetFeopolicy(key FeopolicyKey) (*Feopolicy, error) {
	var results get_feopolicy_result

	id, qs := key.to_id_args()

	if err := c.get("feopolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one feopolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("feopolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_feopolicy_result struct {
	Results []Count `json:"feopolicy"`
}

func (c *NitroClient) CountFeopolicy() (int, error) {
	var results count_feopolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("feopolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsFeopolicy(key FeopolicyKey) (bool, error) {
	var results count_feopolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("feopolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteFeopolicy(key FeopolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("feopolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
