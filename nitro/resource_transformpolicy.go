package nitro

import (
	"fmt"
	"net/url"
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

func (key TransformpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key TransformpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TransformpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type TransformpolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

func (resource Transformpolicy) ToUpdate() TransformpolicyUpdate {
	update := TransformpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Profilename,
		resource.Comment,
		resource.Logaction,
	}

	return update
}

type update_transformpolicy_payload struct {
	Update TransformpolicyUpdate `json:"transformpolicy"`
}

func (c *NitroClient) UpdateTransformpolicy(update TransformpolicyUpdate) error {
	payload := update_transformpolicy_payload{
		update,
	}

	return c.put("transformpolicy", "", nil, payload)
}

//      UNSET

type TransformpolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Profilename bool   `json:"profilename,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
}

func (resource Transformpolicy) ToUnset() TransformpolicyUnset {
	unset := TransformpolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_transformpolicy_payload struct {
	Unset TransformpolicyUnset `json:"transformpolicy"`
}

func (c *NitroClient) UnsetTransformpolicy(unset TransformpolicyUnset) error {
	payload := unset_transformpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("transformpolicy", "", qs, payload)
}

//      RENAME
//      TODO
