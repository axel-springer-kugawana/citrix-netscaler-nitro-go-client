package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowpolicylabel struct {
	Labelname       string `json:"labelname,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
}

type AppflowpolicylabelKey struct {
	Labelname string
}

func (resource Appflowpolicylabel) ToKey() AppflowpolicylabelKey {
	key := AppflowpolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key AppflowpolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AppflowpolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppflowpolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appflowpolicylabel_payload struct {
	Resource Appflowpolicylabel `json:"appflowpolicylabel"`
}

func (c *NitroClient) AddAppflowpolicylabel(resource Appflowpolicylabel) error {
	payload := add_appflowpolicylabel_payload{
		resource,
	}

	return c.post("appflowpolicylabel", "", nil, payload)
}

//      LIST

type list_appflowpolicylabel_result struct {
	Results []Appflowpolicylabel `json:"appflowpolicylabel"`
}

func (c *NitroClient) ListAppflowpolicylabel() ([]Appflowpolicylabel, error) {
	results := list_appflowpolicylabel_result{}

	if err := c.get("appflowpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appflowpolicylabel_result struct {
	Results []Appflowpolicylabel `json:"appflowpolicylabel"`
}

func (c *NitroClient) GetAppflowpolicylabel(key AppflowpolicylabelKey) (*Appflowpolicylabel, error) {
	var results get_appflowpolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("appflowpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appflowpolicylabel_result struct {
	Results []Count `json:"appflowpolicylabel"`
}

func (c *NitroClient) CountAppflowpolicylabel() (int, error) {
	var results count_appflowpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppflowpolicylabel(key AppflowpolicylabelKey) (bool, error) {
	var results count_appflowpolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appflowpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowpolicylabel(key AppflowpolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowpolicylabel", id, qs)
}

//      RENAME
//      TODO
