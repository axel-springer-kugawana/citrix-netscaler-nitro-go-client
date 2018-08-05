package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appfwpolicy struct {
	Name        string `json:"name"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

func appfwpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppfwpolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,string,omitempty"`
	Profilename bool   `json:"profilename,string,omitempty"`
	Comment     bool   `json:"comment,string,omitempty"`
	Logaction   bool   `json:"logaction,string,omitempty"`
}

type update_appfwpolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type rename_appfwpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appfwpolicy_payload struct {
	Resource Appfwpolicy `json:"appfwpolicy"`
}

type rename_appfwpolicy_payload struct {
	Rename rename_appfwpolicy `json:"appfwpolicy"`
}

type unset_appfwpolicy_payload struct {
	Unset AppfwpolicyUnset `json:"appfwpolicy"`
}

type update_appfwpolicy_payload struct {
	Update update_appfwpolicy `json:"appfwpolicy"`
}

type get_appfwpolicy_result struct {
	Results []Appfwpolicy `json:"appfwpolicy"`
}

type count_appfwpolicy_result struct {
	Results []Count `json:"appfwpolicy"`
}

func (c *NitroClient) AddAppfwpolicy(resource Appfwpolicy) error {
	payload := add_appfwpolicy_payload{
		resource,
	}

	return c.post("appfwpolicy", "", nil, payload)
}

func (c *NitroClient) RenameAppfwpolicy(name string, newName string) error {
	payload := rename_appfwpolicy_payload{
		rename_appfwpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appfwpolicy", "", qs, payload)
}

func (c *NitroClient) CountAppfwpolicy() (int, error) {
	var results count_appfwpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appfwpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppfwpolicy(key string) (bool, error) {
	var results count_appfwpolicy_result

	id, qs := appfwpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appfwpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppfwpolicy() ([]Appfwpolicy, error) {
	results := get_appfwpolicy_result{}

	if err := c.get("appfwpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppfwpolicy(key string) (*Appfwpolicy, error) {
	var results get_appfwpolicy_result

	id, qs := appfwpolicy_key_to_id_args(key)

	if err := c.get("appfwpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appfwpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appfwpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppfwpolicy(key string) error {
	id, qs := appfwpolicy_key_to_id_args(key)

	return c.delete("appfwpolicy", id, qs)
}

func (c *NitroClient) UnsetAppfwpolicy(unset AppfwpolicyUnset) error {
	payload := unset_appfwpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appfwpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAppfwpolicy(resource Appfwpolicy) error {
	payload := update_appfwpolicy_payload{
		update_appfwpolicy{
			resource.Name,
			resource.Rule,
			resource.Profilename,
			resource.Comment,
			resource.Logaction,
		},
	}

	return c.put("appfwpolicy", "", nil, payload)
}
