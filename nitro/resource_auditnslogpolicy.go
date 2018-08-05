package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditnslogpolicy struct {
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

func auditnslogpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AuditnslogpolicyUnset struct {
	Name   string `json:"name"`
	Rule   bool   `json:"rule,string,omitempty"`
	Action bool   `json:"action,string,omitempty"`
}

type update_auditnslogpolicy struct {
	Name   string `json:"name"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

type rename_auditnslogpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_auditnslogpolicy_payload struct {
	Resource Auditnslogpolicy `json:"auditnslogpolicy"`
}

type rename_auditnslogpolicy_payload struct {
	Rename rename_auditnslogpolicy `json:"auditnslogpolicy"`
}

type unset_auditnslogpolicy_payload struct {
	Unset AuditnslogpolicyUnset `json:"auditnslogpolicy"`
}

type update_auditnslogpolicy_payload struct {
	Update update_auditnslogpolicy `json:"auditnslogpolicy"`
}

type get_auditnslogpolicy_result struct {
	Results []Auditnslogpolicy `json:"auditnslogpolicy"`
}

type count_auditnslogpolicy_result struct {
	Results []Count `json:"auditnslogpolicy"`
}

func (c *NitroClient) AddAuditnslogpolicy(resource Auditnslogpolicy) error {
	payload := add_auditnslogpolicy_payload{
		resource,
	}

	return c.post("auditnslogpolicy", "", nil, payload)
}

func (c *NitroClient) RenameAuditnslogpolicy(name string, newName string) error {
	payload := rename_auditnslogpolicy_payload{
		rename_auditnslogpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("auditnslogpolicy", "", qs, payload)
}

func (c *NitroClient) CountAuditnslogpolicy() (int, error) {
	var results count_auditnslogpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("auditnslogpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuditnslogpolicy(key string) (bool, error) {
	var results count_auditnslogpolicy_result

	id, qs := auditnslogpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("auditnslogpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuditnslogpolicy() ([]Auditnslogpolicy, error) {
	results := get_auditnslogpolicy_result{}

	if err := c.get("auditnslogpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuditnslogpolicy(key string) (*Auditnslogpolicy, error) {
	var results get_auditnslogpolicy_result

	id, qs := auditnslogpolicy_key_to_id_args(key)

	if err := c.get("auditnslogpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one auditnslogpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("auditnslogpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuditnslogpolicy(key string) error {
	id, qs := auditnslogpolicy_key_to_id_args(key)

	return c.delete("auditnslogpolicy", id, qs)
}

func (c *NitroClient) UnsetAuditnslogpolicy(unset AuditnslogpolicyUnset) error {
	payload := unset_auditnslogpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("auditnslogpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateAuditnslogpolicy(resource Auditnslogpolicy) error {
	payload := update_auditnslogpolicy_payload{
		update_auditnslogpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
		},
	}

	return c.put("auditnslogpolicy", "", nil, payload)
}
