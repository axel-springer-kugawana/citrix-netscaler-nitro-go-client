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

func (resource PolicystringmapPatternBinding) ToKey() PolicystringmapPatternBindingKey {
	key := PolicystringmapPatternBindingKey{
		resource.Name,
		resource.Key,
	}

	return key
}

func (key PolicystringmapPatternBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name
	args = append(args, "key:"+key.Key)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_policystringmap_pattern_binding_payload struct {
	Resource PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

func (c *NitroClient) AddPolicystringmapPatternBinding(resource PolicystringmapPatternBinding) error {
	payload := add_policystringmap_pattern_binding_payload{
		resource,
	}

	return c.put("policystringmap_pattern_binding", "", nil, payload)
}

//      LIST

type list_policystringmap_pattern_binding_result struct {
	Results []PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

func (c *NitroClient) ListPolicystringmapPatternBinding() ([]PolicystringmapPatternBinding, error) {
	results := list_policystringmap_pattern_binding_result{}

	if err := c.get("policystringmap_pattern_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_policystringmap_pattern_binding_result struct {
	Results []PolicystringmapPatternBinding `json:"policystringmap_pattern_binding"`
}

func (c *NitroClient) GetPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) (*PolicystringmapPatternBinding, error) {
	var results get_policystringmap_pattern_binding_result

	id, qs := key.to_id_args()

	if err := c.get("policystringmap_pattern_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policystringmap_pattern_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policystringmap_pattern_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_policystringmap_pattern_binding_result struct {
	Results []Count `json:"policystringmap_pattern_binding"`
}

func (c *NitroClient) ExistsPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) (bool, error) {
	var results count_policystringmap_pattern_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("policystringmap_pattern_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("policystringmap_pattern_binding", id, qs)
}
