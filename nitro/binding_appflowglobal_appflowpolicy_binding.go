package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type AppflowglobalAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type AppflowglobalAppflowpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

type get_appflowglobal_appflowpolicy_binding struct {
	Results []AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

type add_appflowglobal_appflowpolicy_binding_payload struct {
	appflowglobal_appflowpolicy_binding AppflowglobalAppflowpolicyBinding
}

func appflowglobal_appflowpolicy_binding_key_to_id_args(key AppflowglobalAppflowpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "labelname:"+key.Labelname)
	args = append(args, "policyname:"+key.Policyname)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddAppflowglobalAppflowpolicyBinding(binding AppflowglobalAppflowpolicyBinding) error {
	payload := add_appflowglobal_appflowpolicy_binding_payload{
		binding,
	}

	return c.put("appflowglobal_appflowpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) ([]AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding

	id, args := appflowglobal_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowglobal_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListAppflowglobalAppflowpolicyBinding() ([]AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding

	if err := c.get("appflowglobal_appflowpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (*AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding

	id, args := appflowglobal_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowglobal_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowglobal_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("appflowglobal_appflowpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) error {
	id, args := appflowglobal_appflowpolicy_binding_key_to_id_args(key)

	return c.delete("appflowglobal_appflowpolicy_binding", id, "", args)
}
