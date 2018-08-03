package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policydataset struct {
	Name      string `json:"name"`
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
	Type      string `json:"type,omitempty"`
}

type PolicydatasetKey struct {
	Name string `json:"name"`
}

type rename_policydataset struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_policydataset_payload struct {
	Resource Policydataset `json:"policydataset"`
}

type rename_policydataset_payload struct {
	Rename rename_policydataset `json:"policydataset"`
}

type get_policydataset_result struct {
	Results []Policydataset `json:"policydataset"`
}

type count_policydataset_result struct {
	Results []Count `json:"policydataset"`
}

func policydataset_key_to_id_args(key PolicydatasetKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddPolicydataset(resource Policydataset) error {
	payload := add_policydataset_payload{
		resource,
	}

	return c.post("policydataset", "", nil, payload)
}

func (c *NitroClient) RenamePolicydataset(name string, newName string) error {
	payload := rename_policydataset_payload{
		rename_policydataset{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("policydataset", "", qs, payload)
}

func (c *NitroClient) CountPolicydataset() (int, error) {
	var results count_policydataset_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policydataset", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicydataset(key PolicydatasetKey) (bool, error) {
	var results count_policydataset_result

	id, qs := policydataset_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("policydataset", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListPolicydataset() ([]Policydataset, error) {
	var results get_policydataset_result

	if err := c.get("policydataset", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicydataset(key PolicydatasetKey) (*Policydataset, error) {
	var results get_policydataset_result

	id, qs := policydataset_key_to_id_args(key)

	if err := c.get("policydataset", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policydataset element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policydataset element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicydataset(key PolicydatasetKey) error {
	id, qs := policydataset_key_to_id_args(key)

	return c.delete("policydataset", id, qs)
}
