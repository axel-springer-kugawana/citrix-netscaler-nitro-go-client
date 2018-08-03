package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmppolicy struct {
	Name      string `json:"name"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

type CmppolicyKey struct {
	Name string `json:"name"`
}

type CmppolicyUnset struct {
	Name      string `json:"name"`
	Rule      bool   `json:"rule,string,omitempty"`
	Resaction bool   `json:"resaction,string,omitempty"`
}

type update_cmppolicy struct {
	Name      string `json:"name"`
	Rule      string `json:"rule,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

type rename_cmppolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_cmppolicy_payload struct {
	Resource Cmppolicy `json:"cmppolicy"`
}

type rename_cmppolicy_payload struct {
	Rename rename_cmppolicy `json:"cmppolicy"`
}

type unset_cmppolicy_payload struct {
	Unset CmppolicyUnset `json:"cmppolicy"`
}

type update_cmppolicy_payload struct {
	Update update_cmppolicy `json:"cmppolicy"`
}

type get_cmppolicy_result struct {
	Results []Cmppolicy `json:"cmppolicy"`
}

type count_cmppolicy_result struct {
	Results []Count `json:"cmppolicy"`
}

func cmppolicy_key_to_id_args(key CmppolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddCmppolicy(resource Cmppolicy) error {
	payload := add_cmppolicy_payload{
		resource,
	}

	return c.post("cmppolicy", "", nil, payload)
}

func (c *NitroClient) RenameCmppolicy(name string, newName string) error {
	payload := rename_cmppolicy_payload{
		rename_cmppolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cmppolicy", "", qs, payload)
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

func (c *NitroClient) ExistsCmppolicy(key CmppolicyKey) (bool, error) {
	var results count_cmppolicy_result

	id, qs := cmppolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cmppolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCmppolicy() ([]Cmppolicy, error) {
	var results get_cmppolicy_result

	if err := c.get("cmppolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCmppolicy(key CmppolicyKey) (*Cmppolicy, error) {
	var results get_cmppolicy_result

	id, qs := cmppolicy_key_to_id_args(key)

	if err := c.get("cmppolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmppolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("cmppolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCmppolicy(key CmppolicyKey) error {
	id, qs := cmppolicy_key_to_id_args(key)

	return c.delete("cmppolicy", id, qs)
}

func (c *NitroClient) UnsetCmppolicy(unset CmppolicyUnset) error {
	payload := unset_cmppolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("cmppolicy", "", qs, payload)
}

func (c *NitroClient) UpdateCmppolicy(resource Cmppolicy) error {
	payload := update_cmppolicy_payload{
		update_cmppolicy{
			resource.Name,
			resource.Rule,
			resource.Resaction,
		},
	}

	return c.put("cmppolicy", "", nil, payload)
}
