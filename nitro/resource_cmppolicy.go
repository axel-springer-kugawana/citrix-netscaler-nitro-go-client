package nitro

import (
	"fmt"
	"net/url"
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

func (key CmppolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CmppolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CmppolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type CmppolicyUpdate struct {
	Name      string `json:"name,omitempty"`
	Rule      string `json:"rule,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

func (resource Cmppolicy) ToUpdate() CmppolicyUpdate {
	update := CmppolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Resaction,
	}

	return update
}

type update_cmppolicy_payload struct {
	Update CmppolicyUpdate `json:"cmppolicy"`
}

func (c *NitroClient) UpdateCmppolicy(update CmppolicyUpdate) error {
	payload := update_cmppolicy_payload{
		update,
	}

	return c.put("cmppolicy", "", nil, payload)
}

//      UNSET

type CmppolicyUnset struct {
	Name      string `json:"name,omitempty"`
	Rule      bool   `json:"rule,omitempty"`
	Resaction bool   `json:"resaction,omitempty"`
}

func (resource Cmppolicy) ToUnset() CmppolicyUnset {
	unset := CmppolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_cmppolicy_payload struct {
	Unset CmppolicyUnset `json:"cmppolicy"`
}

func (c *NitroClient) UnsetCmppolicy(unset CmppolicyUnset) error {
	payload := unset_cmppolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("cmppolicy", "", qs, payload)
}

//      RENAME
//      TODO
