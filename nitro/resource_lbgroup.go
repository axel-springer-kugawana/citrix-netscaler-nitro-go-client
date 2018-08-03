package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbgroup struct {
	Name                     string `json:"name"`
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,string,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Timeout                  int    `json:"timeout,string,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,string,omitempty"`
}

func lbgroup_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type LbgroupUnset struct {
	Name                     string `json:"name"`
	Persistencetype          bool   `json:"persistencetype,string,omitempty"`
	Persistencebackup        bool   `json:"persistencebackup,string,omitempty"`
	Backuppersistencetimeout bool   `json:"backuppersistencetimeout,string,omitempty"`
	Persistmask              bool   `json:"persistmask,string,omitempty"`
	Cookiename               bool   `json:"cookiename,string,omitempty"`
	V6persistmasklen         bool   `json:"v6persistmasklen,string,omitempty"`
	Cookiedomain             bool   `json:"cookiedomain,string,omitempty"`
	Timeout                  bool   `json:"timeout,string,omitempty"`
	Rule                     bool   `json:"rule,string,omitempty"`
	Usevserverpersistency    bool   `json:"usevserverpersistency,string,omitempty"`
}

type update_lbgroup struct {
	Name                     string `json:"name"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,string,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,string,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Timeout                  int    `json:"timeout,string,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
}

type rename_lbgroup struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_lbgroup_payload struct {
	Resource Lbgroup `json:"lbgroup"`
}

type rename_lbgroup_payload struct {
	Rename rename_lbgroup `json:"lbgroup"`
}

type unset_lbgroup_payload struct {
	Unset LbgroupUnset `json:"lbgroup"`
}

type update_lbgroup_payload struct {
	Update update_lbgroup `json:"lbgroup"`
}

type get_lbgroup_result struct {
	Results []Lbgroup `json:"lbgroup"`
}

type count_lbgroup_result struct {
	Results []Count `json:"lbgroup"`
}

func (c *NitroClient) AddLbgroup(resource Lbgroup) error {
	payload := add_lbgroup_payload{
		resource,
	}

	return c.post("lbgroup", "", nil, payload)
}

func (c *NitroClient) RenameLbgroup(name string, newName string) error {
	payload := rename_lbgroup_payload{
		rename_lbgroup{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbgroup", "", qs, payload)
}

func (c *NitroClient) CountLbgroup() (int, error) {
	var results count_lbgroup_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbgroup", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbgroup(key string) (bool, error) {
	var results count_lbgroup_result

	id, qs := lbgroup_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbgroup", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbgroup() ([]Lbgroup, error) {
	results := get_lbgroup_result{}

	if err := c.get("lbgroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbgroup(key string) (*Lbgroup, error) {
	var results get_lbgroup_result

	id, qs := lbgroup_key_to_id_args(key)

	if err := c.get("lbgroup", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbgroup element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbgroup element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbgroup(key string) error {
	id, qs := lbgroup_key_to_id_args(key)

	return c.delete("lbgroup", id, qs)
}

func (c *NitroClient) UnsetLbgroup(unset LbgroupUnset) error {
	payload := unset_lbgroup_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("lbgroup", "", qs, payload)
}

func (c *NitroClient) UpdateLbgroup(resource Lbgroup) error {
	payload := update_lbgroup_payload{
		update_lbgroup{
			resource.Name,
			resource.Persistencetype,
			resource.Persistencebackup,
			resource.Backuppersistencetimeout,
			resource.Persistmask,
			resource.Cookiename,
			resource.V6persistmasklen,
			resource.Cookiedomain,
			resource.Timeout,
			resource.Rule,
			resource.Usevserverpersistency,
		},
	}

	return c.put("lbgroup", "", nil, payload)
}
