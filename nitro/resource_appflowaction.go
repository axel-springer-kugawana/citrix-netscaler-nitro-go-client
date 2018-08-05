package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appflowaction struct {
	Name                   string   `json:"name"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
	Metricslog             bool     `json:"metricslog,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Transactionlog         string   `json:"transactionlog,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
}

func appflowaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppflowactionUnset struct {
	Name                   string `json:"name"`
	Collectors             bool   `json:"collectors,omitempty"`
	Clientsidemeasurements bool   `json:"clientsidemeasurements,omitempty"`
	Comment                bool   `json:"comment,omitempty"`
	Pagetracking           bool   `json:"pagetracking,omitempty"`
	Webinsight             bool   `json:"webinsight,omitempty"`
	Securityinsight        bool   `json:"securityinsight,omitempty"`
	Videoanalytics         bool   `json:"videoanalytics,omitempty"`
	Distributionalgorithm  bool   `json:"distributionalgorithm,omitempty"`
}

type update_appflowaction struct {
	Name                   string   `json:"name"`
	Collectors             []string `json:"collectors,omitempty"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
}

type rename_appflowaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appflowaction_payload struct {
	Resource Appflowaction `json:"appflowaction"`
}

type rename_appflowaction_payload struct {
	Rename rename_appflowaction `json:"appflowaction"`
}

type unset_appflowaction_payload struct {
	Unset AppflowactionUnset `json:"appflowaction"`
}

type update_appflowaction_payload struct {
	Update update_appflowaction `json:"appflowaction"`
}

type get_appflowaction_result struct {
	Results []Appflowaction `json:"appflowaction"`
}

type count_appflowaction_result struct {
	Results []Count `json:"appflowaction"`
}

func (c *NitroClient) AddAppflowaction(resource Appflowaction) error {
	payload := add_appflowaction_payload{
		resource,
	}

	return c.post("appflowaction", "", nil, payload)
}

func (c *NitroClient) RenameAppflowaction(name string, newName string) error {
	payload := rename_appflowaction_payload{
		rename_appflowaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appflowaction", "", qs, payload)
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

func (c *NitroClient) ExistsAppflowaction(key string) (bool, error) {
	var results count_appflowaction_result

	id, qs := appflowaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appflowaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppflowaction() ([]Appflowaction, error) {
	results := get_appflowaction_result{}

	if err := c.get("appflowaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowaction(key string) (*Appflowaction, error) {
	var results get_appflowaction_result

	id, qs := appflowaction_key_to_id_args(key)

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

func (c *NitroClient) DeleteAppflowaction(key string) error {
	id, qs := appflowaction_key_to_id_args(key)

	return c.delete("appflowaction", id, qs)
}

func (c *NitroClient) UnsetAppflowaction(unset AppflowactionUnset) error {
	payload := unset_appflowaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appflowaction", "", qs, payload)
}

func (c *NitroClient) UpdateAppflowaction(resource Appflowaction) error {
	payload := update_appflowaction_payload{
		update_appflowaction{
			resource.Name,
			resource.Collectors,
			resource.Clientsidemeasurements,
			resource.Comment,
			resource.Pagetracking,
			resource.Webinsight,
			resource.Securityinsight,
			resource.Videoanalytics,
			resource.Distributionalgorithm,
		},
	}

	return c.put("appflowaction", "", nil, payload)
}
