package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmtrafficpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TmtrafficpolicyKey struct {
	Name string
}

func (resource Tmtrafficpolicy) ToKey() TmtrafficpolicyKey {
	key := TmtrafficpolicyKey{
		resource.Name,
	}

	return key
}

func (key TmtrafficpolicyKey) to_id_args() (string, map[string]string) {
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

type add_tmtrafficpolicy_payload struct {
	Resource Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) AddTmtrafficpolicy(resource Tmtrafficpolicy) error {
	payload := add_tmtrafficpolicy_payload{
		resource,
	}

	return c.post("tmtrafficpolicy", "", nil, payload)
}

//      LIST

type list_tmtrafficpolicy_result struct {
	Results []Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) ListTmtrafficpolicy() ([]Tmtrafficpolicy, error) {
	results := list_tmtrafficpolicy_result{}

	if err := c.get("tmtrafficpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_tmtrafficpolicy_result struct {
	Results []Tmtrafficpolicy `json:"tmtrafficpolicy"`
}

func (c *NitroClient) GetTmtrafficpolicy(key TmtrafficpolicyKey) (*Tmtrafficpolicy, error) {
	var results get_tmtrafficpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmtrafficpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmtrafficpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_tmtrafficpolicy_result struct {
	Results []Count `json:"tmtrafficpolicy"`
}

func (c *NitroClient) CountTmtrafficpolicy() (int, error) {
	var results count_tmtrafficpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmtrafficpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTmtrafficpolicy(key TmtrafficpolicyKey) (bool, error) {
	var results count_tmtrafficpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("tmtrafficpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTmtrafficpolicy(key TmtrafficpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("tmtrafficpolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
