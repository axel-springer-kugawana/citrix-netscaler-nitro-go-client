package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformpolicy struct {
	Name        string `json:"name"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

type TransformpolicyKey struct {
	Name string `json:"name"`
}

type TransformpolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,string,omitempty"`
	Profilename bool   `json:"profilename,string,omitempty"`
	Comment     bool   `json:"comment,string,omitempty"`
	Logaction   bool   `json:"logaction,string,omitempty"`
}

type update_transformpolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type rename_transformpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_transformpolicy_payload struct {
	Resource Transformpolicy `json:"transformpolicy"`
}

type rename_transformpolicy_payload struct {
	Rename rename_transformpolicy `json:"transformpolicy"`
}

type unset_transformpolicy_payload struct {
	Unset TransformpolicyUnset `json:"transformpolicy"`
}

type update_transformpolicy_payload struct {
	Update update_transformpolicy `json:"transformpolicy"`
}

type get_transformpolicy_result struct {
	Results []Transformpolicy `json:"transformpolicy"`
}

type count_transformpolicy_result struct {
	Results []Count `json:"transformpolicy"`
}

func transformpolicy_key_to_id_args(key TransformpolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddTransformpolicy(resource Transformpolicy) error {
	payload := add_transformpolicy_payload{
		resource,
	}

	return c.post("transformpolicy", "", nil, payload)
}

func (c *NitroClient) RenameTransformpolicy(name string, newName string) error {
	payload := rename_transformpolicy_payload{
		rename_transformpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("transformpolicy", "", qs, payload)
}

func (c *NitroClient) CountTransformpolicy() (int, error) {
	var results count_transformpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTransformpolicy(key TransformpolicyKey) (bool, error) {
	var results count_transformpolicy_result

	id, qs := transformpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("transformpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTransformpolicy() ([]Transformpolicy, error) {
	var results get_transformpolicy_result

	if err := c.get("transformpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTransformpolicy(key TransformpolicyKey) (*Transformpolicy, error) {
	var results get_transformpolicy_result

	id, qs := transformpolicy_key_to_id_args(key)

	if err := c.get("transformpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("transformpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTransformpolicy(key TransformpolicyKey) error {
	id, qs := transformpolicy_key_to_id_args(key)

	return c.delete("transformpolicy", id, qs)
}

func (c *NitroClient) UnsetTransformpolicy(unset TransformpolicyUnset) error {
	payload := unset_transformpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("transformpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateTransformpolicy(resource Transformpolicy) error {
	payload := update_transformpolicy_payload{
		update_transformpolicy{
			resource.Name,
			resource.Rule,
			resource.Profilename,
			resource.Comment,
			resource.Logaction,
		},
	}

	return c.put("transformpolicy", "", nil, payload)
}
