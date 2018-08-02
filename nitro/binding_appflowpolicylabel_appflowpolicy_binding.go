package nitro

import (
	"fmt"
	"strconv"
)

type AppflowpolicylabelAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invoke_labelname       string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type AppflowpolicylabelAppflowpolicyBindingKey struct {
	Labelname  string
	Policyname string
}

type get_appflowpolicylabel_appflowpolicy_binding struct {
	Results []AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

type add_appflowpolicylabel_appflowpolicy_binding_payload struct {
	appflowpolicylabel_appflowpolicy_binding AppflowpolicylabelAppflowpolicyBinding
}

func appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key AppflowpolicylabelAppflowpolicyBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",labelname:" + key.Labelname
	result = result + ",policyname:" + key.Policyname
	return "", result
}

func (c *NitroClient) AddAppflowpolicylabelAppflowpolicyBinding(binding AppflowpolicylabelAppflowpolicyBinding) error {
	payload := add_appflowpolicylabel_appflowpolicy_binding_payload{
		binding,
	}

	return c.put("appflowpolicylabel_appflowpolicy_binding", "", "", "", payload)
}

func (c *NitroClient) ListAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding

	id, args := appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListAppflowpolicylabelAppflowpolicyBinding() ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) (*AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding

	id, args := appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicylabel_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("appflowpolicylabel_appflowpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) error {
	id, args := appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key)

	return c.delete("appflowpolicylabel_appflowpolicy_binding", id, "", args)
}
