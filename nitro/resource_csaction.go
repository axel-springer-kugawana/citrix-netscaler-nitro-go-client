package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Csaction struct {
	Name              string `json:"name"`
	Comment           string `json:"comment,omitempty"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
}

func csaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type CsactionUnset struct {
	Name              string `json:"name"`
	Targetlbvserver   bool   `json:"targetlbvserver,string,omitempty"`
	Targetvserver     bool   `json:"targetvserver,string,omitempty"`
	Targetvserverexpr bool   `json:"targetvserverexpr,string,omitempty"`
	Comment           bool   `json:"comment,string,omitempty"`
}

type update_csaction struct {
	Name              string `json:"name"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

type rename_csaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_csaction_payload struct {
	Resource Csaction `json:"csaction"`
}

type rename_csaction_payload struct {
	Rename rename_csaction `json:"csaction"`
}

type unset_csaction_payload struct {
	Unset CsactionUnset `json:"csaction"`
}

type update_csaction_payload struct {
	Update update_csaction `json:"csaction"`
}

type get_csaction_result struct {
	Results []Csaction `json:"csaction"`
}

type count_csaction_result struct {
	Results []Count `json:"csaction"`
}

func (c *NitroClient) AddCsaction(resource Csaction) error {
	payload := add_csaction_payload{
		resource,
	}

	return c.post("csaction", "", nil, payload)
}

func (c *NitroClient) RenameCsaction(name string, newName string) error {
	payload := rename_csaction_payload{
		rename_csaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("csaction", "", qs, payload)
}

func (c *NitroClient) CountCsaction() (int, error) {
	var results count_csaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("csaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCsaction(key string) (bool, error) {
	var results count_csaction_result

	id, qs := csaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("csaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCsaction() ([]Csaction, error) {
	results := get_csaction_result{}

	if err := c.get("csaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCsaction(key string) (*Csaction, error) {
	var results get_csaction_result

	id, qs := csaction_key_to_id_args(key)

	if err := c.get("csaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csaction element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCsaction(key string) error {
	id, qs := csaction_key_to_id_args(key)

	return c.delete("csaction", id, qs)
}

func (c *NitroClient) UnsetCsaction(unset CsactionUnset) error {
	payload := unset_csaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("csaction", "", qs, payload)
}

func (c *NitroClient) UpdateCsaction(resource Csaction) error {
	payload := update_csaction_payload{
		update_csaction{
			resource.Name,
			resource.Targetlbvserver,
			resource.Targetvserver,
			resource.Targetvserverexpr,
			resource.Comment,
		},
	}

	return c.put("csaction", "", nil, payload)
}
