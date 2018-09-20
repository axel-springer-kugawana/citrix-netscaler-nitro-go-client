package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Policystringmap struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
}

type PolicystringmapKey struct {
	Name string
}

func (resource Policystringmap) ToKey() PolicystringmapKey {
	key := PolicystringmapKey{
		resource.Name,
	}

	return key
}

func (key PolicystringmapKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key PolicystringmapKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PolicystringmapKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_policystringmap_payload struct {
	Resource Policystringmap `json:"policystringmap"`
}

func (c *NitroClient) AddPolicystringmap(resource Policystringmap) error {
	payload := add_policystringmap_payload{
		resource,
	}

	return c.post("policystringmap", "", nil, payload)
}

//      LIST

type list_policystringmap_result struct {
	Results []Policystringmap `json:"policystringmap"`
}

func (c *NitroClient) ListPolicystringmap() ([]Policystringmap, error) {
	results := list_policystringmap_result{}

	if err := c.get("policystringmap", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_policystringmap_result struct {
	Results []Policystringmap `json:"policystringmap"`
}

func (c *NitroClient) GetPolicystringmap(key PolicystringmapKey) (*Policystringmap, error) {
	var results get_policystringmap_result

	id, qs := key.to_id_args()

	if err := c.get("policystringmap", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policystringmap element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policystringmap element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_policystringmap_result struct {
	Results []Count `json:"policystringmap"`
}

func (c *NitroClient) CountPolicystringmap() (int, error) {
	var results count_policystringmap_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policystringmap", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicystringmap(key PolicystringmapKey) (bool, error) {
	var results count_policystringmap_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("policystringmap", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicystringmap(key PolicystringmapKey) error {
	id, qs := key.to_id_args()

	return c.delete("policystringmap", id, qs)
}

//      UPDATE

type PolicystringmapUpdate struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func (resource Policystringmap) ToUpdate() PolicystringmapUpdate {
	update := PolicystringmapUpdate{
		resource.Name,
		resource.Comment,
	}

	return update
}

type update_policystringmap_payload struct {
	Update PolicystringmapUpdate `json:"policystringmap"`
}

func (c *NitroClient) UpdatePolicystringmap(update PolicystringmapUpdate) error {
	payload := update_policystringmap_payload{
		update,
	}

	return c.put("policystringmap", "", nil, payload)
}

//      UNSET

type PolicystringmapUnset struct {
	Name    string `json:"name,omitempty"`
	Comment bool   `json:"comment,omitempty"`
}

func (resource Policystringmap) ToUnset() PolicystringmapUnset {
	unset := PolicystringmapUnset{
		resource.Name,
		false,
	}

	return unset
}

type unset_policystringmap_payload struct {
	Unset PolicystringmapUnset `json:"policystringmap"`
}

func (c *NitroClient) UnsetPolicystringmap(unset PolicystringmapUnset) error {
	payload := unset_policystringmap_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("policystringmap", "", qs, payload)
}

//      RENAME
//      TODO
