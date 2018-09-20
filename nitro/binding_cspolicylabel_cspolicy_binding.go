package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type CspolicylabelCspolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Invoke_labelname       string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CspolicylabelCspolicyBindingKey struct {
	Labelname  string
	Policyname string
}

func (resource CspolicylabelCspolicyBinding) ToKey() CspolicylabelCspolicyBindingKey {
	key := CspolicylabelCspolicyBindingKey{
		resource.Labelname,
		resource.Policyname,
	}

	return key
}

func (key CspolicylabelCspolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Labelname)
	args = append(args, "policyname:"+url.QueryEscape(key.Policyname))

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CspolicylabelCspolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CspolicylabelCspolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cspolicylabel_cspolicy_binding_payload struct {
	Resource CspolicylabelCspolicyBinding `json:"cspolicylabel_cspolicy_binding"`
}

func (c *NitroClient) AddCspolicylabelCspolicyBinding(resource CspolicylabelCspolicyBinding) error {
	payload := add_cspolicylabel_cspolicy_binding_payload{
		resource,
	}

	return c.put("cspolicylabel_cspolicy_binding", "", nil, payload)
}

//      LIST

type list_cspolicylabel_cspolicy_binding_result struct {
	Results []CspolicylabelCspolicyBinding `json:"cspolicylabel_cspolicy_binding"`
}

func (c *NitroClient) ListCspolicylabelCspolicyBinding() ([]CspolicylabelCspolicyBinding, error) {
	results := list_cspolicylabel_cspolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("cspolicylabel_cspolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_cspolicylabel_cspolicy_binding_result struct {
	Results []Count `json:"cspolicylabel_cspolicy_binding"`
}

func (c *NitroClient) CountCspolicylabelCspolicyBinding() (int, error) {
	results := count_cspolicylabel_cspolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("cspolicylabel_cspolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_cspolicylabel_cspolicy_binding_result struct {
	Results []CspolicylabelCspolicyBinding `json:"cspolicylabel_cspolicy_binding"`
}

func (c *NitroClient) GetCspolicylabelCspolicyBinding(key CspolicylabelCspolicyBindingKey) (*CspolicylabelCspolicyBinding, error) {
	var results get_cspolicylabel_cspolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("cspolicylabel_cspolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cspolicylabel_cspolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cspolicylabel_cspolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsCspolicylabelCspolicyBinding(key CspolicylabelCspolicyBindingKey) (bool, error) {
	var results count_cspolicylabel_cspolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("cspolicylabel_cspolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCspolicylabelCspolicyBinding(key CspolicylabelCspolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("cspolicylabel_cspolicy_binding", id, qs)
}
