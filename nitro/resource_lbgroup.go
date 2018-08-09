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

func (key LbgroupKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbgroupKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbgroupKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type LbgroupUpdate struct {
	Name                     string `json:"name,omitempty"`
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

func (resource Lbgroup) ToUpdate() LbgroupUpdate {
	update := LbgroupUpdate{
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
	}

	return update
}

type update_lbgroup_payload struct {
	Update LbgroupUpdate `json:"lbgroup"`
}

func (c *NitroClient) UpdateLbgroup(update LbgroupUpdate) error {
	payload := update_lbgroup_payload{
		update,
	}

	return c.put("lbgroup", "", nil, payload)
}

//      UNSET

type LbgroupUnset struct {
	Name                     string `json:"name,omitempty"`
	Persistencetype          bool   `json:"persistencetype,omitempty"`
	Persistencebackup        bool   `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout bool   `json:"backuppersistencetimeout,omitempty"`
	Persistmask              bool   `json:"persistmask,omitempty"`
	Cookiename               bool   `json:"cookiename,omitempty"`
	V6persistmasklen         bool   `json:"v6persistmasklen,omitempty"`
	Cookiedomain             bool   `json:"cookiedomain,omitempty"`
	Timeout                  bool   `json:"timeout,omitempty"`
	Rule                     bool   `json:"rule,omitempty"`
	Usevserverpersistency    bool   `json:"usevserverpersistency,omitempty"`
}

func (resource Lbgroup) ToUnset() LbgroupUnset {
	unset := LbgroupUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_lbgroup_payload struct {
	Unset LbgroupUnset `json:"lbgroup"`
}

func (c *NitroClient) UnsetLbgroup(unset LbgroupUnset) error {
	payload := unset_lbgroup_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("lbgroup", "", qs, payload)
}

//      RENAME
//      TODO
