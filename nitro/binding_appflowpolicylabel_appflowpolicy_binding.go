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

func (resource AppflowpolicylabelAppflowpolicyBinding) ToKey() AppflowpolicylabelAppflowpolicyBindingKey {
	key := AppflowpolicylabelAppflowpolicyBindingKey{
		resource.Labelname,
		resource.Policyname,
	}

	return key
}

func (key AppflowpolicylabelAppflowpolicyBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Labelname
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_appflowpolicylabel_appflowpolicy_binding_payload struct {
	Resource AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

func (c *NitroClient) AddAppflowpolicylabelAppflowpolicyBinding(resource AppflowpolicylabelAppflowpolicyBinding) error {
	payload := add_appflowpolicylabel_appflowpolicy_binding_payload{
		resource,
	}

	return c.put("appflowpolicylabel_appflowpolicy_binding", "", nil, payload)
}

//      LIST

type list_appflowpolicylabel_appflowpolicy_binding_result struct {
	Results []AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

func (c *NitroClient) ListAppflowpolicylabelAppflowpolicyBinding() ([]AppflowpolicylabelAppflowpolicyBinding, error) {
	results := list_appflowpolicylabel_appflowpolicy_binding_result{}

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appflowpolicylabel_appflowpolicy_binding_result struct {
	Results []AppflowpolicylabelAppflowpolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
}

func (c *NitroClient) GetAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) (*AppflowpolicylabelAppflowpolicyBinding, error) {
	var results get_appflowpolicylabel_appflowpolicy_binding_result

	id, qs := key.to_id_args()

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appflowpolicylabel_appflowpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appflowpolicylabel_appflowpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_appflowpolicylabel_appflowpolicy_binding_result struct {
	Results []Count `json:"appflowpolicylabel_appflowpolicy_binding"`
}

func (c *NitroClient) ExistsAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) (bool, error) {
	var results count_appflowpolicylabel_appflowpolicy_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appflowpolicylabel_appflowpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppflowpolicylabelAppflowpolicyBinding(key AppflowpolicylabelAppflowpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("appflowpolicylabel_appflowpolicy_binding", id, qs)
}
