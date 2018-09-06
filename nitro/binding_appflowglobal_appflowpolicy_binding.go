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

func (resource AppflowglobalAppflowpolicyBinding) ToKey() AppflowglobalAppflowpolicyBindingKey {
	key := AppflowglobalAppflowpolicyBindingKey{
		resource.Labelname,
		resource.Policyname,
	}

	return key
}

func (key AppflowglobalAppflowpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Labelname
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AppflowglobalAppflowpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppflowglobalAppflowpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appflowglobal_appflowpolicy_binding_payload struct {
	Resource AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

func (c *NitroClient) AddAppflowglobalAppflowpolicyBinding(resource AppflowglobalAppflowpolicyBinding) error {
	payload := add_appflowglobal_appflowpolicy_binding_payload{
		resource,
	}

	return c.put("appflowglobal_appflowpolicy_binding", "", nil, payload)
}

//      LIST

type list_appflowglobal_appflowpolicy_binding_result struct {
	Results []AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

func (c *NitroClient) ListAppflowglobalAppflowpolicyBinding() ([]AppflowglobalAppflowpolicyBinding, error) {
	results := list_appflowglobal_appflowpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("appflowglobal_appflowpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_appflowglobal_appflowpolicy_binding_result struct {
	Results []Count `json:"appflowglobal_appflowpolicy_binding"`
}

func (c *NitroClient) CountAppflowglobalAppflowpolicyBinding() (int, error) {
	results := count_appflowglobal_appflowpolicy_binding_result{}

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

//      READ

type get_appflowglobal_appflowpolicy_binding_result struct {
	Results []AppflowglobalAppflowpolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
}

func (c *NitroClient) GetAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (*AppflowglobalAppflowpolicyBinding, error) {
	var results get_appflowglobal_appflowpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("appflowglobal_appflowpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowglobal_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowglobal_appflowpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) (bool, error) {
	var results count_appflowglobal_appflowpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("appflowglobal_appflowpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowglobalAppflowpolicyBinding(key AppflowglobalAppflowpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowglobal_appflowpolicy_binding", id, qs)
}
