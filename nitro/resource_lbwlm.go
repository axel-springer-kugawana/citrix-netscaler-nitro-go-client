package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbwlm struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Katimeout int    `json:"katimeout,string,omitempty"`
	Lbuid     string `json:"lbuid,omitempty"`
	Port      int    `json:"port,omitempty"`
	Wlmname   string `json:"wlmname,omitempty"`
}

type LbwlmKey struct {
	Wlmname string
}

func (resource Lbwlm) ToKey() LbwlmKey {
	key := LbwlmKey{
		resource.Wlmname,
	}

	return key
}

func (key LbwlmKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Wlmname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbwlmKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbwlmKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbwlm_payload struct {
	Resource Lbwlm `json:"lbwlm"`
}

func (c *NitroClient) AddLbwlm(resource Lbwlm) error {
	payload := add_lbwlm_payload{
		resource,
	}

	return c.post("lbwlm", "", nil, payload)
}

//      LIST

type list_lbwlm_result struct {
	Results []Lbwlm `json:"lbwlm"`
}

func (c *NitroClient) ListLbwlm() ([]Lbwlm, error) {
	results := list_lbwlm_result{}

	if err := c.get("lbwlm", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbwlm_result struct {
	Results []Lbwlm `json:"lbwlm"`
}

func (c *NitroClient) GetLbwlm(key LbwlmKey) (*Lbwlm, error) {
	var results get_lbwlm_result

	id, qs := key.to_id_args()

	if err := c.get("lbwlm", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbwlm element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbwlm element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbwlm_result struct {
	Results []Count `json:"lbwlm"`
}

func (c *NitroClient) CountLbwlm() (int, error) {
	var results count_lbwlm_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbwlm", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbwlm(key LbwlmKey) (bool, error) {
	var results count_lbwlm_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbwlm", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbwlm(key LbwlmKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbwlm", id, qs)
}

//      UPDATE

type LbwlmUpdate struct {
	Wlmname   string `json:"wlmname,omitempty"`
	Katimeout int    `json:"katimeout,string,omitempty"`
}

func (resource Lbwlm) ToUpdate() LbwlmUpdate {
	update := LbwlmUpdate{
		resource.Wlmname,
		resource.Katimeout,
	}

	return update
}

type update_lbwlm_payload struct {
	Update LbwlmUpdate `json:"lbwlm"`
}

func (c *NitroClient) UpdateLbwlm(update LbwlmUpdate) error {
	payload := update_lbwlm_payload{
		update,
	}

	return c.put("lbwlm", "", nil, payload)
}

//      UNSET

type LbwlmUnset struct {
	Wlmname   string `json:"wlmname,omitempty"`
	Katimeout bool   `json:"katimeout,omitempty"`
}

func (resource Lbwlm) ToUnset() LbwlmUnset {
	unset := LbwlmUnset{
		resource.Wlmname,
		false,
	}

	return unset
}

type unset_lbwlm_payload struct {
	Unset LbwlmUnset `json:"lbwlm"`
}

func (c *NitroClient) UnsetLbwlm(unset LbwlmUnset) error {
	payload := unset_lbwlm_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("lbwlm", "", qs, payload)
}

//      RENAME
//      TODO
