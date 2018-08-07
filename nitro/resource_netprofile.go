package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Netprofile struct {
	Name             string `json:"name,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,string,omitempty"`
}

type NetprofileKey struct {
	Name string
}

func (resource Netprofile) ToKey() NetprofileKey {
	key := NetprofileKey{
		resource.Name,
	}

	return key
}

func (key NetprofileKey) to_id_args() (string, map[string]string) {
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

type add_netprofile_payload struct {
	Resource Netprofile `json:"netprofile"`
}

func (c *NitroClient) AddNetprofile(resource Netprofile) error {
	payload := add_netprofile_payload{
		resource,
	}

	return c.post("netprofile", "", nil, payload)
}

//      LIST

type list_netprofile_result struct {
	Results []Netprofile `json:"netprofile"`
}

func (c *NitroClient) ListNetprofile() ([]Netprofile, error) {
	results := list_netprofile_result{}

	if err := c.get("netprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_netprofile_result struct {
	Results []Netprofile `json:"netprofile"`
}

func (c *NitroClient) GetNetprofile(key NetprofileKey) (*Netprofile, error) {
	var results get_netprofile_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_netprofile_result struct {
	Results []Count `json:"netprofile"`
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

//      EXISTS

func (c *NitroClient) ExistsNetprofile(key NetprofileKey) (bool, error) {
	var results count_netprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("netprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteNetprofile(key NetprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("netprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
