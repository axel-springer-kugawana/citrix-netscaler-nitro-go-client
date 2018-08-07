package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dbdbprofile struct {
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Name                   string `json:"name,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
}

type DbdbprofileKey struct {
	Name string
}

func (resource Dbdbprofile) ToKey() DbdbprofileKey {
	key := DbdbprofileKey{
		resource.Name,
	}

	return key
}

func (key DbdbprofileKey) to_id_args() (string, map[string]string) {
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

type add_dbdbprofile_payload struct {
	Resource Dbdbprofile `json:"dbdbprofile"`
}

func (c *NitroClient) AddDbdbprofile(resource Dbdbprofile) error {
	payload := add_dbdbprofile_payload{
		resource,
	}

	return c.post("dbdbprofile", "", nil, payload)
}

//      LIST

type list_dbdbprofile_result struct {
	Results []Dbdbprofile `json:"dbdbprofile"`
}

func (c *NitroClient) ListDbdbprofile() ([]Dbdbprofile, error) {
	results := list_dbdbprofile_result{}

	if err := c.get("dbdbprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_dbdbprofile_result struct {
	Results []Dbdbprofile `json:"dbdbprofile"`
}

func (c *NitroClient) GetDbdbprofile(key DbdbprofileKey) (*Dbdbprofile, error) {
	var results get_dbdbprofile_result

	id, qs := key.to_id_args()

	if err := c.get("dbdbprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dbdbprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dbdbprofile element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_dbdbprofile_result struct {
	Results []Count `json:"dbdbprofile"`
}

func (c *NitroClient) CountDbdbprofile() (int, error) {
	var results count_dbdbprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dbdbprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsDbdbprofile(key DbdbprofileKey) (bool, error) {
	var results count_dbdbprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("dbdbprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteDbdbprofile(key DbdbprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("dbdbprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
