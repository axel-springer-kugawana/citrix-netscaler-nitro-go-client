package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformprofile struct {
	Name                      string `json:"name"`
	Comment                   string `json:"comment,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Type                      string `json:"type,omitempty"`
}

func transformprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type TransformprofileUnset struct {
	Name                      string `json:"name"`
	Type                      bool   `json:"type,omitempty"`
	Onlytransformabsurlinbody bool   `json:"onlytransformabsurlinbody,omitempty"`
	Comment                   bool   `json:"comment,omitempty"`
}

type update_transformprofile struct {
	Name                      string `json:"name"`
	Type                      string `json:"type,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Comment                   string `json:"comment,omitempty"`
}

type rename_transformprofile struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_transformprofile_payload struct {
	Resource Transformprofile `json:"transformprofile"`
}

type rename_transformprofile_payload struct {
	Rename rename_transformprofile `json:"transformprofile"`
}

type unset_transformprofile_payload struct {
	Unset TransformprofileUnset `json:"transformprofile"`
}

type update_transformprofile_payload struct {
	Update update_transformprofile `json:"transformprofile"`
}

type get_transformprofile_result struct {
	Results []Transformprofile `json:"transformprofile"`
}

type count_transformprofile_result struct {
	Results []Count `json:"transformprofile"`
}

func (c *NitroClient) AddTransformprofile(resource Transformprofile) error {
	payload := add_transformprofile_payload{
		resource,
	}

	return c.post("transformprofile", "", nil, payload)
}

func (c *NitroClient) RenameTransformprofile(name string, newName string) error {
	payload := rename_transformprofile_payload{
		rename_transformprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("transformprofile", "", qs, payload)
}

func (c *NitroClient) CountTransformprofile() (int, error) {
	var results count_transformprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTransformprofile(key string) (bool, error) {
	var results count_transformprofile_result

	id, qs := transformprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("transformprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTransformprofile() ([]Transformprofile, error) {
	results := get_transformprofile_result{}

	if err := c.get("transformprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTransformprofile(key string) (*Transformprofile, error) {
	var results get_transformprofile_result

	id, qs := transformprofile_key_to_id_args(key)

	if err := c.get("transformprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("transformprofile element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTransformprofile(key string) error {
	id, qs := transformprofile_key_to_id_args(key)

	return c.delete("transformprofile", id, qs)
}

func (c *NitroClient) UnsetTransformprofile(unset TransformprofileUnset) error {
	payload := unset_transformprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("transformprofile", "", qs, payload)
}

func (c *NitroClient) UpdateTransformprofile(resource Transformprofile) error {
	payload := update_transformprofile_payload{
		update_transformprofile{
			resource.Name,
			resource.Type,
			resource.Onlytransformabsurlinbody,
			resource.Comment,
		},
	}

	return c.put("transformprofile", "", nil, payload)
}
