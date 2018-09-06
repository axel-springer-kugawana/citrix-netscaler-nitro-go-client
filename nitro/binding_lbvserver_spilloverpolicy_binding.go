package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbvserverSpilloverpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverSpilloverpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (resource LbvserverSpilloverpolicyBinding) ToKey() LbvserverSpilloverpolicyBindingKey {
	key := LbvserverSpilloverpolicyBindingKey{
		resource.Name,
		resource.Policyname,
		resource.Bindpoint,
	}

	return key
}

func (key LbvserverSpilloverpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbvserverSpilloverpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbvserverSpilloverpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbvserver_spilloverpolicy_binding_payload struct {
	Resource LbvserverSpilloverpolicyBinding `json:"lbvserver_spilloverpolicy_binding"`
}

func (c *NitroClient) AddLbvserverSpilloverpolicyBinding(resource LbvserverSpilloverpolicyBinding) error {
	payload := add_lbvserver_spilloverpolicy_binding_payload{
		resource,
	}

	return c.put("lbvserver_spilloverpolicy_binding", "", nil, payload)
}

//      LIST

type list_lbvserver_spilloverpolicy_binding_result struct {
	Results []LbvserverSpilloverpolicyBinding `json:"lbvserver_spilloverpolicy_binding"`
}

func (c *NitroClient) ListLbvserverSpilloverpolicyBinding() ([]LbvserverSpilloverpolicyBinding, error) {
	results := list_lbvserver_spilloverpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbvserver_spilloverpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbvserver_spilloverpolicy_binding_result struct {
	Results []Count `json:"lbvserver_spilloverpolicy_binding"`
}

func (c *NitroClient) CountLbvserverSpilloverpolicyBinding() (int, error) {
	results := count_lbvserver_spilloverpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbvserver_spilloverpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_lbvserver_spilloverpolicy_binding_result struct {
	Results []LbvserverSpilloverpolicyBinding `json:"lbvserver_spilloverpolicy_binding"`
}

func (c *NitroClient) GetLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) (*LbvserverSpilloverpolicyBinding, error) {
	var results get_lbvserver_spilloverpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("lbvserver_spilloverpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_spilloverpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver_spilloverpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) (bool, error) {
	var results count_lbvserver_spilloverpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbvserver_spilloverpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver_spilloverpolicy_binding", id, qs)
}
