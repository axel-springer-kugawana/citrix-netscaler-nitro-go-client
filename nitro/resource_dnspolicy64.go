package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnspolicy64 struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type Dnspolicy64Key struct {
	Name string
}

func (resource Dnspolicy64) ToKey() Dnspolicy64Key {
	key := Dnspolicy64Key{
		resource.Name,
	}

	return key
}

func (key Dnspolicy64Key) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key Dnspolicy64Key) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key Dnspolicy64Key) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_dnspolicy64_payload struct {
	Resource Dnspolicy64 `json:"dnspolicy64"`
}

func (c *NitroClient) AddDnspolicy64(resource Dnspolicy64) error {
	payload := add_dnspolicy64_payload{
		resource,
	}

	return c.post("dnspolicy64", "", nil, payload)
}

//      LIST

type list_dnspolicy64_result struct {
	Results []Dnspolicy64 `json:"dnspolicy64"`
}

func (c *NitroClient) ListDnspolicy64() ([]Dnspolicy64, error) {
	results := list_dnspolicy64_result{}

	if err := c.get("dnspolicy64", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_dnspolicy64_result struct {
	Results []Dnspolicy64 `json:"dnspolicy64"`
}

func (c *NitroClient) GetDnspolicy64(key Dnspolicy64Key) (*Dnspolicy64, error) {
	var results get_dnspolicy64_result

	id, qs := key.to_id_args()

	if err := c.get("dnspolicy64", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dnspolicy64 element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dnspolicy64 element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_dnspolicy64_result struct {
	Results []Count `json:"dnspolicy64"`
}

func (c *NitroClient) CountDnspolicy64() (int, error) {
	var results count_dnspolicy64_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dnspolicy64", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsDnspolicy64(key Dnspolicy64Key) (bool, error) {
	var results count_dnspolicy64_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("dnspolicy64", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteDnspolicy64(key Dnspolicy64Key) error {
	id, qs := key.to_id_args()

	return c.delete("dnspolicy64", id, qs)
}

//      UPDATE

type Dnspolicy64Update struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Dnspolicy64) ToUpdate() Dnspolicy64Update {
	update := Dnspolicy64Update{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_dnspolicy64_payload struct {
	Update Dnspolicy64Update `json:"dnspolicy64"`
}

func (c *NitroClient) UpdateDnspolicy64(update Dnspolicy64Update) error {
	payload := update_dnspolicy64_payload{
		update,
	}

	return c.put("dnspolicy64", "", nil, payload)
}

//      UNSET

type Dnspolicy64Unset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Dnspolicy64) ToUnset() Dnspolicy64Unset {
	unset := Dnspolicy64Unset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_dnspolicy64_payload struct {
	Unset Dnspolicy64Unset `json:"dnspolicy64"`
}

func (c *NitroClient) UnsetDnspolicy64(unset Dnspolicy64Unset) error {
	payload := unset_dnspolicy64_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("dnspolicy64", "", qs, payload)
}

//      RENAME
//      TODO
