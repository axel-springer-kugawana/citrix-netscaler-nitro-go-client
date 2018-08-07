package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnsaction64 struct {
	Actionname  string `json:"actionname,omitempty"`
	Excluderule string `json:"excluderule,omitempty"`
	Mappedrule  string `json:"mappedrule,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
}

type Dnsaction64Key struct {
	Actionname string
}

func (resource Dnsaction64) ToKey() Dnsaction64Key {
	key := Dnsaction64Key{
		resource.Actionname,
	}

	return key
}

func (key Dnsaction64Key) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Actionname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key Dnsaction64Key) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key Dnsaction64Key) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_dnsaction64_payload struct {
	Resource Dnsaction64 `json:"dnsaction64"`
}

func (c *NitroClient) AddDnsaction64(resource Dnsaction64) error {
	payload := add_dnsaction64_payload{
		resource,
	}

	return c.post("dnsaction64", "", nil, payload)
}

//      LIST

type list_dnsaction64_result struct {
	Results []Dnsaction64 `json:"dnsaction64"`
}

func (c *NitroClient) ListDnsaction64() ([]Dnsaction64, error) {
	results := list_dnsaction64_result{}

	if err := c.get("dnsaction64", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_dnsaction64_result struct {
	Results []Dnsaction64 `json:"dnsaction64"`
}

func (c *NitroClient) GetDnsaction64(key Dnsaction64Key) (*Dnsaction64, error) {
	var results get_dnsaction64_result

	id, qs := key.to_id_args()

	if err := c.get("dnsaction64", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dnsaction64 element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dnsaction64 element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_dnsaction64_result struct {
	Results []Count `json:"dnsaction64"`
}

func (c *NitroClient) CountDnsaction64() (int, error) {
	var results count_dnsaction64_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dnsaction64", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsDnsaction64(key Dnsaction64Key) (bool, error) {
	var results count_dnsaction64_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("dnsaction64", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteDnsaction64(key Dnsaction64Key) error {
	id, qs := key.to_id_args()

	return c.delete("dnsaction64", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
