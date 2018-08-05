package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbprofile struct {
	Lbprofilename                 string `json:"lbprofilename"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
}

func lbprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type LbprofileUnset struct {
	Lbprofilename                 string `json:"lbprofilename"`
	Cookiepassphrase              bool   `json:"cookiepassphrase,string,omitempty"`
	Dbslb                         bool   `json:"dbslb,string,omitempty"`
	Processlocal                  bool   `json:"processlocal,string,omitempty"`
	Httponlycookieflag            bool   `json:"httponlycookieflag,string,omitempty"`
	Usesecuredpersistencecookie   bool   `json:"usesecuredpersistencecookie,string,omitempty"`
	Useencryptedpersistencecookie bool   `json:"useencryptedpersistencecookie,string,omitempty"`
}

type update_lbprofile struct {
	Lbprofilename                 string `json:"lbprofilename"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
}

type rename_lbprofile struct {
	Name    string `json:"lbprofilename"`
	Newname string `json:"newname"`
}

type add_lbprofile_payload struct {
	Resource Lbprofile `json:"lbprofile"`
}

type rename_lbprofile_payload struct {
	Rename rename_lbprofile `json:"lbprofile"`
}

type unset_lbprofile_payload struct {
	Unset LbprofileUnset `json:"lbprofile"`
}

type update_lbprofile_payload struct {
	Update update_lbprofile `json:"lbprofile"`
}

type get_lbprofile_result struct {
	Results []Lbprofile `json:"lbprofile"`
}

type count_lbprofile_result struct {
	Results []Count `json:"lbprofile"`
}

func (c *NitroClient) AddLbprofile(resource Lbprofile) error {
	payload := add_lbprofile_payload{
		resource,
	}

	return c.post("lbprofile", "", nil, payload)
}

func (c *NitroClient) RenameLbprofile(name string, newName string) error {
	payload := rename_lbprofile_payload{
		rename_lbprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbprofile", "", qs, payload)
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

func (c *NitroClient) ExistsLbprofile(key string) (bool, error) {
	var results count_lbprofile_result

	id, qs := lbprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbprofile", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbprofile() ([]Lbprofile, error) {
	results := get_lbprofile_result{}

	if err := c.get("lbprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbprofile(key string) (*Lbprofile, error) {
	var results get_lbprofile_result

	id, qs := lbprofile_key_to_id_args(key)

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

func (c *NitroClient) DeleteLbprofile(key string) error {
	id, qs := lbprofile_key_to_id_args(key)

	return c.delete("lbprofile", id, qs)
}

func (c *NitroClient) UnsetLbprofile(unset LbprofileUnset) error {
	payload := unset_lbprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("lbprofile", "", qs, payload)
}

func (c *NitroClient) UpdateLbprofile(resource Lbprofile) error {
	payload := update_lbprofile_payload{
		update_lbprofile{
			resource.Lbprofilename,
			resource.Cookiepassphrase,
			resource.Dbslb,
			resource.Processlocal,
			resource.Httponlycookieflag,
			resource.Usesecuredpersistencecookie,
			resource.Useencryptedpersistencecookie,
		},
	}

	return c.put("lbprofile", "", nil, payload)
}
