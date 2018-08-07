package nitro

import (
	"fmt"
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

	var id string
	var args []string

	id = key.Name

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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
