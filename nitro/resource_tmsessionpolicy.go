package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Tmsessionpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TmsessionpolicyKey struct {
	Name string
}

func (resource Tmsessionpolicy) ToKey() TmsessionpolicyKey {
	key := TmsessionpolicyKey{
		resource.Name,
	}

	return key
}

func (key TmsessionpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TmsessionpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TmsessionpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_tmsessionpolicy_payload struct {
	Resource Tmsessionpolicy `json:"tmsessionpolicy"`
}

func (c *NitroClient) AddTmsessionpolicy(resource Tmsessionpolicy) error {
	payload := add_tmsessionpolicy_payload{
		resource,
	}

	return c.post("tmsessionpolicy", "", nil, payload)
}

//      LIST

type list_tmsessionpolicy_result struct {
	Results []Tmsessionpolicy `json:"tmsessionpolicy"`
}

func (c *NitroClient) ListTmsessionpolicy() ([]Tmsessionpolicy, error) {
	results := list_tmsessionpolicy_result{}

	if err := c.get("tmsessionpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_tmsessionpolicy_result struct {
	Results []Tmsessionpolicy `json:"tmsessionpolicy"`
}

func (c *NitroClient) GetTmsessionpolicy(key TmsessionpolicyKey) (*Tmsessionpolicy, error) {
	var results get_tmsessionpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("tmsessionpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmsessionpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmsessionpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_tmsessionpolicy_result struct {
	Results []Count `json:"tmsessionpolicy"`
}

func (c *NitroClient) CountTmsessionpolicy() (int, error) {
	var results count_tmsessionpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmsessionpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTmsessionpolicy(key TmsessionpolicyKey) (bool, error) {
	var results count_tmsessionpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("tmsessionpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTmsessionpolicy(key TmsessionpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("tmsessionpolicy", id, qs)
}

//      UPDATE

type TmsessionpolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Tmsessionpolicy) ToUpdate() TmsessionpolicyUpdate {
	update := TmsessionpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_tmsessionpolicy_payload struct {
	Update TmsessionpolicyUpdate `json:"tmsessionpolicy"`
}

func (c *NitroClient) UpdateTmsessionpolicy(update TmsessionpolicyUpdate) error {
	payload := update_tmsessionpolicy_payload{
		update,
	}

	return c.put("tmsessionpolicy", "", nil, payload)
}

//      UNSET

type TmsessionpolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Tmsessionpolicy) ToUnset() TmsessionpolicyUnset {
	unset := TmsessionpolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_tmsessionpolicy_payload struct {
	Unset TmsessionpolicyUnset `json:"tmsessionpolicy"`
}

func (c *NitroClient) UnsetTmsessionpolicy(unset TmsessionpolicyUnset) error {
	payload := unset_tmsessionpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("tmsessionpolicy", "", qs, payload)
}

//      RENAME
//      TODO
