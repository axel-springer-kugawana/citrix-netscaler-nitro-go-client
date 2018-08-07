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

func (key Dnspolicy64Key) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
