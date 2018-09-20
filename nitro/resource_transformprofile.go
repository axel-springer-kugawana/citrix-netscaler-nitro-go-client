package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Transformprofile struct {
	Comment                   string `json:"comment,omitempty"`
	Name                      string `json:"name,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Type                      string `json:"type,omitempty"`
}

type TransformprofileKey struct {
	Name string
}

func (resource Transformprofile) ToKey() TransformprofileKey {
	key := TransformprofileKey{
		resource.Name,
	}

	return key
}

func (key TransformprofileKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key TransformprofileKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TransformprofileKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_transformprofile_payload struct {
	Resource Transformprofile `json:"transformprofile"`
}

func (c *NitroClient) AddTransformprofile(resource Transformprofile) error {
	payload := add_transformprofile_payload{
		resource,
	}

	return c.post("transformprofile", "", nil, payload)
}

//      LIST

type list_transformprofile_result struct {
	Results []Transformprofile `json:"transformprofile"`
}

func (c *NitroClient) ListTransformprofile() ([]Transformprofile, error) {
	results := list_transformprofile_result{}

	if err := c.get("transformprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_transformprofile_result struct {
	Results []Transformprofile `json:"transformprofile"`
}

func (c *NitroClient) GetTransformprofile(key TransformprofileKey) (*Transformprofile, error) {
	var results get_transformprofile_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_transformprofile_result struct {
	Results []Count `json:"transformprofile"`
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

//      EXISTS

func (c *NitroClient) ExistsTransformprofile(key TransformprofileKey) (bool, error) {
	var results count_transformprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("transformprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTransformprofile(key TransformprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("transformprofile", id, qs)
}

//      UPDATE

type TransformprofileUpdate struct {
	Name                      string `json:"name,omitempty"`
	Type                      string `json:"type,omitempty"`
	Onlytransformabsurlinbody string `json:"onlytransformabsurlinbody,omitempty"`
	Comment                   string `json:"comment,omitempty"`
}

func (resource Transformprofile) ToUpdate() TransformprofileUpdate {
	update := TransformprofileUpdate{
		resource.Name,
		resource.Type,
		resource.Onlytransformabsurlinbody,
		resource.Comment,
	}

	return update
}

type update_transformprofile_payload struct {
	Update TransformprofileUpdate `json:"transformprofile"`
}

func (c *NitroClient) UpdateTransformprofile(update TransformprofileUpdate) error {
	payload := update_transformprofile_payload{
		update,
	}

	return c.put("transformprofile", "", nil, payload)
}

//      UNSET

type TransformprofileUnset struct {
	Name                      string `json:"name,omitempty"`
	Type                      bool   `json:"type,omitempty"`
	Onlytransformabsurlinbody bool   `json:"onlytransformabsurlinbody,omitempty"`
	Comment                   bool   `json:"comment,omitempty"`
}

func (resource Transformprofile) ToUnset() TransformprofileUnset {
	unset := TransformprofileUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_transformprofile_payload struct {
	Unset TransformprofileUnset `json:"transformprofile"`
}

func (c *NitroClient) UnsetTransformprofile(unset TransformprofileUnset) error {
	payload := unset_transformprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("transformprofile", "", qs, payload)
}

//      RENAME
//      TODO
