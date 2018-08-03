package nitro

import (
	"fmt"
	"strconv"
	"strings"
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

type add_appflowpolicylabel_appflowpolicy_binding_payload struct {
	Resource AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

type get_appflowpolicylabel_appflowpolicy_binding_result struct {
	Results []AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

type count_appflowpolicylabel_appflowpolicy_binding_result struct {
	Results []Count `json:"appflowpolicylabel_appflowpolicy_binding"`
}

func appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key AppflowpolicylabelAppflowpolicyBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "labelname:"+key.Labelname)
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddAppflowpolicylabelAppflowpolicyBinding(binding AppflowpolicylabelAppflowpolicyBinding) error {
	payload := add_appflowpolicylabel_appflowpolicy_binding_payload{
		binding,
	}

	return c.put("appflowpolicylabel_appflowpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountAppflowpolicylabelAppflowpolicyBinding() (int, error) {
	var results count_appflowpolicylabel_appflowpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountAppflowpolicylabelAppflowpolicyBinding(id string) (int, error) {
	var results count_appflowpolicylabel_appflowpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppflowpolicylabelAppflowpolicyBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountAppflowpolicylabelAppflowpolicyBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListAppflowpolicylabelAppflowpolicyBinding() ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListAppflowpolicylabelAppflowpolicyBinding(id string) ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding_result

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) (*AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding_result

	id, qs := appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicylabel_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("appflowpolicylabel_appflowpolicy_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) error {
	id, qs := appflowpolicylabel_appflowpolicy_binding_key_to_id_args(key)

	return c.delete("appflowpolicylabel_appflowpolicy_binding", id, qs)
}
