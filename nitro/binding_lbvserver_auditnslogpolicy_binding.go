package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAuditnslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
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
	Bindpoint  string
}

type get_lbvserver_auditnslogpolicy_binding struct {
	Results []LbvserverAuditnslogpolicyBinding `json:"lbvserver_auditnslogpolicy_binding"`
}

type add_lbvserver_auditnslogpolicy_binding_payload struct {
	lbvserver_auditnslogpolicy_binding LbvserverAuditnslogpolicyBinding
}

func lbvserver_auditnslogpolicy_binding_key_to_id_args(key LbvserverAuditnslogpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverAuditnslogpolicyBinding(binding LbvserverAuditnslogpolicyBinding) error {
	payload := add_lbvserver_auditnslogpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_auditnslogpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) ([]LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding

	id, args := lbvserver_auditnslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAuditnslogpolicyBinding() ([]LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding

	if err := c.get("lbvserver_auditnslogpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) (*LbvserverAuditnslogpolicyBinding, error) {
	var results get_lbvserver_auditnslogpolicy_binding

	id, args := lbvserver_auditnslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditnslogpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_auditnslogpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_auditnslogpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) error {
	id, args := lbvserver_auditnslogpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_auditnslogpolicy_binding", id, "", args)
}
