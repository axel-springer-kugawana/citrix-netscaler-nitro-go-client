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

func (key AppqoepolicyKey) to_id_args() (string, map[string]string) {
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
