package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmtrafficaction struct {
	Name             string `json:"name"`
	Apptimeout       int    `json:"apptimeout,string,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,string,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
}

type TmtrafficactionKey struct {
	Name string `json:"name"`
}

type TmtrafficactionUnset struct {
	Name             string `json:"name"`
	Apptimeout       bool   `json:"apptimeout,string,omitempty"`
	Sso              bool   `json:"sso,string,omitempty"`
	Formssoaction    bool   `json:"formssoaction,string,omitempty"`
	Persistentcookie bool   `json:"persistentcookie,string,omitempty"`
	Initiatelogout   bool   `json:"initiatelogout,string,omitempty"`
	Kcdaccount       bool   `json:"kcdaccount,string,omitempty"`
	Samlssoprofile   bool   `json:"samlssoprofile,string,omitempty"`
	Forcedtimeout    bool   `json:"forcedtimeout,string,omitempty"`
	Forcedtimeoutval bool   `json:"forcedtimeoutval,string,omitempty"`
	Userexpression   bool   `json:"userexpression,string,omitempty"`
	Passwdexpression bool   `json:"passwdexpression,string,omitempty"`
}

type update_tmtrafficaction struct {
	Name             string `json:"name"`
	Apptimeout       int    `json:"apptimeout,string,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,string,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
}

type rename_tmtrafficaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_tmtrafficaction_payload struct {
	Resource Tmtrafficaction `json:"tmtrafficaction"`
}

type rename_tmtrafficaction_payload struct {
	Rename rename_tmtrafficaction `json:"tmtrafficaction"`
}

type unset_tmtrafficaction_payload struct {
	Unset TmtrafficactionUnset `json:"tmtrafficaction"`
}

type update_tmtrafficaction_payload struct {
	Update update_tmtrafficaction `json:"tmtrafficaction"`
}

type get_tmtrafficaction_result struct {
	Results []Tmtrafficaction `json:"tmtrafficaction"`
}

type count_tmtrafficaction_result struct {
	Results []Count `json:"tmtrafficaction"`
}

func tmtrafficaction_key_to_id_args(key TmtrafficactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddTmtrafficaction(resource Tmtrafficaction) error {
	payload := add_tmtrafficaction_payload{
		resource,
	}

	return c.post("tmtrafficaction", "", nil, payload)
}

func (c *NitroClient) RenameTmtrafficaction(name string, newName string) error {
	payload := rename_tmtrafficaction_payload{
		rename_tmtrafficaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("tmtrafficaction", "", qs, payload)
}

func (c *NitroClient) CountTmtrafficaction() (int, error) {
	var results count_tmtrafficaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmtrafficaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTmtrafficaction(key TmtrafficactionKey) (bool, error) {
	var results count_tmtrafficaction_result

	id, qs := tmtrafficaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("tmtrafficaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTmtrafficaction() ([]Tmtrafficaction, error) {
	var results get_tmtrafficaction_result

	if err := c.get("tmtrafficaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTmtrafficaction(key TmtrafficactionKey) (*Tmtrafficaction, error) {
	var results get_tmtrafficaction_result

	id, qs := tmtrafficaction_key_to_id_args(key)

	if err := c.get("tmtrafficaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmtrafficaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("tmtrafficaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTmtrafficaction(key TmtrafficactionKey) error {
	id, qs := tmtrafficaction_key_to_id_args(key)

	return c.delete("tmtrafficaction", id, qs)
}

func (c *NitroClient) UnsetTmtrafficaction(unset TmtrafficactionUnset) error {
	payload := unset_tmtrafficaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("tmtrafficaction", "", qs, payload)
}

func (c *NitroClient) UpdateTmtrafficaction(resource Tmtrafficaction) error {
	payload := update_tmtrafficaction_payload{
		update_tmtrafficaction{
			resource.Name,
			resource.Apptimeout,
			resource.Sso,
			resource.Formssoaction,
			resource.Persistentcookie,
			resource.Initiatelogout,
			resource.Kcdaccount,
			resource.Samlssoprofile,
			resource.Forcedtimeout,
			resource.Forcedtimeoutval,
			resource.Userexpression,
			resource.Passwdexpression,
		},
	}

	return c.put("tmtrafficaction", "", nil, payload)
}
