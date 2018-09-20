package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Cspolicylabel struct {
	Cspolicylabeltype string `json:"cspolicylabeltype,omitempty"`
	Labelname         string `json:"labelname,omitempty"`
}

type CspolicylabelKey struct {
	Labelname string
}

func (resource Cspolicylabel) ToKey() CspolicylabelKey {
	key := CspolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key CspolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Labelname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CspolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CspolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cspolicylabel_payload struct {
	Resource Cspolicylabel `json:"cspolicylabel"`
}

func (c *NitroClient) AddCspolicylabel(resource Cspolicylabel) error {
	payload := add_cspolicylabel_payload{
		resource,
	}

	return c.post("cspolicylabel", "", nil, payload)
}

//      LIST

type list_cspolicylabel_result struct {
	Results []Cspolicylabel `json:"cspolicylabel"`
}

func (c *NitroClient) ListCspolicylabel() ([]Cspolicylabel, error) {
	results := list_cspolicylabel_result{}

	if err := c.get("cspolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cspolicylabel_result struct {
	Results []Cspolicylabel `json:"cspolicylabel"`
}

func (c *NitroClient) GetCspolicylabel(key CspolicylabelKey) (*Cspolicylabel, error) {
	var results get_cspolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("cspolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cspolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cspolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cspolicylabel_result struct {
	Results []Count `json:"cspolicylabel"`
}

func (c *NitroClient) CountCspolicylabel() (int, error) {
	var results count_cspolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cspolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCspolicylabel(key CspolicylabelKey) (bool, error) {
	var results count_cspolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cspolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCspolicylabel(key CspolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("cspolicylabel", id, qs)
}

//      RENAME
//      TODO
