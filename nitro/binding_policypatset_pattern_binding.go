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

type add_policypatset_pattern_binding_payload struct {
	Resource PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

type get_policypatset_pattern_binding_result struct {
	Results []PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

type count_policypatset_pattern_binding_result struct {
	Results []Count `json:"policypatset_pattern_binding"`
}

func policypatset_pattern_binding_key_to_id_qs(key PolicypatsetPatternBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "string:"+key.String)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func policypatset_pattern_binding_key_to_id_args(key PolicypatsetPatternBindingKey) (string, map[string]string) {
	return policypatset_pattern_binding_key_to_id_qs(key, "args")
}

func policypatset_pattern_binding_key_to_id_filter(key PolicypatsetPatternBindingKey) (string, map[string]string) {
	return policypatset_pattern_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddPolicypatsetPatternBinding(binding PolicypatsetPatternBinding) error {
	payload := add_policypatset_pattern_binding_payload{
		binding,
	}

	return c.put("policypatset_pattern_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountPolicypatsetPatternBinding() (int, error) {
	var results count_policypatset_pattern_binding_result

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

func (c *NitroClient) CountPolicypatsetPatternBinding(id string) (int, error) {
	var results count_policypatset_pattern_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policypatset_pattern_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (bool, error) {
	var results count_policypatset_pattern_binding_result

	id, qs := policypatset_pattern_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("policypatset_pattern_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one policypatset_pattern_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListPolicypatsetPatternBinding() ([]PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("policypatset_pattern_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListPolicypatsetPatternBinding(id string) ([]PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding_result

	if err := c.get("policypatset_pattern_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (*PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding_result

	id, qs := policypatset_pattern_binding_key_to_id_args(key)

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

func (c *NitroClient) DeletePolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) error {
	id, qs := policypatset_pattern_binding_key_to_id_args(key)

	return c.delete("policypatset_pattern_binding", id, qs)
}
