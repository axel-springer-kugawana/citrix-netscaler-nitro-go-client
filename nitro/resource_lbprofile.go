package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbprofile struct {
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
}

type LbprofileKey struct {
	Lbprofilename string
}

func (resource Lbprofile) ToKey() LbprofileKey {
	key := LbprofileKey{
		resource.Lbprofilename,
	}

	return key
}

func (key LbprofileKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Lbprofilename

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_lbprofile_payload struct {
	Resource Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) AddLbprofile(resource Lbprofile) error {
	payload := add_lbprofile_payload{
		resource,
	}

	return c.post("lbprofile", "", nil, payload)
}

//      LIST

type list_lbprofile_result struct {
	Results []Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) ListLbprofile() ([]Lbprofile, error) {
	results := list_lbprofile_result{}

	if err := c.get("lbprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbprofile_result struct {
	Results []Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) GetLbprofile(key LbprofileKey) (*Lbprofile, error) {
	var results get_lbprofile_result

	id, qs := key.to_id_args()

	if err := c.get("lbprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbprofile element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbprofile_result struct {
	Results []Count `json:"lbprofile"`
}

func (c *NitroClient) CountLbprofile() (int, error) {
	var results count_lbprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbprofile(key LbprofileKey) (bool, error) {
	var results count_lbprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbprofile(key LbprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
