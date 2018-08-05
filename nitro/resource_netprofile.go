package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Netprofile struct {
	Name             string `json:"name"`
	Overridelsn      string `json:"overridelsn,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,string,omitempty"`
}

func netprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type NetprofileUnset struct {
	Name             string `json:"name"`
	Srcip            bool   `json:"srcip,omitempty"`
	Srcippersistency bool   `json:"srcippersistency,omitempty"`
	Overridelsn      bool   `json:"overridelsn,omitempty"`
}

type update_netprofile struct {
	Name             string `json:"name"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
}

type rename_netprofile struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_netprofile_payload struct {
	Resource Netprofile `json:"netprofile"`
}

type rename_netprofile_payload struct {
	Rename rename_netprofile `json:"netprofile"`
}

type unset_netprofile_payload struct {
	Unset NetprofileUnset `json:"netprofile"`
}

type update_netprofile_payload struct {
	Update update_netprofile `json:"netprofile"`
}

type get_netprofile_result struct {
	Results []Netprofile `json:"netprofile"`
}

type count_netprofile_result struct {
	Results []Count `json:"netprofile"`
}

func (c *NitroClient) AddNetprofile(resource Netprofile) error {
	payload := add_netprofile_payload{
		resource,
	}

	return c.post("netprofile", "", nil, payload)
}

func (c *NitroClient) RenameNetprofile(name string, newName string) error {
	payload := rename_netprofile_payload{
		rename_netprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("netprofile", "", qs, payload)
}

func (c *NitroClient) CountNetprofile() (int, error) {
	var results count_netprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("netprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsNetprofile(key string) (bool, error) {
	var results count_netprofile_result

	id, qs := netprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("netprofile", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListNetprofile() ([]Netprofile, error) {
	results := get_netprofile_result{}

	if err := c.get("netprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetNetprofile(key string) (*Netprofile, error) {
	var results get_netprofile_result

	id, qs := netprofile_key_to_id_args(key)

	if err := c.get("netprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one netprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("netprofile element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteNetprofile(key string) error {
	id, qs := netprofile_key_to_id_args(key)

	return c.delete("netprofile", id, qs)
}

func (c *NitroClient) UnsetNetprofile(unset NetprofileUnset) error {
	payload := unset_netprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("netprofile", "", qs, payload)
}

func (c *NitroClient) UpdateNetprofile(resource Netprofile) error {
	payload := update_netprofile_payload{
		update_netprofile{
			resource.Name,
			resource.Srcip,
			resource.Srcippersistency,
			resource.Overridelsn,
		},
	}

	return c.put("netprofile", "", nil, payload)
}
