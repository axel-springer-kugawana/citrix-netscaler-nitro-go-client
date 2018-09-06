package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Responderpolicylabel struct {
	Comment         string `json:"comment,omitempty"`
	Labelname       string `json:"labelname,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type ResponderpolicylabelKey struct {
	Labelname string
}

func (resource Responderpolicylabel) ToKey() ResponderpolicylabelKey {
	key := ResponderpolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key ResponderpolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ResponderpolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ResponderpolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_responderpolicylabel_payload struct {
	Resource Responderpolicylabel `json:"responderpolicylabel"`
}

func (c *NitroClient) AddResponderpolicylabel(resource Responderpolicylabel) error {
	payload := add_responderpolicylabel_payload{
		resource,
	}

	return c.post("responderpolicylabel", "", nil, payload)
}

//      LIST

type list_responderpolicylabel_result struct {
	Results []Responderpolicylabel `json:"responderpolicylabel"`
}

func (c *NitroClient) ListResponderpolicylabel() ([]Responderpolicylabel, error) {
	results := list_responderpolicylabel_result{}

	if err := c.get("responderpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_responderpolicylabel_result struct {
	Results []Responderpolicylabel `json:"responderpolicylabel"`
}

func (c *NitroClient) GetResponderpolicylabel(key ResponderpolicylabelKey) (*Responderpolicylabel, error) {
	var results get_responderpolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("responderpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one responderpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("responderpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_responderpolicylabel_result struct {
	Results []Count `json:"responderpolicylabel"`
}

func (c *NitroClient) CountResponderpolicylabel() (int, error) {
	var results count_responderpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("responderpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsResponderpolicylabel(key ResponderpolicylabelKey) (bool, error) {
	var results count_responderpolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("responderpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteResponderpolicylabel(key ResponderpolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("responderpolicylabel", id, qs)
}

//      RENAME
//      TODO
