package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Appflowpolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type AppflowpolicyKey struct {
	Name string
}

func (resource Appflowpolicy) ToKey() AppflowpolicyKey {
	key := AppflowpolicyKey{
		resource.Name,
	}

	return key
}

func (key AppflowpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AppflowpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppflowpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appflowpolicy_payload struct {
	Resource Appflowpolicy `json:"appflowpolicy"`
}

func (c *NitroClient) AddAppflowpolicy(resource Appflowpolicy) error {
	payload := add_appflowpolicy_payload{
		resource,
	}

	return c.post("appflowpolicy", "", nil, payload)
}

//      LIST

type list_appflowpolicy_result struct {
	Results []Appflowpolicy `json:"appflowpolicy"`
}

func (c *NitroClient) ListAppflowpolicy() ([]Appflowpolicy, error) {
	results := list_appflowpolicy_result{}

	if err := c.get("appflowpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appflowpolicy_result struct {
	Results []Appflowpolicy `json:"appflowpolicy"`
}

func (c *NitroClient) GetAppflowpolicy(key AppflowpolicyKey) (*Appflowpolicy, error) {
	var results get_appflowpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("appflowpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appflowpolicy_result struct {
	Results []Count `json:"appflowpolicy"`
}

func (c *NitroClient) CountAppflowpolicy() (int, error) {
	var results count_appflowpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppflowpolicy(key AppflowpolicyKey) (bool, error) {
	var results count_appflowpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appflowpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowpolicy(key AppflowpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowpolicy", id, qs)
}

//      UPDATE

type AppflowpolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

func (resource Appflowpolicy) ToUpdate() AppflowpolicyUpdate {
	update := AppflowpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
	}

	return update
}

type update_appflowpolicy_payload struct {
	Update AppflowpolicyUpdate `json:"appflowpolicy"`
}

func (c *NitroClient) UpdateAppflowpolicy(update AppflowpolicyUpdate) error {
	payload := update_appflowpolicy_payload{
		update,
	}

	return c.put("appflowpolicy", "", nil, payload)
}

//      UNSET

type AppflowpolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
}

func (resource Appflowpolicy) ToUnset() AppflowpolicyUnset {
	unset := AppflowpolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_appflowpolicy_payload struct {
	Unset AppflowpolicyUnset `json:"appflowpolicy"`
}

func (c *NitroClient) UnsetAppflowpolicy(unset AppflowpolicyUnset) error {
	payload := unset_appflowpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appflowpolicy", "", qs, payload)
}

//      RENAME
//      TODO
