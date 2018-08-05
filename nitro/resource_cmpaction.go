package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmpaction struct {
	Name            string `json:"name"`
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Cmptype         string `json:"cmptype,omitempty"`
	Deltatype       string `json:"deltatype,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

func cmpaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type CmpactionUnset struct {
	Name            string `json:"name"`
	Cmptype         bool   `json:"cmptype,string,omitempty"`
	Addvaryheader   bool   `json:"addvaryheader,string,omitempty"`
	Varyheadervalue bool   `json:"varyheadervalue,string,omitempty"`
}

type update_cmpaction struct {
	Name            string `json:"name"`
	Cmptype         string `json:"cmptype,omitempty"`
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

type rename_cmpaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_cmpaction_payload struct {
	Resource Cmpaction `json:"cmpaction"`
}

type rename_cmpaction_payload struct {
	Rename rename_cmpaction `json:"cmpaction"`
}

type unset_cmpaction_payload struct {
	Unset CmpactionUnset `json:"cmpaction"`
}

type update_cmpaction_payload struct {
	Update update_cmpaction `json:"cmpaction"`
}

type get_cmpaction_result struct {
	Results []Cmpaction `json:"cmpaction"`
}

type count_cmpaction_result struct {
	Results []Count `json:"cmpaction"`
}

func (c *NitroClient) AddCmpaction(resource Cmpaction) error {
	payload := add_cmpaction_payload{
		resource,
	}

	return c.post("cmpaction", "", nil, payload)
}

func (c *NitroClient) RenameCmpaction(name string, newName string) error {
	payload := rename_cmpaction_payload{
		rename_cmpaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cmpaction", "", qs, payload)
}

func (c *NitroClient) CountCmpaction() (int, error) {
	var results count_cmpaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cmpaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCmpaction(key string) (bool, error) {
	var results count_cmpaction_result

	id, qs := cmpaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cmpaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCmpaction() ([]Cmpaction, error) {
	results := get_cmpaction_result{}

	if err := c.get("cmpaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCmpaction(key string) (*Cmpaction, error) {
	var results get_cmpaction_result

	id, qs := cmpaction_key_to_id_args(key)

	if err := c.get("cmpaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmpaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cmpaction element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCmpaction(key string) error {
	id, qs := cmpaction_key_to_id_args(key)

	return c.delete("cmpaction", id, qs)
}

func (c *NitroClient) UnsetCmpaction(unset CmpactionUnset) error {
	payload := unset_cmpaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("cmpaction", "", qs, payload)
}

func (c *NitroClient) UpdateCmpaction(resource Cmpaction) error {
	payload := update_cmpaction_payload{
		update_cmpaction{
			resource.Name,
			resource.Cmptype,
			resource.Addvaryheader,
			resource.Varyheadervalue,
		},
	}

	return c.put("cmpaction", "", nil, payload)
}
