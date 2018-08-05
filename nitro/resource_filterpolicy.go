package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Filterpolicy struct {
	Name      string `json:"name"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

func filterpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type FilterpolicyUnset struct {
	Name      string `json:"name"`
	Rule      bool   `json:"rule,string,omitempty"`
	Reqaction bool   `json:"reqaction,string,omitempty"`
	Resaction bool   `json:"resaction,string,omitempty"`
}

type update_filterpolicy struct {
	Name      string `json:"name"`
	Rule      string `json:"rule,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

type rename_filterpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_filterpolicy_payload struct {
	Resource Filterpolicy `json:"filterpolicy"`
}

type rename_filterpolicy_payload struct {
	Rename rename_filterpolicy `json:"filterpolicy"`
}

type unset_filterpolicy_payload struct {
	Unset FilterpolicyUnset `json:"filterpolicy"`
}

type update_filterpolicy_payload struct {
	Update update_filterpolicy `json:"filterpolicy"`
}

type get_filterpolicy_result struct {
	Results []Filterpolicy `json:"filterpolicy"`
}

type count_filterpolicy_result struct {
	Results []Count `json:"filterpolicy"`
}

func (c *NitroClient) AddFilterpolicy(resource Filterpolicy) error {
	payload := add_filterpolicy_payload{
		resource,
	}

	return c.post("filterpolicy", "", nil, payload)
}

func (c *NitroClient) RenameFilterpolicy(name string, newName string) error {
	payload := rename_filterpolicy_payload{
		rename_filterpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("filterpolicy", "", qs, payload)
}

func (c *NitroClient) CountFilterpolicy() (int, error) {
	var results count_filterpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("filterpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsFilterpolicy(key string) (bool, error) {
	var results count_filterpolicy_result

	id, qs := filterpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("filterpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListFilterpolicy() ([]Filterpolicy, error) {
	results := get_filterpolicy_result{}

	if err := c.get("filterpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetFilterpolicy(key string) (*Filterpolicy, error) {
	var results get_filterpolicy_result

	id, qs := filterpolicy_key_to_id_args(key)

	if err := c.get("filterpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one filterpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("filterpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteFilterpolicy(key string) error {
	id, qs := filterpolicy_key_to_id_args(key)

	return c.delete("filterpolicy", id, qs)
}

func (c *NitroClient) UnsetFilterpolicy(unset FilterpolicyUnset) error {
	payload := unset_filterpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("filterpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateFilterpolicy(resource Filterpolicy) error {
	payload := update_filterpolicy_payload{
		update_filterpolicy{
			resource.Name,
			resource.Rule,
			resource.Reqaction,
			resource.Resaction,
		},
	}

	return c.put("filterpolicy", "", nil, payload)
}
