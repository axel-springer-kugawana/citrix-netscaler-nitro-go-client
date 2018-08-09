package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appqoepolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AppqoepolicyKey struct {
	Name string
}

func (resource Appqoepolicy) ToKey() AppqoepolicyKey {
	key := AppqoepolicyKey{
		resource.Name,
	}

	return key
}

func (key AppqoepolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AppqoepolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppqoepolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appqoepolicy_payload struct {
	Resource Appqoepolicy `json:"appqoepolicy"`
}

func (c *NitroClient) AddAppqoepolicy(resource Appqoepolicy) error {
	payload := add_appqoepolicy_payload{
		resource,
	}

	return c.post("appqoepolicy", "", nil, payload)
}

//      LIST

type list_appqoepolicy_result struct {
	Results []Appqoepolicy `json:"appqoepolicy"`
}

func (c *NitroClient) ListAppqoepolicy() ([]Appqoepolicy, error) {
	results := list_appqoepolicy_result{}

	if err := c.get("appqoepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appqoepolicy_result struct {
	Results []Appqoepolicy `json:"appqoepolicy"`
}

func (c *NitroClient) GetAppqoepolicy(key AppqoepolicyKey) (*Appqoepolicy, error) {
	var results get_appqoepolicy_result

	id, qs := key.to_id_args()

	if err := c.get("appqoepolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appqoepolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appqoepolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appqoepolicy_result struct {
	Results []Count `json:"appqoepolicy"`
}

func (c *NitroClient) CountAppqoepolicy() (int, error) {
	var results count_appqoepolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appqoepolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppqoepolicy(key AppqoepolicyKey) (bool, error) {
	var results count_appqoepolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appqoepolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppqoepolicy(key AppqoepolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("appqoepolicy", id, qs)
}

//      UPDATE

type AppqoepolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Appqoepolicy) ToUpdate() AppqoepolicyUpdate {
	update := AppqoepolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_appqoepolicy_payload struct {
	Update AppqoepolicyUpdate `json:"appqoepolicy"`
}

func (c *NitroClient) UpdateAppqoepolicy(update AppqoepolicyUpdate) error {
	payload := update_appqoepolicy_payload{
		update,
	}

	return c.put("appqoepolicy", "", nil, payload)
}

//      UNSET

type AppqoepolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Appqoepolicy) ToUnset() AppqoepolicyUnset {
	unset := AppqoepolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_appqoepolicy_payload struct {
	Unset AppqoepolicyUnset `json:"appqoepolicy"`
}

func (c *NitroClient) UnsetAppqoepolicy(unset AppqoepolicyUnset) error {
	payload := unset_appqoepolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appqoepolicy", "", qs, payload)
}

//      RENAME
//      TODO
