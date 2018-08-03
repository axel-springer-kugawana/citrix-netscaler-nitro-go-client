package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowcollector struct {
	Name       string `json:"name"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
	Port       int    `json:"port,omitempty"`
	Transport  string `json:"transport,omitempty"`
}

func appflowcollector_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppflowcollectorUnset struct {
	Name       string `json:"name"`
	Ipaddress  bool   `json:"ipaddress,string,omitempty"`
	Port       bool   `json:"port,string,omitempty"`
	Netprofile bool   `json:"netprofile,string,omitempty"`
}

type update_appflowcollector struct {
	Name       string `json:"name"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Port       int    `json:"port,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
}

type rename_appflowcollector struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appflowcollector_payload struct {
	Resource Appflowcollector `json:"appflowcollector"`
}

type rename_appflowcollector_payload struct {
	Rename rename_appflowcollector `json:"appflowcollector"`
}

type unset_appflowcollector_payload struct {
	Unset AppflowcollectorUnset `json:"appflowcollector"`
}

type update_appflowcollector_payload struct {
	Update update_appflowcollector `json:"appflowcollector"`
}

type get_appflowcollector_result struct {
	Results []Appflowcollector `json:"appflowcollector"`
}

type count_appflowcollector_result struct {
	Results []Count `json:"appflowcollector"`
}

func (c *NitroClient) AddAppflowcollector(resource Appflowcollector) error {
	payload := add_appflowcollector_payload{
		resource,
	}

	return c.post("appflowcollector", "", nil, payload)
}

func (c *NitroClient) RenameAppflowcollector(name string, newName string) error {
	payload := rename_appflowcollector_payload{
		rename_appflowcollector{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appflowcollector", "", qs, payload)
}

func (c *NitroClient) CountAppflowcollector() (int, error) {
	var results count_appflowcollector_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowcollector", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppflowcollector(key string) (bool, error) {
	var results count_appflowcollector_result

	id, qs := appflowcollector_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appflowcollector", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppflowcollector() ([]Appflowcollector, error) {
	results := get_appflowcollector_result{}

	if err := c.get("appflowcollector", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowcollector(key string) (*Appflowcollector, error) {
	var results get_appflowcollector_result

	id, qs := appflowcollector_key_to_id_args(key)

	if err := c.get("appflowcollector", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowcollector element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("appflowcollector element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowcollector(key string) error {
	id, qs := appflowcollector_key_to_id_args(key)

	return c.delete("appflowcollector", id, qs)
}

func (c *NitroClient) UnsetAppflowcollector(unset AppflowcollectorUnset) error {
	payload := unset_appflowcollector_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appflowcollector", "", qs, payload)
}

func (c *NitroClient) UpdateAppflowcollector(resource Appflowcollector) error {
	payload := update_appflowcollector_payload{
		update_appflowcollector{
			resource.Name,
			resource.Ipaddress,
			resource.Port,
			resource.Netprofile,
		},
	}

	return c.put("appflowcollector", "", nil, payload)
}
