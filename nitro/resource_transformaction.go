package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformaction struct {
	Name             string `json:"name"`
	Comment          string `json:"comment,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Priority         int    `json:"priority,string,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	State            string `json:"state,omitempty"`
}

type TransformactionKey struct {
	Name string `json:"name"`
}

type TransformactionUnset struct {
	Name             string `json:"name"`
	Priority         bool   `json:"priority,string,omitempty"`
	Requrlfrom       bool   `json:"requrlfrom,string,omitempty"`
	Requrlinto       bool   `json:"requrlinto,string,omitempty"`
	Resurlfrom       bool   `json:"resurlfrom,string,omitempty"`
	Resurlinto       bool   `json:"resurlinto,string,omitempty"`
	Cookiedomainfrom bool   `json:"cookiedomainfrom,string,omitempty"`
	Cookiedomaininto bool   `json:"cookiedomaininto,string,omitempty"`
	State            bool   `json:"state,string,omitempty"`
	Comment          bool   `json:"comment,string,omitempty"`
}

type update_transformaction struct {
	Name             string `json:"name"`
	Priority         int    `json:"priority,string,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	State            string `json:"state,omitempty"`
	Comment          string `json:"comment,omitempty"`
}

type rename_transformaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_transformaction_payload struct {
	Resource Transformaction `json:"transformaction"`
}

type rename_transformaction_payload struct {
	Rename rename_transformaction `json:"transformaction"`
}

type unset_transformaction_payload struct {
	Unset TransformactionUnset `json:"transformaction"`
}

type update_transformaction_payload struct {
	Update update_transformaction `json:"transformaction"`
}

type get_transformaction_result struct {
	Results []Transformaction `json:"transformaction"`
}

type count_transformaction_result struct {
	Results []Count `json:"transformaction"`
}

func transformaction_key_to_id_args(key TransformactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddTransformaction(resource Transformaction) error {
	payload := add_transformaction_payload{
		resource,
	}

	return c.post("transformaction", "", nil, payload)
}

func (c *NitroClient) RenameTransformaction(name string, newName string) error {
	payload := rename_transformaction_payload{
		rename_transformaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("transformaction", "", qs, payload)
}

func (c *NitroClient) CountTransformaction() (int, error) {
	var results count_transformaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTransformaction(key TransformactionKey) (bool, error) {
	var results count_transformaction_result

	id, qs := transformaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("transformaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTransformaction() ([]Transformaction, error) {
	var results get_transformaction_result

	if err := c.get("transformaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTransformaction(key TransformactionKey) (*Transformaction, error) {
	var results get_transformaction_result

	id, qs := transformaction_key_to_id_args(key)

	if err := c.get("transformaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("transformaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTransformaction(key TransformactionKey) error {
	id, qs := transformaction_key_to_id_args(key)

	return c.delete("transformaction", id, qs)
}

func (c *NitroClient) UnsetTransformaction(unset TransformactionUnset) error {
	payload := unset_transformaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("transformaction", "", qs, payload)
}

func (c *NitroClient) UpdateTransformaction(resource Transformaction) error {
	payload := update_transformaction_payload{
		update_transformaction{
			resource.Name,
			resource.Priority,
			resource.Requrlfrom,
			resource.Requrlinto,
			resource.Resurlfrom,
			resource.Resurlinto,
			resource.Cookiedomainfrom,
			resource.Cookiedomaininto,
			resource.State,
			resource.Comment,
		},
	}

	return c.put("transformaction", "", nil, payload)
}
