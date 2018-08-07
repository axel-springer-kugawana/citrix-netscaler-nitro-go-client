package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Transformpolicy struct {
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

type TransformpolicyKey struct {
	Name string
}

func (resource Transformpolicy) ToKey() TransformpolicyKey {
	key := TransformpolicyKey{
		resource.Name,
	}

	return key
}

func (key TransformpolicyKey) to_id_args() (string, map[string]string) {
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

type add_transformpolicy_payload struct {
	Resource Transformpolicy `json:"transformpolicy"`
}

func (c *NitroClient) AddTransformpolicy(resource Transformpolicy) error {
	payload := add_transformpolicy_payload{
		resource,
	}

	return c.post("transformpolicy", "", nil, payload)
}

//      LIST

type list_transformpolicy_result struct {
	Results []Transformpolicy `json:"transformpolicy"`
}

func (c *NitroClient) ListTransformpolicy() ([]Transformpolicy, error) {
	results := list_transformpolicy_result{}

	if err := c.get("transformpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_transformpolicy_result struct {
	Results []Transformpolicy `json:"transformpolicy"`
}

func (c *NitroClient) GetTransformpolicy(key TransformpolicyKey) (*Transformpolicy, error) {
	var results get_transformpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("transformpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one transformpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("transformpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_transformpolicy_result struct {
	Results []Count `json:"transformpolicy"`
}

func (c *NitroClient) CountTransformpolicy() (int, error) {
	var results count_transformpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("transformpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsTransformpolicy(key TransformpolicyKey) (bool, error) {
	var results count_transformpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("transformpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTransformpolicy(key TransformpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("transformpolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
