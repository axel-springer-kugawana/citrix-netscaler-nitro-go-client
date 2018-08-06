package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAuditnslogpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuditnslogpolicyBindingKey struct {
	Name       string
	Policyname string
}

type add_lbvserver_auditnslogpolicy_binding_payload struct {
	Resource LbvserverAuditnslogpolicyBinding `json:"lbvserver_auditnslogpolicy_binding"`
}

type get_lbvserver_auditnslogpolicy_binding_result struct {
	Results []LbvserverAuditnslogpolicyBinding `json:"lbvserver_auditnslogpolicy_binding"`
}

type count_lbvserver_auditnslogpolicy_binding_result struct {
	Results []Count `json:"lbvserver_auditnslogpolicy_binding"`
}

func lbvserver_auditnslogpolicy_binding_key_to_id_qs(key LbvserverAuditnslogpolicyBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbvserver_auditnslogpolicy_binding_key_to_id_args(key LbvserverAuditnslogpolicyBindingKey) (string, map[string]string) {
	return lbvserver_auditnslogpolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_auditnslogpolicy_binding_key_to_id_filter(key LbvserverAuditnslogpolicyBindingKey) (string, map[string]string) {
	return lbvserver_auditnslogpolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverAuditnslogpolicyBinding(binding LbvserverAuditnslogpolicyBinding) error {
	payload := add_lbvserver_auditnslogpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_auditnslogpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverAuditnslogpolicyBinding() (int, error) {
	var results count_lbvserver_auditnslogpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_auditnslogpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverAuditnslogpolicyBinding(id string) (int, error) {
	var results count_lbvserver_auditnslogpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) (bool, error) {
	var results count_lbvserver_auditnslogpolicy_binding_result

	id, qs := lbvserver_auditnslogpolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_auditnslogpolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverAuditnslogpolicyBinding() ([]LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_auditnslogpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverAuditnslogpolicyBinding(id string) ([]LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding_result

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) (*LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding_result

	id, qs := lbvserver_auditnslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_auditnslogpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_auditnslogpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) error {
	id, qs := lbvserver_auditnslogpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_auditnslogpolicy_binding", id, qs)
}
