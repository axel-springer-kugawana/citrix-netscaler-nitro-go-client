package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnsaction64 struct {
	Actionname  string `json:"actionname"`
	Excluderule string `json:"excluderule,omitempty"`
	Mappedrule  string `json:"mappedrule,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
}

func dnsaction64_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type Dnsaction64Unset struct {
	Actionname  string `json:"actionname"`
	Prefix      bool   `json:"prefix,string,omitempty"`
	Mappedrule  bool   `json:"mappedrule,string,omitempty"`
	Excluderule bool   `json:"excluderule,string,omitempty"`
}

type update_dnsaction64 struct {
	Actionname  string `json:"actionname"`
	Prefix      string `json:"prefix,omitempty"`
	Mappedrule  string `json:"mappedrule,omitempty"`
	Excluderule string `json:"excluderule,omitempty"`
}

type rename_dnsaction64 struct {
	Name    string `json:"actionname"`
	Newname string `json:"newname"`
}

type add_dnsaction64_payload struct {
	Resource Dnsaction64 `json:"dnsaction64"`
}

type rename_dnsaction64_payload struct {
	Rename rename_dnsaction64 `json:"dnsaction64"`
}

type unset_dnsaction64_payload struct {
	Unset Dnsaction64Unset `json:"dnsaction64"`
}

type update_dnsaction64_payload struct {
	Update update_dnsaction64 `json:"dnsaction64"`
}

type get_dnsaction64_result struct {
	Results []Dnsaction64 `json:"dnsaction64"`
}

type count_dnsaction64_result struct {
	Results []Count `json:"dnsaction64"`
}

func (c *NitroClient) AddDnsaction64(resource Dnsaction64) error {
	payload := add_dnsaction64_payload{
		resource,
	}

	return c.post("dnsaction64", "", nil, payload)
}

func (c *NitroClient) RenameDnsaction64(name string, newName string) error {
	payload := rename_dnsaction64_payload{
		rename_dnsaction64{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("dnsaction64", "", qs, payload)
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

func (c *NitroClient) ExistsDnsaction64(key string) (bool, error) {
	var results count_dnsaction64_result

	id, qs := dnsaction64_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("dnsaction64", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListDnsaction64() ([]Dnsaction64, error) {
	results := get_dnsaction64_result{}

	if err := c.get("dnsaction64", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetDnsaction64(key string) (*Dnsaction64, error) {
	var results get_dnsaction64_result

	id, qs := dnsaction64_key_to_id_args(key)

	if err := c.get("dnsaction64", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dnsaction64 element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("dnsaction64 element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteDnsaction64(key string) error {
	id, qs := dnsaction64_key_to_id_args(key)

	return c.delete("dnsaction64", id, qs)
}

func (c *NitroClient) UnsetDnsaction64(unset Dnsaction64Unset) error {
	payload := unset_dnsaction64_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("dnsaction64", "", qs, payload)
}

func (c *NitroClient) UpdateDnsaction64(resource Dnsaction64) error {
	payload := update_dnsaction64_payload{
		update_dnsaction64{
			resource.Actionname,
			resource.Prefix,
			resource.Mappedrule,
			resource.Excluderule,
		},
	}

	return c.put("dnsaction64", "", nil, payload)
}
