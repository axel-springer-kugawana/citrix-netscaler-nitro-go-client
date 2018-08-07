package nitro

import (
	"fmt"
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

func (key TransformactionKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
