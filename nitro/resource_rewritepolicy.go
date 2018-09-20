package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Rewritepolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type RewritepolicyKey struct {
	Name string
}

func (resource Rewritepolicy) ToKey() RewritepolicyKey {
	key := RewritepolicyKey{
		resource.Name,
	}

	return key
}

func (key RewritepolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key RewritepolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key RewritepolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_rewritepolicy_payload struct {
	Resource Rewritepolicy `json:"rewritepolicy"`
}

func (c *NitroClient) AddRewritepolicy(resource Rewritepolicy) error {
	payload := add_rewritepolicy_payload{
		resource,
	}

	return c.post("rewritepolicy", "", nil, payload)
}

//      LIST

type list_rewritepolicy_result struct {
	Results []Rewritepolicy `json:"rewritepolicy"`
}

func (c *NitroClient) ListRewritepolicy() ([]Rewritepolicy, error) {
	results := list_rewritepolicy_result{}

	if err := c.get("rewritepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_rewritepolicy_result struct {
	Results []Rewritepolicy `json:"rewritepolicy"`
}

func (c *NitroClient) GetRewritepolicy(key RewritepolicyKey) (*Rewritepolicy, error) {
	var results get_rewritepolicy_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_rewritepolicy_result struct {
	Results []Count `json:"rewritepolicy"`
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

//      EXISTS

func (c *NitroClient) ExistsRewritepolicy(key RewritepolicyKey) (bool, error) {
	var results count_rewritepolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("rewritepolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteRewritepolicy(key RewritepolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("rewritepolicy", id, qs)
}

//      UPDATE

type RewritepolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

func (resource Rewritepolicy) ToUpdate() RewritepolicyUpdate {
	update := RewritepolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
		resource.Logaction,
	}

	return update
}

type update_rewritepolicy_payload struct {
	Update RewritepolicyUpdate `json:"rewritepolicy"`
}

func (c *NitroClient) UpdateRewritepolicy(update RewritepolicyUpdate) error {
	payload := update_rewritepolicy_payload{
		update,
	}

	return c.put("rewritepolicy", "", nil, payload)
}

//      UNSET

type RewritepolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
}

func (resource Rewritepolicy) ToUnset() RewritepolicyUnset {
	unset := RewritepolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_rewritepolicy_payload struct {
	Unset RewritepolicyUnset `json:"rewritepolicy"`
}

func (c *NitroClient) UnsetRewritepolicy(unset RewritepolicyUnset) error {
	payload := unset_rewritepolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("rewritepolicy", "", qs, payload)
}

//      RENAME
//      TODO
