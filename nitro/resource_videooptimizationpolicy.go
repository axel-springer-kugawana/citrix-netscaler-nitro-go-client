package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationpolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type VideooptimizationpolicyKey struct {
	Name string
}

func (resource Videooptimizationpolicy) ToKey() VideooptimizationpolicyKey {
	key := VideooptimizationpolicyKey{
		resource.Name,
	}

	return key
}

func (key VideooptimizationpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key VideooptimizationpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key VideooptimizationpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_videooptimizationpolicy_payload struct {
	Resource Videooptimizationpolicy `json:"videooptimizationpolicy"`
}

func (c *NitroClient) AddVideooptimizationpolicy(resource Videooptimizationpolicy) error {
	payload := add_videooptimizationpolicy_payload{
		resource,
	}

	return c.post("videooptimizationpolicy", "", nil, payload)
}

//      LIST

type list_videooptimizationpolicy_result struct {
	Results []Videooptimizationpolicy `json:"videooptimizationpolicy"`
}

func (c *NitroClient) ListVideooptimizationpolicy() ([]Videooptimizationpolicy, error) {
	results := list_videooptimizationpolicy_result{}

	if err := c.get("videooptimizationpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_videooptimizationpolicy_result struct {
	Results []Videooptimizationpolicy `json:"videooptimizationpolicy"`
}

func (c *NitroClient) GetVideooptimizationpolicy(key VideooptimizationpolicyKey) (*Videooptimizationpolicy, error) {
	var results get_videooptimizationpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("videooptimizationpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one videooptimizationpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("videooptimizationpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_videooptimizationpolicy_result struct {
	Results []Count `json:"videooptimizationpolicy"`
}

func (c *NitroClient) CountVideooptimizationpolicy() (int, error) {
	var results count_videooptimizationpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("videooptimizationpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsVideooptimizationpolicy(key VideooptimizationpolicyKey) (bool, error) {
	var results count_videooptimizationpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("videooptimizationpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteVideooptimizationpolicy(key VideooptimizationpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("videooptimizationpolicy", id, qs)
}

//      UPDATE

type VideooptimizationpolicyUpdate struct {
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

func (resource Videooptimizationpolicy) ToUpdate() VideooptimizationpolicyUpdate {
	update := VideooptimizationpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
		resource.Undefaction,
		resource.Comment,
		resource.Logaction,
	}

	return update
}

type update_videooptimizationpolicy_payload struct {
	Update VideooptimizationpolicyUpdate `json:"videooptimizationpolicy"`
}

func (c *NitroClient) UpdateVideooptimizationpolicy(update VideooptimizationpolicyUpdate) error {
	payload := update_videooptimizationpolicy_payload{
		update,
	}

	return c.put("videooptimizationpolicy", "", nil, payload)
}

//      UNSET

type VideooptimizationpolicyUnset struct {
	Name        string `json:"name,omitempty"`
	Rule        bool   `json:"rule,omitempty"`
	Action      bool   `json:"action,omitempty"`
	Undefaction bool   `json:"undefaction,omitempty"`
	Comment     bool   `json:"comment,omitempty"`
	Logaction   bool   `json:"logaction,omitempty"`
}

func (resource Videooptimizationpolicy) ToUnset() VideooptimizationpolicyUnset {
	unset := VideooptimizationpolicyUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_videooptimizationpolicy_payload struct {
	Unset VideooptimizationpolicyUnset `json:"videooptimizationpolicy"`
}

func (c *NitroClient) UnsetVideooptimizationpolicy(unset VideooptimizationpolicyUnset) error {
	payload := unset_videooptimizationpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("videooptimizationpolicy", "", qs, payload)
}

//      RENAME
//      TODO
