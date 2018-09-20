package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Capolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type CapolicyKey struct {
	Name string
}

func (resource Capolicy) ToKey() CapolicyKey {
	key := CapolicyKey{
		resource.Name,
	}

	return key
}

func (key CapolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CapolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CapolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_capolicy_payload struct {
	Resource Capolicy `json:"capolicy"`
}

func (c *NitroClient) AddCapolicy(resource Capolicy) error {
	payload := add_capolicy_payload{
		resource,
	}

	return c.post("capolicy", "", nil, payload)
}

//      LIST

type list_capolicy_result struct {
	Results []Capolicy `json:"capolicy"`
}

func (c *NitroClient) ListCapolicy() ([]Capolicy, error) {
	results := list_capolicy_result{}

	if err := c.get("capolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_capolicy_result struct {
	Results []Capolicy `json:"capolicy"`
}

func (c *NitroClient) GetCapolicy(key CapolicyKey) (*Capolicy, error) {
	var results get_capolicy_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_capolicy_result struct {
	Results []Count `json:"capolicy"`
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

//      EXISTS

func (c *NitroClient) ExistsCapolicy(key CapolicyKey) (bool, error) {
	var results count_capolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("capolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCapolicy(key CapolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("capolicy", id, qs)
}

//      UPDATE

type CapolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

func (resource Capolicy) ToUpdate() CapolicyUpdate {
	update := CapolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Comment,
		resource.Logaction,
		resource.Undefaction,
	}

	return update
}

type update_capolicy_payload struct {
	Update CapolicyUpdate `json:"capolicy"`
}

func (c *NitroClient) UpdateCapolicy(update CapolicyUpdate) error {
	payload := update_capolicy_payload{
		update,
	}

	return c.put("capolicy", "", nil, payload)
}

//      UNSET

type CapolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
}

func (resource Capolicy) ToUnset() CapolicyUnset {
	unset := CapolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_capolicy_payload struct {
	Unset CapolicyUnset `json:"capolicy"`
}

func (c *NitroClient) UnsetCapolicy(unset CapolicyUnset) error {
	payload := unset_capolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("capolicy", "", qs, payload)
}

//      RENAME
//      TODO
