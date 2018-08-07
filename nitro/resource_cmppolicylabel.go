package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmppolicylabel struct {
	Labelname string `json:"labelname,omitempty"`
	Type      string `json:"type,omitempty"`
}

type CmppolicylabelKey struct {
	Labelname string
}

func (resource Cmppolicylabel) ToKey() CmppolicylabelKey {
	key := CmppolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key CmppolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CmppolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CmppolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cmppolicylabel_payload struct {
	Resource Cmppolicylabel `json:"cmppolicylabel"`
}

func (c *NitroClient) AddCmppolicylabel(resource Cmppolicylabel) error {
	payload := add_cmppolicylabel_payload{
		resource,
	}

	return c.post("cmppolicylabel", "", nil, payload)
}

//      LIST

type list_cmppolicylabel_result struct {
	Results []Cmppolicylabel `json:"cmppolicylabel"`
}

func (c *NitroClient) ListCmppolicylabel() ([]Cmppolicylabel, error) {
	results := list_cmppolicylabel_result{}

	if err := c.get("cmppolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cmppolicylabel_result struct {
	Results []Cmppolicylabel `json:"cmppolicylabel"`
}

func (c *NitroClient) GetCmppolicylabel(key CmppolicylabelKey) (*Cmppolicylabel, error) {
	var results get_cmppolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("cmppolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmppolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cmppolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cmppolicylabel_result struct {
	Results []Count `json:"cmppolicylabel"`
}

func (c *NitroClient) CountCmppolicylabel() (int, error) {
	var results count_cmppolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cmppolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCmppolicylabel(key CmppolicylabelKey) (bool, error) {
	var results count_cmppolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cmppolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCmppolicylabel(key CmppolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("cmppolicylabel", id, qs)
}

//      RENAME
//      TODO
