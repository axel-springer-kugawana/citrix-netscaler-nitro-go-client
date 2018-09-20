package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Appflowcollector struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Name       string `json:"name,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
	Port       int    `json:"port,omitempty"`
}

type AppflowcollectorKey struct {
	Name string
}

func (resource Appflowcollector) ToKey() AppflowcollectorKey {
	key := AppflowcollectorKey{
		resource.Name,
	}

	return key
}

func (key AppflowcollectorKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AppflowcollectorKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppflowcollectorKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appflowcollector_payload struct {
	Resource Appflowcollector `json:"appflowcollector"`
}

func (c *NitroClient) AddAppflowcollector(resource Appflowcollector) error {
	payload := add_appflowcollector_payload{
		resource,
	}

	return c.post("appflowcollector", "", nil, payload)
}

//      LIST

type list_appflowcollector_result struct {
	Results []Appflowcollector `json:"appflowcollector"`
}

func (c *NitroClient) ListAppflowcollector() ([]Appflowcollector, error) {
	results := list_appflowcollector_result{}

	if err := c.get("appflowcollector", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appflowcollector_result struct {
	Results []Appflowcollector `json:"appflowcollector"`
}

func (c *NitroClient) GetAppflowcollector(key AppflowcollectorKey) (*Appflowcollector, error) {
	var results get_appflowcollector_result

	id, qs := key.to_id_args()

	if err := c.get("appflowcollector", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowcollector element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowcollector element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appflowcollector_result struct {
	Results []Count `json:"appflowcollector"`
}

func (c *NitroClient) CountAppflowcollector() (int, error) {
	var results count_appflowcollector_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowcollector", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppflowcollector(key AppflowcollectorKey) (bool, error) {
	var results count_appflowcollector_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appflowcollector", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowcollector(key AppflowcollectorKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowcollector", id, qs)
}

//      UPDATE

type AppflowcollectorUpdate struct {
	Name       string `json:"name,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Port       int    `json:"port,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
}

func (resource Appflowcollector) ToUpdate() AppflowcollectorUpdate {
	update := AppflowcollectorUpdate{
		resource.Name,
		resource.Ipaddress,
		resource.Port,
		resource.Netprofile,
	}

	return update
}

type update_appflowcollector_payload struct {
	Update AppflowcollectorUpdate `json:"appflowcollector"`
}

func (c *NitroClient) UpdateAppflowcollector(update AppflowcollectorUpdate) error {
	payload := update_appflowcollector_payload{
		update,
	}

	return c.put("appflowcollector", "", nil, payload)
}

//      UNSET

type AppflowcollectorUnset struct {
	Name       string `json:"name,omitempty"`
	Ipaddress  bool   `json:"ipaddress,omitempty"`
	Port       bool   `json:"port,omitempty"`
	Netprofile bool   `json:"netprofile,omitempty"`
}

func (resource Appflowcollector) ToUnset() AppflowcollectorUnset {
	unset := AppflowcollectorUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_appflowcollector_payload struct {
	Unset AppflowcollectorUnset `json:"appflowcollector"`
}

func (c *NitroClient) UnsetAppflowcollector(unset AppflowcollectorUnset) error {
	payload := unset_appflowcollector_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appflowcollector", "", qs, payload)
}

//      RENAME
//      TODO
