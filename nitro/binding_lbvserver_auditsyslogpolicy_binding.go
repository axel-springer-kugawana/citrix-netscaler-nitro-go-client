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

func (resource LbvserverAuditsyslogpolicyBinding) ToKey() LbvserverAuditsyslogpolicyBindingKey {
	key := LbvserverAuditsyslogpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key LbvserverAuditsyslogpolicyBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_lbvserver_auditsyslogpolicy_binding_payload struct {
	Resource LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

func (c *NitroClient) AddLbvserverAuditsyslogpolicyBinding(resource LbvserverAuditsyslogpolicyBinding) error {
	payload := add_lbvserver_auditsyslogpolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_auditsyslogpolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_auditsyslogpolicy_binding_result struct {
	Results []LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

func (c *NitroClient) ListLbvserverAuditsyslogpolicyBinding() ([]LbvserverAuditsyslogpolicyBinding, error) {
	results := list_lbvserver_auditsyslogpolicy_binding_result{}

	if err := c.get("lbvserver_auditsyslogpolicy_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbvserver_auditsyslogpolicy_binding_result struct {
	Results []LbvserverAuditsyslogpolicyBinding `json:"lbvserver_auditsyslogpolicy_binding"`
}

func (c *NitroClient) GetLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (*LbvserverAuditsyslogpolicyBinding, error) {
	var results get_lbvserver_auditsyslogpolicy_binding_result

	id, qs := key.to_id_args()

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

//      EXISTS

type count_lbvserver_auditsyslogpolicy_binding_result struct {
	Results []Count `json:"lbvserver_auditsyslogpolicy_binding"`
}

func (c *NitroClient) ExistsLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (bool, error) {
	var results count_lbvserver_auditsyslogpolicy_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbvserver_auditsyslogpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_auditsyslogpolicy_binding", id, qs)
}
