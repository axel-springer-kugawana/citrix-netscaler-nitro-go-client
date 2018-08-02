package nitro

import (
	"fmt"
	"strconv"
)

type LbvserverAuditsyslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuditsyslogpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

type get_lbvserver_auditsyslogpolicy_binding struct {
	Results []LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

type add_lbvserver_auditsyslogpolicy_binding_payload struct {
	lbvserver_auditsyslogpolicy_binding LbvserverAuditsyslogpolicyBinding
}

func lbvserver_auditsyslogpolicy_binding_key_to_id_args(key LbvserverAuditsyslogpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",policyname:" + key.Policyname
	result = result + ",bindpoint:" + key.Bindpoint
	return "", result
}

func (c *NitroClient) AddLbvserverAuditsyslogpolicyBinding(binding LbvserverAuditsyslogpolicyBinding) error {
	payload := add_lbvserver_auditsyslogpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_auditsyslogpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) ([]LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding

	id, args := lbvserver_auditsyslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverAuditsyslogpolicyBinding() ([]LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding

	if err := c.get("lbvserver_auditsyslogpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (*LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding

	id, args := lbvserver_auditsyslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_auditsyslogpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("lbvserver_auditsyslogpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) error {
	id, args := lbvserver_auditsyslogpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_auditsyslogpolicy_binding", id, "", args)
}
