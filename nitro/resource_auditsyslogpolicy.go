package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditsyslogpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func auditsyslogpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AuditsyslogpolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_auditsyslogpolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_auditsyslogpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_auditsyslogpolicy_payload struct {
	Resource Auditsyslogpolicy `json:"auditsyslogpolicy"`
}

type rename_auditsyslogpolicy_payload struct {
	Rename rename_auditsyslogpolicy `json:"auditsyslogpolicy"`
}

type unset_auditsyslogpolicy_payload struct {
	Unset AuditsyslogpolicyUnset `json:"auditsyslogpolicy"`
}

type update_auditsyslogpolicy_payload struct {
	Update update_auditsyslogpolicy `json:"auditsyslogpolicy"`
}

type get_auditsyslogpolicy_result struct {
	Results []Auditsyslogpolicy `json:"auditsyslogpolicy"`
}

type count_auditsyslogpolicy_result struct {
	Results []Count `json:"auditsyslogpolicy"`
}

func (c *NitroClient) AddAuditsyslogpolicy(resource Auditsyslogpolicy) error {
	payload := add_auditsyslogpolicy_payload{
		resource,
	}

	return c.post("auditsyslogpolicy", "", nil, payload)
}

func (c *NitroClient) RenameAuditsyslogpolicy(name string, newName string) error {
	payload := rename_auditsyslogpolicy_payload{
		rename_auditsyslogpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("auditsyslogpolicy", "", qs, payload)
}

func (c *NitroClient) CountAuditsyslogpolicy() (int, error) {
	var results count_auditsyslogpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("auditsyslogpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuditsyslogpolicy(key string) (bool, error) {
	var results count_auditsyslogpolicy_result

	id, qs := auditsyslogpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("auditsyslogpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuditsyslogpolicy() ([]Auditsyslogpolicy, error) {
	results := get_auditsyslogpolicy_result{}

	if err := c.get("auditsyslogpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuditsyslogpolicy(key string) (*Auditsyslogpolicy, error) {
	var results get_auditsyslogpolicy_result

	id, qs := auditsyslogpolicy_key_to_id_args(key)

	if err := c.get("auditsyslogpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one auditsyslogpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("auditsyslogpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuditsyslogpolicy(key string) error {
	id, qs := auditsyslogpolicy_key_to_id_args(key)

	return c.delete("auditsyslogpolicy", id, qs)
}

func (c *NitroClient) UnsetAuditsyslogpolicy(unset AuditsyslogpolicyUnset) error {
	payload := unset_auditsyslogpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("auditsyslogpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAuditsyslogpolicy(resource Auditsyslogpolicy) error {
	payload := update_auditsyslogpolicy_payload{
		update_auditsyslogpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("auditsyslogpolicy", "", nil, payload)
}
