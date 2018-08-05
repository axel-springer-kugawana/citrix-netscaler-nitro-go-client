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

type add_appflowglobal_appflowpolicy_binding_payload struct {
	Resource AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

type get_appflowglobal_appflowpolicy_binding_result struct {
	Results []AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

type count_appflowglobal_appflowpolicy_binding_result struct {
	Results []Count `json:"appflowglobal_appflowpolicy_binding"`
}

func appflowglobal_appflowpolicy_binding_key_to_id_args(key AppflowglobalAppflowpolicyBindingKey) (string, map[string]string) {
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

func (c *NitroClient) AddAppflowglobalAppflowpolicyBinding(binding AppflowglobalAppflowpolicyBinding) error {
	payload := add_appflowglobal_appflowpolicy_binding_payload{
		binding,
	}

	return c.put("appflowglobal_appflowpolicy_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountAppflowglobalAppflowpolicyBinding() (int, error) {
	var results count_appflowglobal_appflowpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("appflowglobal_appflowpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountAppflowglobalAppflowpolicyBinding(id string) (int, error) {
	var results count_appflowglobal_appflowpolicy_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appflowglobal_appflowpolicy_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (bool, error) {
	// TODO : wrong implementation
	return false, nil
	//        if count, err := c.CountAppflowglobalAppflowpolicyBinding(id); err != nil {
	//                return false, err
	//        } else {
	//                return count == 1, nil
	//        }
}

func (c *NitroClient) BulkListAppflowglobalAppflowpolicyBinding() ([]AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("appflowglobal_appflowpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListAppflowglobalAppflowpolicyBinding(id string) ([]AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding_result

	if err := c.get("appflowglobal_appflowpolicy_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (*AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding_result

	id, qs := appflowglobal_appflowpolicy_binding_key_to_id_args(key)

	if err := c.get("appflowglobal_appflowpolicy_binding", id, qs, &results); err != nil {
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
	id, qs := appflowglobal_appflowpolicy_binding_key_to_id_args(key)

	return c.delete("appflowglobal_appflowpolicy_binding", id, qs)
}
