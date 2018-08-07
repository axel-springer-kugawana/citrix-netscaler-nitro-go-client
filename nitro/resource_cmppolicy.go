package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmppolicy struct {
	Name      string `json:"name,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

type CmppolicyKey struct {
	Name string
}

func (resource Cmppolicy) ToKey() CmppolicyKey {
	key := CmppolicyKey{
		resource.Name,
	}

	return key
}

func (key CmppolicyKey) to_id_args() (string, map[string]string) {
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

type add_cmppolicy_payload struct {
	Resource Cmppolicy `json:"cmppolicy"`
}

func (c *NitroClient) AddCmppolicy(resource Cmppolicy) error {
	payload := add_cmppolicy_payload{
		resource,
	}

	return c.post("cmppolicy", "", nil, payload)
}

//      LIST

type list_cmppolicy_result struct {
	Results []Cmppolicy `json:"cmppolicy"`
}

func (c *NitroClient) ListCmppolicy() ([]Cmppolicy, error) {
	results := list_cmppolicy_result{}

	if err := c.get("cmppolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cmppolicy_result struct {
	Results []Cmppolicy `json:"cmppolicy"`
}

func (c *NitroClient) GetCmppolicy(key CmppolicyKey) (*Cmppolicy, error) {
	var results get_cmppolicy_result

	id, qs := key.to_id_args()

	if err := c.get("cmppolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmppolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cmppolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cmppolicy_result struct {
	Results []Count `json:"cmppolicy"`
}

func (c *NitroClient) CountCmppolicy() (int, error) {
	var results count_cmppolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cmppolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCmppolicy(key CmppolicyKey) (bool, error) {
	var results count_cmppolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cmppolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCmppolicy(key CmppolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("cmppolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
