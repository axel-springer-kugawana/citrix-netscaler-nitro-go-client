package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Filteraction struct {
	Name        string `json:"name"`
	Page        string `json:"page,omitempty"`
	Qual        string `json:"qual,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
}

type FilteractionKey struct {
	Name string `json:"name"`
}

type FilteractionUnset struct {
	Name        string `json:"name"`
	Servicename bool   `json:"servicename,string,omitempty"`
	Value       bool   `json:"value,string,omitempty"`
	Respcode    bool   `json:"respcode,string,omitempty"`
	Page        bool   `json:"page,string,omitempty"`
}

type update_filteraction struct {
	Name        string `json:"name"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Page        string `json:"page,omitempty"`
}

type rename_filteraction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_filteraction_payload struct {
	Resource Filteraction `json:"filteraction"`
}

type rename_filteraction_payload struct {
	Rename rename_filteraction `json:"filteraction"`
}

type unset_filteraction_payload struct {
	Unset FilteractionUnset `json:"filteraction"`
}

type update_filteraction_payload struct {
	Update update_filteraction `json:"filteraction"`
}

type get_filteraction_result struct {
	Results []Filteraction `json:"filteraction"`
}

type count_filteraction_result struct {
	Results []Count `json:"filteraction"`
}

func filteraction_key_to_id_args(key FilteractionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddFilteraction(resource Filteraction) error {
	payload := add_filteraction_payload{
		resource,
	}

	return c.post("filteraction", "", nil, payload)
}

func (c *NitroClient) RenameFilteraction(name string, newName string) error {
	payload := rename_filteraction_payload{
		rename_filteraction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("filteraction", "", qs, payload)
}

func (c *NitroClient) CountFilteraction() (int, error) {
	var results count_filteraction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("filteraction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsFilteraction(key FilteractionKey) (bool, error) {
	var results count_filteraction_result

	id, qs := filteraction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("filteraction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListFilteraction() ([]Filteraction, error) {
	var results get_filteraction_result

	if err := c.get("filteraction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetFilteraction(key FilteractionKey) (*Filteraction, error) {
	var results get_filteraction_result

	id, qs := filteraction_key_to_id_args(key)

	if err := c.get("filteraction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one filteraction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("filteraction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteFilteraction(key FilteractionKey) error {
	id, qs := filteraction_key_to_id_args(key)

	return c.delete("filteraction", id, qs)
}

func (c *NitroClient) UnsetFilteraction(unset FilteractionUnset) error {
	payload := unset_filteraction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("filteraction", "", qs, payload)
}

func (c *NitroClient) UpdateFilteraction(resource Filteraction) error {
	payload := update_filteraction_payload{
		update_filteraction{
			resource.Name,
			resource.Servicename,
			resource.Value,
			resource.Respcode,
			resource.Page,
		},
	}

	return c.put("filteraction", "", nil, payload)
}
