package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowpolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

func appflowpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppflowpolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,string,omitempty"`
	Action      bool   `json:"action,string,omitempty"`
	Undefaction bool   `json:"undefaction,string,omitempty"`
	Comment     bool   `json:"comment,string,omitempty"`
}

type update_appflowpolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type rename_appflowpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appflowpolicy_payload struct {
	Resource Appflowpolicy `json:"appflowpolicy"`
}

type rename_appflowpolicy_payload struct {
	Rename rename_appflowpolicy `json:"appflowpolicy"`
}

type unset_appflowpolicy_payload struct {
	Unset AppflowpolicyUnset `json:"appflowpolicy"`
}

type update_appflowpolicy_payload struct {
	Update update_appflowpolicy `json:"appflowpolicy"`
}

type get_appflowpolicy_result struct {
	Results []Appflowpolicy `json:"appflowpolicy"`
}

type count_appflowpolicy_result struct {
	Results []Count `json:"appflowpolicy"`
}

func (c *NitroClient) AddAppflowpolicy(resource Appflowpolicy) error {
	payload := add_appflowpolicy_payload{
		resource,
	}

	return c.post("appflowpolicy", "", nil, payload)
}

func (c *NitroClient) RenameAppflowpolicy(name string, newName string) error {
	payload := rename_appflowpolicy_payload{
		rename_appflowpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appflowpolicy", "", qs, payload)
}

func (c *NitroClient) CountAppflowpolicy() (int, error) {
	var results count_appflowpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppflowpolicy(key string) (bool, error) {
	var results count_appflowpolicy_result

	id, qs := appflowpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appflowpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppflowpolicy() ([]Appflowpolicy, error) {
	results := get_appflowpolicy_result{}

	if err := c.get("appflowpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowpolicy(key string) (*Appflowpolicy, error) {
	var results get_appflowpolicy_result

	id, qs := appflowpolicy_key_to_id_args(key)

	if err := c.get("appflowpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowpolicy(key string) error {
	id, qs := appflowpolicy_key_to_id_args(key)

	return c.delete("appflowpolicy", id, qs)
}

func (c *NitroClient) UnsetAppflowpolicy(unset AppflowpolicyUnset) error {
	payload := unset_appflowpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appflowpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAppflowpolicy(resource Appflowpolicy) error {
	payload := update_appflowpolicy_payload{
		update_appflowpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Undefaction,
			resource.Comment,
		},
	}

	return c.put("appflowpolicy", "", nil, payload)
}
