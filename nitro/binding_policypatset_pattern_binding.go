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

type get_policypatset_pattern_binding struct {
	Results []PolicypatsetPatternBinding `json:"policypatset_pattern_binding"`
}

type add_policypatset_pattern_binding_payload struct {
	policypatset_pattern_binding PolicypatsetPatternBinding
}

func policypatset_pattern_binding_key_to_id_args(key PolicypatsetPatternBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "string:"+key.String)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddPolicypatsetPatternBinding(binding PolicypatsetPatternBinding) error {
	payload := add_policypatset_pattern_binding_payload{
		binding,
	}

	return c.put("policypatset_pattern_binding", "", "", "", payload)
}

func (c *NitroClient) ListPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) ([]PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding

	id, args := policypatset_pattern_binding_key_to_id_args(key)

	if err := c.get("policypatset_pattern_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListPolicypatsetPatternBinding() ([]PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding

	if err := c.get("policypatset_pattern_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (*PolicypatsetPatternBinding, error) {
	var results get_policypatset_pattern_binding

	id, args := policypatset_pattern_binding_key_to_id_args(key)

	if err := c.get("policypatset_pattern_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policypatset_pattern_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policypatset_pattern_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) error {
	id, args := policypatset_pattern_binding_key_to_id_args(key)

	return c.delete("policypatset_pattern_binding", id, "", args)
}
