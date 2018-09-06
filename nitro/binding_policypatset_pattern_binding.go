package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type PolicypatsetPatternBinding struct {
	Charset string `json:"charset,omitempty"`
	Index   int    `json:"index,string,omitempty"`
	Name    string `json:"name,omitempty"`
	String  string `json:"string,omitempty"`
}

type PolicypatsetPatternBindingKey struct {
	Name   string
	String string
}

func (resource PolicypatsetPatternBinding) ToKey() PolicypatsetPatternBindingKey {
	key := PolicypatsetPatternBindingKey{
		resource.Name,
		resource.String,
	}

	return key
}

func (key PolicypatsetPatternBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "string:"+key.String)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key PolicypatsetPatternBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PolicypatsetPatternBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_policypatset_pattern_binding_payload struct {
	Resource PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

func (c *NitroClient) AddPolicypatsetPatternBinding(resource PolicypatsetPatternBinding) error {
	payload := add_policypatset_pattern_binding_payload{
		resource,
	}

	return c.put("policypatset_pattern_binding", "", nil, payload)
}

//      LIST

type list_policypatset_pattern_binding_result struct {
	Results []PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

func (c *NitroClient) ListPolicypatsetPatternBinding() ([]PolicypatsetPatternBinding, error) {
	results := list_policypatset_pattern_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("policypatset_pattern_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_policypatset_pattern_binding_result struct {
	Results []Count `json:"policypatset_pattern_binding"`
}

func (c *NitroClient) CountPolicypatsetPatternBinding() (int, error) {
	results := count_policypatset_pattern_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("policypatset_pattern_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_policypatset_pattern_binding_result struct {
	Results []PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

func (c *NitroClient) GetPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (*PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("policypatset_pattern_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policypatset_pattern_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policypatset_pattern_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (bool, error) {
	var results count_policypatset_pattern_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("policypatset_pattern_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("policypatset_pattern_binding", id, qs)
}
