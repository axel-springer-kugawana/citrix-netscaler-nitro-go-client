package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Auditnslogpolicy struct {
	Action string `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type AuditnslogpolicyKey struct {
	Name string
}

func (resource Auditnslogpolicy) ToKey() AuditnslogpolicyKey {
	key := AuditnslogpolicyKey{
		resource.Name,
	}

	return key
}

func (key AuditnslogpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AuditnslogpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuditnslogpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_auditnslogpolicy_payload struct {
	Resource Auditnslogpolicy `json:"auditnslogpolicy"`
}

func (c *NitroClient) AddAuditnslogpolicy(resource Auditnslogpolicy) error {
	payload := add_auditnslogpolicy_payload{
		resource,
	}

	return c.post("auditnslogpolicy", "", nil, payload)
}

//      LIST

type list_auditnslogpolicy_result struct {
	Results []Auditnslogpolicy `json:"auditnslogpolicy"`
}

func (c *NitroClient) ListAuditnslogpolicy() ([]Auditnslogpolicy, error) {
	results := list_auditnslogpolicy_result{}

	if err := c.get("auditnslogpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_auditnslogpolicy_result struct {
	Results []Auditnslogpolicy `json:"auditnslogpolicy"`
}

func (c *NitroClient) GetAuditnslogpolicy(key AuditnslogpolicyKey) (*Auditnslogpolicy, error) {
	var results get_auditnslogpolicy_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_auditnslogpolicy_result struct {
	Results []Count `json:"auditnslogpolicy"`
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

//      EXISTS

func (c *NitroClient) ExistsAuditnslogpolicy(key AuditnslogpolicyKey) (bool, error) {
	var results count_auditnslogpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("auditnslogpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuditnslogpolicy(key AuditnslogpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("auditnslogpolicy", id, qs)
}

//      UPDATE

type AuditnslogpolicyUpdate struct {
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Action string `json:"action,omitempty"`
}

func (resource Auditnslogpolicy) ToUpdate() AuditnslogpolicyUpdate {
	update := AuditnslogpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Action,
	}

	return update
}

type update_auditnslogpolicy_payload struct {
	Update AuditnslogpolicyUpdate `json:"auditnslogpolicy"`
}

func (c *NitroClient) UpdateAuditnslogpolicy(update AuditnslogpolicyUpdate) error {
	payload := update_auditnslogpolicy_payload{
		update,
	}

	return c.put("auditnslogpolicy", "", nil, payload)
}

//      UNSET

type AuditnslogpolicyUnset struct {
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
	Action bool   `json:"action,omitempty"`
}

func (resource Auditnslogpolicy) ToUnset() AuditnslogpolicyUnset {
	unset := AuditnslogpolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_auditnslogpolicy_payload struct {
	Unset AuditnslogpolicyUnset `json:"auditnslogpolicy"`
}

func (c *NitroClient) UnsetAuditnslogpolicy(unset AuditnslogpolicyUnset) error {
	payload := unset_auditnslogpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("auditnslogpolicy", "", qs, payload)
}

//      RENAME
//      TODO
