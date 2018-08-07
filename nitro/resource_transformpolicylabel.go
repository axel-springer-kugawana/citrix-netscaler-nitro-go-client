package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformpolicylabel struct {
	Labelname       string `json:"labelname,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type TransformpolicylabelKey struct {
	Labelname string
}

func (resource Transformpolicylabel) ToKey() TransformpolicylabelKey {
	key := TransformpolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key TransformpolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TransformpolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TransformpolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_transformpolicylabel_payload struct {
	Resource Transformpolicylabel `json:"transformpolicylabel"`
}

func (c *NitroClient) AddTransformpolicylabel(resource Transformpolicylabel) error {
	payload := add_transformpolicylabel_payload{
		resource,
	}

	return c.post("transformpolicylabel", "", nil, payload)
}

//      LIST

type list_transformpolicylabel_result struct {
	Results []Transformpolicylabel `json:"transformpolicylabel"`
}

func (c *NitroClient) ListTransformpolicylabel() ([]Transformpolicylabel, error) {
	results := list_transformpolicylabel_result{}

	if err := c.get("transformpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_transformpolicylabel_result struct {
	Results []Transformpolicylabel `json:"transformpolicylabel"`
}

func (c *NitroClient) GetTransformpolicylabel(key TransformpolicylabelKey) (*Transformpolicylabel, error) {
	var results get_transformpolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("transformpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("transformpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_transformpolicylabel_result struct {
	Results []Count `json:"transformpolicylabel"`
}

func (c *NitroClient) CountTransformpolicylabel() (int, error) {
	var results count_transformpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTransformpolicylabel(key TransformpolicylabelKey) (bool, error) {
	var results count_transformpolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("transformpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTransformpolicylabel(key TransformpolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("transformpolicylabel", id, qs)
}

//      RENAME
//      TODO
