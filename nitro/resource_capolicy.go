package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Capolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

func capolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type CapolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
}

type update_capolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type rename_capolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_capolicy_payload struct {
	Resource Capolicy `json:"capolicy"`
}

type rename_capolicy_payload struct {
	Rename rename_capolicy `json:"capolicy"`
}

type unset_capolicy_payload struct {
	Unset CapolicyUnset `json:"capolicy"`
}

type update_capolicy_payload struct {
	Update update_capolicy `json:"capolicy"`
}

type get_capolicy_result struct {
	Results []Capolicy `json:"capolicy"`
}

type count_capolicy_result struct {
	Results []Count `json:"capolicy"`
}

func (c *NitroClient) AddCapolicy(resource Capolicy) error {
	payload := add_capolicy_payload{
		resource,
	}

	return c.post("capolicy", "", nil, payload)
}

func (c *NitroClient) RenameCapolicy(name string, newName string) error {
	payload := rename_capolicy_payload{
		rename_capolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("capolicy", "", qs, payload)
}

func (c *NitroClient) CountCapolicy() (int, error) {
	var results count_capolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("capolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCapolicy(key string) (bool, error) {
	var results count_capolicy_result

	id, qs := capolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("capolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCapolicy() ([]Capolicy, error) {
	results := get_capolicy_result{}

	if err := c.get("capolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCapolicy(key string) (*Capolicy, error) {
	var results get_capolicy_result

	id, qs := capolicy_key_to_id_args(key)

	if err := c.get("capolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one capolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("capolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCapolicy(key string) error {
	id, qs := capolicy_key_to_id_args(key)

	return c.delete("capolicy", id, qs)
}

func (c *NitroClient) UnsetCapolicy(unset CapolicyUnset) error {
	payload := unset_capolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("capolicy", "", qs, payload)
}

func (c *NitroClient) UpdateCapolicy(resource Capolicy) error {
	payload := update_capolicy_payload{
		update_capolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Comment,
			resource.Logaction,
			resource.Undefaction,
		},
	}

	return c.put("capolicy", "", nil, payload)
}
