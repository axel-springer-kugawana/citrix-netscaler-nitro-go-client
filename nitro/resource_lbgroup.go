package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbgroup struct {
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,string,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	Name                     string `json:"name,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Timeout                  int    `json:"timeout,string,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,string,omitempty"`
}

type LbgroupKey struct {
	Name string
}

func (resource Lbgroup) ToKey() LbgroupKey {
	key := LbgroupKey{
		resource.Name,
	}

	return key
}

func (key LbgroupKey) to_id_args() (string, map[string]string) {
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

type add_lbgroup_payload struct {
	Resource Lbgroup `json:"lbgroup"`
}

func (c *NitroClient) AddLbgroup(resource Lbgroup) error {
	payload := add_lbgroup_payload{
		resource,
	}

	return c.post("lbgroup", "", nil, payload)
}

//      LIST

type list_lbgroup_result struct {
	Results []Lbgroup `json:"lbgroup"`
}

func (c *NitroClient) ListLbgroup() ([]Lbgroup, error) {
	results := list_lbgroup_result{}

	if err := c.get("lbgroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbgroup_result struct {
	Results []Lbgroup `json:"lbgroup"`
}

func (c *NitroClient) GetLbgroup(key LbgroupKey) (*Lbgroup, error) {
	var results get_lbgroup_result

	id, qs := key.to_id_args()

	if err := c.get("lbgroup", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbgroup element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbgroup element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbgroup_result struct {
	Results []Count `json:"lbgroup"`
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

//      EXISTS

func (c *NitroClient) ExistsLbgroup(key LbgroupKey) (bool, error) {
	var results count_lbgroup_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbgroup", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbgroup(key LbgroupKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbgroup", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
