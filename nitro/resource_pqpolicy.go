package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Pqpolicy struct {
	Policyname string `json:"policyname"`
	Polqdepth  int    `json:"polqdepth,string,omitempty"`
	Priority   int    `json:"priority,string,omitempty"`
	Qdepth     int    `json:"qdepth,string,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Weight     int    `json:"weight,string,omitempty"`
}

type PqpolicyKey struct {
	Policyname string `json:"policyname"`
}

type PqpolicyUnset struct {
	Policyname string `json:"policyname"`
	Weight     bool   `json:"weight,string,omitempty"`
	Qdepth     bool   `json:"qdepth,string,omitempty"`
	Polqdepth  bool   `json:"polqdepth,string,omitempty"`
}

type update_pqpolicy struct {
	Policyname string `json:"policyname"`
	Weight     int    `json:"weight,string,omitempty"`
	Qdepth     int    `json:"qdepth,string,omitempty"`
	Polqdepth  int    `json:"polqdepth,string,omitempty"`
}

type rename_pqpolicy struct {
	Name    string `json:"policyname"`
	Newname string `json:"newname"`
}

type add_pqpolicy_payload struct {
	Resource Pqpolicy `json:"pqpolicy"`
}

type rename_pqpolicy_payload struct {
	Rename rename_pqpolicy `json:"pqpolicy"`
}

type unset_pqpolicy_payload struct {
	Unset PqpolicyUnset `json:"pqpolicy"`
}

type update_pqpolicy_payload struct {
	Update update_pqpolicy `json:"pqpolicy"`
}

type get_pqpolicy_result struct {
	Results []Pqpolicy `json:"pqpolicy"`
}

type count_pqpolicy_result struct {
	Results []Count `json:"pqpolicy"`
}

func pqpolicy_key_to_id_args(key PqpolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Policyname, qs
}

func (c *NitroClient) AddPqpolicy(resource Pqpolicy) error {
	payload := add_pqpolicy_payload{
		resource,
	}

	return c.post("pqpolicy", "", nil, payload)
}

func (c *NitroClient) RenamePqpolicy(name string, newName string) error {
	payload := rename_pqpolicy_payload{
		rename_pqpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("pqpolicy", "", qs, payload)
}

func (c *NitroClient) CountPqpolicy() (int, error) {
	var results count_pqpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("pqpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPqpolicy(key PqpolicyKey) (bool, error) {
	var results count_pqpolicy_result

	id, qs := pqpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("pqpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListPqpolicy() ([]Pqpolicy, error) {
	var results get_pqpolicy_result

	if err := c.get("pqpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPqpolicy(key PqpolicyKey) (*Pqpolicy, error) {
	var results get_pqpolicy_result

	id, qs := pqpolicy_key_to_id_args(key)

	if err := c.get("pqpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one pqpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("pqpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePqpolicy(key PqpolicyKey) error {
	id, qs := pqpolicy_key_to_id_args(key)

	return c.delete("pqpolicy", id, qs)
}

func (c *NitroClient) UnsetPqpolicy(unset PqpolicyUnset) error {
	payload := unset_pqpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("pqpolicy", "", qs, payload)
}

func (c *NitroClient) UpdatePqpolicy(resource Pqpolicy) error {
	payload := update_pqpolicy_payload{
		update_pqpolicy{
			resource.Policyname,
			resource.Weight,
			resource.Qdepth,
			resource.Polqdepth,
		},
	}

	return c.put("pqpolicy", "", nil, payload)
}
