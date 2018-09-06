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

func (key AppflowactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AppflowactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppflowactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type AppflowactionUpdate struct {
	Name                   string   `json:"name,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Comment                string   `json:"comment,omitempty"`
}

func (resource Appflowaction) ToUpdate() AppflowactionUpdate {
	update := AppflowactionUpdate{
		resource.Name,
		resource.Collectors,
		resource.Clientsidemeasurements,
		resource.Pagetracking,
		resource.Webinsight,
		resource.Securityinsight,
		resource.Comment,
	}

	return update
}

type update_appflowaction_payload struct {
	Update AppflowactionUpdate `json:"appflowaction"`
}

func (c *NitroClient) UpdateAppflowaction(update AppflowactionUpdate) error {
	payload := update_appflowaction_payload{
		update,
	}

	return c.put("appflowaction", "", nil, payload)
}

//      UNSET

type AppflowactionUnset struct {
	Name                   string `json:"name,omitempty"`
	Collectors             bool   `json:"collectors,omitempty"`
	Clientsidemeasurements bool   `json:"clientsidemeasurements,omitempty"`
	Pagetracking           bool   `json:"pagetracking,omitempty"`
	Webinsight             bool   `json:"webinsight,omitempty"`
	Securityinsight        bool   `json:"securityinsight,omitempty"`
	Comment                bool   `json:"comment,omitempty"`
}

func (resource Appflowaction) ToUnset() AppflowactionUnset {
	unset := AppflowactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_appflowaction_payload struct {
	Unset AppflowactionUnset `json:"appflowaction"`
}

func (c *NitroClient) UnsetAppflowaction(unset AppflowactionUnset) error {
	payload := unset_appflowaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appflowaction", "", qs, payload)
}

//      RENAME
//      TODO
