package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Responderpolicy struct {
	Name          string `json:"name"`
	Action        string `json:"action,omitempty"`
	Appflowaction string `json:"appflowaction,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Rule          string `json:"rule,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
}

type ResponderpolicyKey struct {
	Name string `json:"name"`
}

type ResponderpolicyUnset struct {
	Name          string `json:"name"`
	Rule          bool   `json:"rule,string,omitempty"`
	Action        bool   `json:"action,string,omitempty"`
	Undefaction   bool   `json:"undefaction,string,omitempty"`
	Comment       bool   `json:"comment,string,omitempty"`
	Logaction     bool   `json:"logaction,string,omitempty"`
	Appflowaction bool   `json:"appflowaction,string,omitempty"`
}

type update_responderpolicy struct {
	Name          string `json:"name"`
	Rule          string `json:"rule,omitempty"`
	Action        string `json:"action,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Appflowaction string `json:"appflowaction,omitempty"`
}

type rename_responderpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_responderpolicy_payload struct {
	Resource Responderpolicy `json:"responderpolicy"`
}

type rename_responderpolicy_payload struct {
	Rename rename_responderpolicy `json:"responderpolicy"`
}

type unset_responderpolicy_payload struct {
	Unset ResponderpolicyUnset `json:"responderpolicy"`
}

type update_responderpolicy_payload struct {
	Update update_responderpolicy `json:"responderpolicy"`
}

type get_responderpolicy_result struct {
	Results []Responderpolicy `json:"responderpolicy"`
}

type count_responderpolicy_result struct {
	Results []Count `json:"responderpolicy"`
}

func responderpolicy_key_to_id_args(key ResponderpolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddResponderpolicy(resource Responderpolicy) error {
	payload := add_responderpolicy_payload{
		resource,
	}

	return c.post("responderpolicy", "", nil, payload)
}

func (c *NitroClient) RenameResponderpolicy(name string, newName string) error {
	payload := rename_responderpolicy_payload{
		rename_responderpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("responderpolicy", "", qs, payload)
}

func (c *NitroClient) CountResponderpolicy() (int, error) {
	var results count_responderpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("responderpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsResponderpolicy(key ResponderpolicyKey) (bool, error) {
	var results count_responderpolicy_result

	id, qs := responderpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("responderpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListResponderpolicy() ([]Responderpolicy, error) {
	var results get_responderpolicy_result

	if err := c.get("responderpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetResponderpolicy(key ResponderpolicyKey) (*Responderpolicy, error) {
	var results get_responderpolicy_result

	id, qs := responderpolicy_key_to_id_args(key)

	if err := c.get("responderpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one responderpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("responderpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteResponderpolicy(key ResponderpolicyKey) error {
	id, qs := responderpolicy_key_to_id_args(key)

	return c.delete("responderpolicy", id, qs)
}

func (c *NitroClient) UnsetResponderpolicy(unset ResponderpolicyUnset) error {
	payload := unset_responderpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("responderpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateResponderpolicy(resource Responderpolicy) error {
	payload := update_responderpolicy_payload{
		update_responderpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Undefaction,
			resource.Comment,
			resource.Logaction,
			resource.Appflowaction,
		},
	}

	return c.put("responderpolicy", "", nil, payload)
}
