package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverScpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverScpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource LbvserverScpolicyBinding) ToKey() LbvserverScpolicyBindingKey {
	key := LbvserverScpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key LbvserverScpolicyBindingKey) to_id_args() (string, map[string]string) {
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

type add_lbvserver_scpolicy_binding_payload struct {
	Resource LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

func (c *NitroClient) AddLbvserverScpolicyBinding(resource LbvserverScpolicyBinding) error {
	payload := add_lbvserver_scpolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_scpolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_scpolicy_binding_result struct {
	Results []LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

func (c *NitroClient) ListLbvserverScpolicyBinding() ([]LbvserverScpolicyBinding, error) {
	results := list_lbvserver_scpolicy_binding_result{}

	if err := c.get("lbvserver_scpolicy_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbvserver_scpolicy_binding_result struct {
	Results []LbvserverScpolicyBinding `json:"lbvserver_scpolicy_binding"`
}

func (c *NitroClient) GetLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (*LbvserverScpolicyBinding, error) {
	var results get_lbvserver_scpolicy_binding_result

	id, qs := key.to_id_args()

	if err := c.get("lbvserver_scpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_scpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_scpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_lbvserver_scpolicy_binding_result struct {
	Results []Count `json:"lbvserver_scpolicy_binding"`
}

func (c *NitroClient) ExistsLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) (bool, error) {
	var results count_lbvserver_scpolicy_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbvserver_scpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverScpolicyBinding(key LbvserverScpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_scpolicy_binding", id, qs)
}
