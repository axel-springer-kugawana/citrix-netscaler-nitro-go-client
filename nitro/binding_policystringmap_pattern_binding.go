package nitro

import (
	"fmt"
	"strconv"
	"strings"
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

type add_policystringmap_pattern_binding_payload struct {
	Resource PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

type get_policystringmap_pattern_binding_result struct {
	Results []PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

type count_policystringmap_pattern_binding_result struct {
	Results []Count `json:"policystringmap_pattern_binding"`
}

func policystringmap_pattern_binding_key_to_id_args(key PolicystringmapPatternBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "key:"+key.Key)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddPolicystringmapPatternBinding(binding PolicystringmapPatternBinding) error {
	payload := add_policystringmap_pattern_binding_payload{
		binding,
	}

	return c.put("policystringmap_pattern_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountPolicystringmapPatternBinding() (int, error) {
	var results count_policystringmap_pattern_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("policystringmap_pattern_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountPolicystringmapPatternBinding(id string) (int, error) {
	var results count_policystringmap_pattern_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policystringmap_pattern_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicystringmapPatternBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountPolicystringmapPatternBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListPolicystringmapPatternBinding() ([]PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("policystringmap_pattern_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListPolicystringmapPatternBinding(id string) ([]PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding_result

	if err := c.get("policystringmap_pattern_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) (*PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding_result

	id, qs := policystringmap_pattern_binding_key_to_id_args(key)

	if err := c.get("policystringmap_pattern_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policystringmap_pattern_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policystringmap_pattern_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) error {
	id, qs := policystringmap_pattern_binding_key_to_id_args(key)

	return c.delete("policystringmap_pattern_binding", id, qs)
}
