package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverAuditsyslogpolicyBinding struct {
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
}

type add_lbvserver_auditsyslogpolicy_binding_payload struct {
	Resource LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

type get_lbvserver_auditsyslogpolicy_binding_result struct {
	Results []LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

type count_lbvserver_auditsyslogpolicy_binding_result struct {
	Results []Count `json:"lbvserver_auditsyslogpolicy_binding"`
}

func lbvserver_auditsyslogpolicy_binding_key_to_id_qs(key LbvserverAuditsyslogpolicyBindingKey, arg string) (string, map[string]string) {
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

func lbvserver_auditsyslogpolicy_binding_key_to_id_args(key LbvserverAuditsyslogpolicyBindingKey) (string, map[string]string) {
	return lbvserver_auditsyslogpolicy_binding_key_to_id_qs(key, "args")
}

func lbvserver_auditsyslogpolicy_binding_key_to_id_filter(key LbvserverAuditsyslogpolicyBindingKey) (string, map[string]string) {
	return lbvserver_auditsyslogpolicy_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbvserverAuditsyslogpolicyBinding(binding LbvserverAuditsyslogpolicyBinding) error {
	payload := add_lbvserver_auditsyslogpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_auditsyslogpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbvserverAuditsyslogpolicyBinding() (int, error) {
	var results count_lbvserver_auditsyslogpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_auditsyslogpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbvserverAuditsyslogpolicyBinding(id string) (int, error) {
	var results count_lbvserver_auditsyslogpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (bool, error) {
	var results count_lbvserver_auditsyslogpolicy_binding_result

	id, qs := lbvserver_auditsyslogpolicy_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbvserver_auditsyslogpolicy_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbvserverAuditsyslogpolicyBinding() ([]LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_auditsyslogpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbvserverAuditsyslogpolicyBinding(id string) ([]LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding_result

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (*LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding_result

	id, qs := lbvserver_auditsyslogpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_auditsyslogpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_auditsyslogpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) error {
	id, qs := lbvserver_auditsyslogpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_auditsyslogpolicy_binding", id, qs)
}
