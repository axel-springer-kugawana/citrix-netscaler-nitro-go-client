package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowaction struct {
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
}

type AppflowactionKey struct {
	Name string
}

func (resource Appflowaction) ToKey() AppflowactionKey {
	key := AppflowactionKey{
		resource.Name,
	}

	return key
}

func (key AppflowactionKey) to_id_args() (string, map[string]string) {
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

type add_appflowaction_payload struct {
	Resource Appflowaction `json:"appflowaction"`
}

func (c *NitroClient) AddAppflowaction(resource Appflowaction) error {
	payload := add_appflowaction_payload{
		resource,
	}

	return c.post("appflowaction", "", nil, payload)
}

//      LIST

type list_appflowaction_result struct {
	Results []Appflowaction `json:"appflowaction"`
}

func (c *NitroClient) ListAppflowaction() ([]Appflowaction, error) {
	results := list_appflowaction_result{}

	if err := c.get("appflowaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appflowaction_result struct {
	Results []Appflowaction `json:"appflowaction"`
}

func (c *NitroClient) GetAppflowaction(key AppflowactionKey) (*Appflowaction, error) {
	var results get_appflowaction_result

	id, qs := key.to_id_args()

	if err := c.get("appflowaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appflowaction_result struct {
	Results []Count `json:"appflowaction"`
}

func (c *NitroClient) CountAppflowaction() (int, error) {
	var results count_appflowaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppflowaction(key AppflowactionKey) (bool, error) {
	var results count_appflowaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appflowaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowaction(key AppflowactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowaction", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
