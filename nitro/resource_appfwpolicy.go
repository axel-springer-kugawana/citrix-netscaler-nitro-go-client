package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appfwpolicy struct {
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}

type AppfwpolicyKey struct {
	Name string
}

func (resource Appfwpolicy) ToKey() AppfwpolicyKey {
	key := AppfwpolicyKey{
		resource.Name,
	}

	return key
}

func (key AppfwpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AppfwpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppfwpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appfwpolicy_payload struct {
	Resource Appfwpolicy `json:"appfwpolicy"`
}

func (c *NitroClient) AddAppfwpolicy(resource Appfwpolicy) error {
	payload := add_appfwpolicy_payload{
		resource,
	}

	return c.post("appfwpolicy", "", nil, payload)
}

//      LIST

type list_appfwpolicy_result struct {
	Results []Appfwpolicy `json:"appfwpolicy"`
}

func (c *NitroClient) ListAppfwpolicy() ([]Appfwpolicy, error) {
	results := list_appfwpolicy_result{}

	if err := c.get("appfwpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appfwpolicy_result struct {
	Results []Appfwpolicy `json:"appfwpolicy"`
}

func (c *NitroClient) GetAppfwpolicy(key AppfwpolicyKey) (*Appfwpolicy, error) {
	var results get_appfwpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("appfwpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appfwpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appfwpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appfwpolicy_result struct {
	Results []Count `json:"appfwpolicy"`
}

func (c *NitroClient) CountAppfwpolicy() (int, error) {
	var results count_appfwpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appfwpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppfwpolicy(key AppfwpolicyKey) (bool, error) {
	var results count_appfwpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appfwpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppfwpolicy(key AppfwpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("appfwpolicy", id, qs)
}

//      UPDATE

type AppfwpolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

func (resource Appfwpolicy) ToUpdate() AppfwpolicyUpdate {
	update := AppfwpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Profilename,
		resource.Comment,
		resource.Logaction,
	}

	return update
}

type update_appfwpolicy_payload struct {
	Update AppfwpolicyUpdate `json:"appfwpolicy"`
}

func (c *NitroClient) UpdateAppfwpolicy(update AppfwpolicyUpdate) error {
	payload := update_appfwpolicy_payload{
		update,
	}

	return c.put("appfwpolicy", "", nil, payload)
}

//      UNSET

type AppfwpolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Profilename bool   `json:"profilename,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
}

func (resource Appfwpolicy) ToUnset() AppfwpolicyUnset {
	unset := AppfwpolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_appfwpolicy_payload struct {
	Unset AppfwpolicyUnset `json:"appfwpolicy"`
}

func (c *NitroClient) UnsetAppfwpolicy(unset AppfwpolicyUnset) error {
	payload := unset_appfwpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appfwpolicy", "", qs, payload)
}

//      RENAME
//      TODO
