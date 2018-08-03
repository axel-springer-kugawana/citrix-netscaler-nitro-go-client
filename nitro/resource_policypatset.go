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

func policypatset_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
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

func (c *NitroClient) ExistsPolicypatset(key string) (bool, error) {
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
	results := get_policypatset_result{}

	if err := c.get("policypatset", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicypatset(key string) (*Policypatset, error) {
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

func (c *NitroClient) DeletePolicypatset(key string) error {
	id, qs := policypatset_key_to_id_args(key)

	return c.delete("policypatset", id, qs)
}
