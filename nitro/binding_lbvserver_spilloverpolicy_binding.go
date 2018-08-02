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

type get_lbvserver_spilloverpolicy_binding struct {
	Results []LbvserverSpilloverpolicyBinding `json:"lbvserver_spilloverpolicy_binding"`
}

type add_lbvserver_spilloverpolicy_binding_payload struct {
	lbvserver_spilloverpolicy_binding LbvserverSpilloverpolicyBinding
}

func lbvserver_spilloverpolicy_binding_key_to_id_args(key LbvserverSpilloverpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "policyname:"+key.Policyname)
	args = append(args, "bindpoint:"+key.Bindpoint)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbvserverSpilloverpolicyBinding(binding LbvserverSpilloverpolicyBinding) error {
	payload := add_lbvserver_spilloverpolicy_binding_payload{
		binding,
	}

	return c.put("lbvserver_spilloverpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) ([]LbvserverSpilloverpolicyBinding, error) {
	var results get_lbvserver_spilloverpolicy_binding

	id, args := lbvserver_spilloverpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_spilloverpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbvserverSpilloverpolicyBinding() ([]LbvserverSpilloverpolicyBinding, error) {
	var results get_lbvserver_spilloverpolicy_binding

	if err := c.get("lbvserver_spilloverpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) (*LbvserverSpilloverpolicyBinding, error) {
	var results get_lbvserver_spilloverpolicy_binding

	id, args := lbvserver_spilloverpolicy_binding_key_to_id_args(key)

	if err := c.get("lbvserver_spilloverpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver_spilloverpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver_spilloverpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) error {
	id, args := lbvserver_spilloverpolicy_binding_key_to_id_args(key)

	return c.delete("lbvserver_spilloverpolicy_binding", id, "", args)
}
