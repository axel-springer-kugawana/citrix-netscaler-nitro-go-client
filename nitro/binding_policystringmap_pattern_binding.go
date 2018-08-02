package nitro

import (
	"fmt"
	"strconv"
)

type PolicystringmapPatternBinding struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PolicystringmapPatternBindingKey struct {
	Name string
	Key  string
}

type get_policystringmap_pattern_binding struct {
	Results []PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

type add_policystringmap_pattern_binding_payload struct {
	policystringmap_pattern_binding PolicystringmapPatternBinding
}

func policystringmap_pattern_binding_key_to_id_args(key PolicystringmapPatternBindingKey) (string, string) {
	var _ = strconv.Itoa

	result := ""

	result = result + ",name:" + key.Name
	result = result + ",key:" + key.Key
	return "", result
}

func (c *NitroClient) AddPolicystringmapPatternBinding(binding PolicystringmapPatternBinding) error {
	payload := add_policystringmap_pattern_binding_payload{
		binding,
	}

	return c.put("policystringmap_pattern_binding", "", "", "", payload)
}

func (c *NitroClient) ListPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) ([]PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding

	id, args := policystringmap_pattern_binding_key_to_id_args(key)

	if err := c.get("policystringmap_pattern_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListPolicystringmapPatternBinding() ([]PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding

	if err := c.get("policystringmap_pattern_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) (*PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding

	id, args := policystringmap_pattern_binding_key_to_id_args(key)

	if err := c.get("policystringmap_pattern_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policystringmap_pattern_binding element found")
		} else if len(results.Results) < 1 {
			//                        return nil, fmt.Errorf("policystringmap_pattern_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) error {
	id, args := policystringmap_pattern_binding_key_to_id_args(key)

	return c.delete("policystringmap_pattern_binding", id, "", args)
}
