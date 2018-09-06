package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbmetrictable struct {
	Metrictable string `json:"metrictable,omitempty"`
}

type LbmetrictableKey struct {
	Metrictable string
}

func (resource Lbmetrictable) ToKey() LbmetrictableKey {
	key := LbmetrictableKey{
		resource.Metrictable,
	}

	return key
}

func (key LbmetrictableKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Metrictable

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbmetrictableKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbmetrictableKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbmetrictable_payload struct {
	Resource Lbmetrictable `json:"lbmetrictable"`
}

func (c *NitroClient) AddLbmetrictable(resource Lbmetrictable) error {
	payload := add_lbmetrictable_payload{
		resource,
	}

	return c.post("lbmetrictable", "", nil, payload)
}

//      LIST

type list_lbmetrictable_result struct {
	Results []Lbmetrictable `json:"lbmetrictable"`
}

func (c *NitroClient) ListLbmetrictable() ([]Lbmetrictable, error) {
	results := list_lbmetrictable_result{}

	if err := c.get("lbmetrictable", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbmetrictable_result struct {
	Results []Lbmetrictable `json:"lbmetrictable"`
}

func (c *NitroClient) GetLbmetrictable(key LbmetrictableKey) (*Lbmetrictable, error) {
	var results get_lbmetrictable_result

	id, qs := key.to_id_args()

	if err := c.get("lbmetrictable", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmetrictable element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbmetrictable element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbmetrictable_result struct {
	Results []Count `json:"lbmetrictable"`
}

func (c *NitroClient) CountLbmetrictable() (int, error) {
	var results count_lbmetrictable_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmetrictable", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbmetrictable(key LbmetrictableKey) (bool, error) {
	var results count_lbmetrictable_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbmetrictable", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbmetrictable(key LbmetrictableKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbmetrictable", id, qs)
}

//      RENAME
//      TODO
