package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Transformaction struct {
	Comment          string `json:"comment,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Name             string `json:"name,omitempty"`
	Priority         int    `json:"priority,string,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	State            string `json:"state,omitempty"`
}

type TransformactionKey struct {
	Name string
}

func (resource Transformaction) ToKey() TransformactionKey {
	key := TransformactionKey{
		resource.Name,
	}

	return key
}

func (key TransformactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TransformactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TransformactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_transformaction_payload struct {
	Resource Transformaction `json:"transformaction"`
}

func (c *NitroClient) AddTransformaction(resource Transformaction) error {
	payload := add_transformaction_payload{
		resource,
	}

	return c.post("transformaction", "", nil, payload)
}

//      LIST

type list_transformaction_result struct {
	Results []Transformaction `json:"transformaction"`
}

func (c *NitroClient) ListTransformaction() ([]Transformaction, error) {
	results := list_transformaction_result{}

	if err := c.get("transformaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_transformaction_result struct {
	Results []Transformaction `json:"transformaction"`
}

func (c *NitroClient) GetTransformaction(key TransformactionKey) (*Transformaction, error) {
	var results get_transformaction_result

	id, qs := key.to_id_args()

	if err := c.get("transformaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("transformaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_transformaction_result struct {
	Results []Count `json:"transformaction"`
}

func (c *NitroClient) CountTransformaction() (int, error) {
	var results count_transformaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTransformaction(key TransformactionKey) (bool, error) {
	var results count_transformaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("transformaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTransformaction(key TransformactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("transformaction", id, qs)
}

//      UPDATE

type TransformactionUpdate struct {
	Name             string `json:"name,omitempty"`
	Priority         int    `json:"priority,string,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	State            string `json:"state,omitempty"`
	Comment          string `json:"comment,omitempty"`
}

func (resource Transformaction) ToUpdate() TransformactionUpdate {
	update := TransformactionUpdate{
		resource.Name,
		resource.Priority,
		resource.Requrlfrom,
		resource.Requrlinto,
		resource.Resurlfrom,
		resource.Resurlinto,
		resource.Cookiedomainfrom,
		resource.Cookiedomaininto,
		resource.State,
		resource.Comment,
	}

	return update
}

type update_transformaction_payload struct {
	Update TransformactionUpdate `json:"transformaction"`
}

func (c *NitroClient) UpdateTransformaction(update TransformactionUpdate) error {
	payload := update_transformaction_payload{
		update,
	}

	return c.put("transformaction", "", nil, payload)
}

//      UNSET

type TransformactionUnset struct {
	Name             string `json:"name,omitempty"`
	Priority         bool   `json:"priority,omitempty"`
	Requrlfrom       bool   `json:"requrlfrom,omitempty"`
	Requrlinto       bool   `json:"requrlinto,omitempty"`
	Resurlfrom       bool   `json:"resurlfrom,omitempty"`
	Resurlinto       bool   `json:"resurlinto,omitempty"`
	Cookiedomainfrom bool   `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto bool   `json:"cookiedomaininto,omitempty"`
	State            bool   `json:"state,omitempty"`
	Comment          bool   `json:"comment,omitempty"`
}

func (resource Transformaction) ToUnset() TransformactionUnset {
	unset := TransformactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_transformaction_payload struct {
	Unset TransformactionUnset `json:"transformaction"`
}

func (c *NitroClient) UnsetTransformaction(unset TransformactionUnset) error {
	payload := unset_transformaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("transformaction", "", qs, payload)
}

//      RENAME
//      TODO
