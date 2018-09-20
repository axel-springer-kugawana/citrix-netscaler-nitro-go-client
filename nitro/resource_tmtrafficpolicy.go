package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Tmtrafficpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TmtrafficpolicyKey struct {
	Name string
}

func (resource Tmtrafficpolicy) ToKey() TmtrafficpolicyKey {
	key := TmtrafficpolicyKey{
		resource.Name,
	}

	return key
}

func (key TmtrafficpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TmtrafficpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TmtrafficpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_tmtrafficpolicy_payload struct {
	Resource Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) AddTmtrafficpolicy(resource Tmtrafficpolicy) error {
	payload := add_tmtrafficpolicy_payload{
		resource,
	}

	return c.post("tmtrafficpolicy", "", nil, payload)
}

//      LIST

type list_tmtrafficpolicy_result struct {
	Results []Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) ListTmtrafficpolicy() ([]Tmtrafficpolicy, error) {
	results := list_tmtrafficpolicy_result{}

	if err := c.get("tmtrafficpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_tmtrafficpolicy_result struct {
	Results []Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) GetTmtrafficpolicy(key TmtrafficpolicyKey) (*Tmtrafficpolicy, error) {
	var results get_tmtrafficpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmtrafficpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmtrafficpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_tmtrafficpolicy_result struct {
	Results []Count `json:"tmtrafficpolicy"`
}

func (c *NitroClient) CountTmtrafficpolicy() (int, error) {
	var results count_tmtrafficpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmtrafficpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTmtrafficpolicy(key TmtrafficpolicyKey) (bool, error) {
	var results count_tmtrafficpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTmtrafficpolicy(key TmtrafficpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("tmtrafficpolicy", id, qs)
}

//      UPDATE

type TmtrafficpolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Tmtrafficpolicy) ToUpdate() TmtrafficpolicyUpdate {
	update := TmtrafficpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_tmtrafficpolicy_payload struct {
	Update TmtrafficpolicyUpdate `json:"tmtrafficpolicy"`
}

func (c *NitroClient) UpdateTmtrafficpolicy(update TmtrafficpolicyUpdate) error {
	payload := update_tmtrafficpolicy_payload{
		update,
	}

	return c.put("tmtrafficpolicy", "", nil, payload)
}

//      UNSET

type TmtrafficpolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Tmtrafficpolicy) ToUnset() TmtrafficpolicyUnset {
	unset := TmtrafficpolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_tmtrafficpolicy_payload struct {
	Unset TmtrafficpolicyUnset `json:"tmtrafficpolicy"`
}

func (c *NitroClient) UnsetTmtrafficpolicy(unset TmtrafficpolicyUnset) error {
	payload := unset_tmtrafficpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("tmtrafficpolicy", "", qs, payload)
}

//      RENAME
//      TODO
