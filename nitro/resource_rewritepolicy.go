package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Rewritepolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

func rewritepolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type RewritepolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
}

type update_rewritepolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type rename_rewritepolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_rewritepolicy_payload struct {
	Resource Rewritepolicy `json:"rewritepolicy"`
}

type rename_rewritepolicy_payload struct {
	Rename rename_rewritepolicy `json:"rewritepolicy"`
}

type unset_rewritepolicy_payload struct {
	Unset RewritepolicyUnset `json:"rewritepolicy"`
}

type update_rewritepolicy_payload struct {
	Update update_rewritepolicy `json:"rewritepolicy"`
}

type get_rewritepolicy_result struct {
	Results []Rewritepolicy `json:"rewritepolicy"`
}

type count_rewritepolicy_result struct {
	Results []Count `json:"rewritepolicy"`
}

func (c *NitroClient) AddRewritepolicy(resource Rewritepolicy) error {
	payload := add_rewritepolicy_payload{
		resource,
	}

	return c.post("rewritepolicy", "", nil, payload)
}

func (c *NitroClient) RenameRewritepolicy(name string, newName string) error {
	payload := rename_rewritepolicy_payload{
		rename_rewritepolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("rewritepolicy", "", qs, payload)
}

func (c *NitroClient) CountRewritepolicy() (int, error) {
	var results count_rewritepolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("rewritepolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsRewritepolicy(key string) (bool, error) {
	var results count_rewritepolicy_result

	id, qs := rewritepolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("rewritepolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListRewritepolicy() ([]Rewritepolicy, error) {
	results := get_rewritepolicy_result{}

	if err := c.get("rewritepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetRewritepolicy(key string) (*Rewritepolicy, error) {
	var results get_rewritepolicy_result

	id, qs := rewritepolicy_key_to_id_args(key)

	if err := c.get("rewritepolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one rewritepolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("rewritepolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteRewritepolicy(key string) error {
	id, qs := rewritepolicy_key_to_id_args(key)

	return c.delete("rewritepolicy", id, qs)
}

func (c *NitroClient) UnsetRewritepolicy(unset RewritepolicyUnset) error {
	payload := unset_rewritepolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("rewritepolicy", "", qs, payload)
}

func (c *NitroClient) UpdateRewritepolicy(resource Rewritepolicy) error {
	payload := update_rewritepolicy_payload{
		update_rewritepolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Undefaction,
			resource.Comment,
			resource.Logaction,
		},
	}

	return c.put("rewritepolicy", "", nil, payload)
}
