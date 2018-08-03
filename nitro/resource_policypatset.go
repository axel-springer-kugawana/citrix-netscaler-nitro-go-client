package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policypatset struct {
	Name      string `json:"name"`
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
}

type PolicypatsetKey struct {
	Name string `json:"name"`
}

type rename_policypatset struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_policypatset_payload struct {
	Resource Policypatset `json:"policypatset"`
}

type rename_policypatset_payload struct {
	Rename rename_policypatset `json:"policypatset"`
}

type get_policypatset_result struct {
	Results []Policypatset `json:"policypatset"`
}

type count_policypatset_result struct {
	Results []Count `json:"policypatset"`
}

func policypatset_key_to_id_args(key PolicypatsetKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddPolicypatset(resource Policypatset) error {
	payload := add_policypatset_payload{
		resource,
	}

	return c.post("policypatset", "", nil, payload)
}

func (c *NitroClient) RenamePolicypatset(name string, newName string) error {
	payload := rename_policypatset_payload{
		rename_policypatset{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("policypatset", "", qs, payload)
}

func (c *NitroClient) CountPolicypatset() (int, error) {
	var results count_policypatset_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policypatset", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicypatset(key PolicypatsetKey) (bool, error) {
	var results count_policypatset_result

	id, qs := policypatset_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("policypatset", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListPolicypatset() ([]Policypatset, error) {
	var results get_policypatset_result

	if err := c.get("policypatset", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicypatset(key PolicypatsetKey) (*Policypatset, error) {
	var results get_policypatset_result

	id, qs := policypatset_key_to_id_args(key)

	if err := c.get("policypatset", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policypatset element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policypatset element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicypatset(key PolicypatsetKey) error {
	id, qs := policypatset_key_to_id_args(key)

	return c.delete("policypatset", id, qs)
}
