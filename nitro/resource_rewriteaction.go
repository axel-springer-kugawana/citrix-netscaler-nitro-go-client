package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Rewriteaction struct {
	Name              string `json:"name"`
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Search            string `json:"search,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Target            string `json:"target,omitempty"`
	Type              string `json:"type,omitempty"`
}

type RewriteactionKey struct {
	Name string `json:"name"`
}

type RewriteactionUnset struct {
	Name              string `json:"name"`
	Target            bool   `json:"target,string,omitempty"`
	Stringbuilderexpr bool   `json:"stringbuilderexpr,string,omitempty"`
	Pattern           bool   `json:"pattern,string,omitempty"`
	Search            bool   `json:"search,string,omitempty"`
	Bypasssafetycheck bool   `json:"bypasssafetycheck,string,omitempty"`
	Refinesearch      bool   `json:"refinesearch,string,omitempty"`
	Comment           bool   `json:"comment,string,omitempty"`
}

type update_rewriteaction struct {
	Name              string `json:"name"`
	Target            string `json:"target,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Search            string `json:"search,omitempty"`
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

type rename_rewriteaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_rewriteaction_payload struct {
	Resource Rewriteaction `json:"rewriteaction"`
}

type rename_rewriteaction_payload struct {
	Rename rename_rewriteaction `json:"rewriteaction"`
}

type unset_rewriteaction_payload struct {
	Unset RewriteactionUnset `json:"rewriteaction"`
}

type update_rewriteaction_payload struct {
	Update update_rewriteaction `json:"rewriteaction"`
}

type get_rewriteaction_result struct {
	Results []Rewriteaction `json:"rewriteaction"`
}

type count_rewriteaction_result struct {
	Results []Count `json:"rewriteaction"`
}

func rewriteaction_key_to_id_args(key RewriteactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddRewriteaction(resource Rewriteaction) error {
	payload := add_rewriteaction_payload{
		resource,
	}

	return c.post("rewriteaction", "", nil, payload)
}

func (c *NitroClient) RenameRewriteaction(name string, newName string) error {
	payload := rename_rewriteaction_payload{
		rename_rewriteaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("rewriteaction", "", qs, payload)
}

func (c *NitroClient) CountRewriteaction() (int, error) {
	var results count_rewriteaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("rewriteaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsRewriteaction(key RewriteactionKey) (bool, error) {
	var results count_rewriteaction_result

	id, qs := rewriteaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("rewriteaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListRewriteaction() ([]Rewriteaction, error) {
	var results get_rewriteaction_result

	if err := c.get("rewriteaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetRewriteaction(key RewriteactionKey) (*Rewriteaction, error) {
	var results get_rewriteaction_result

	id, qs := rewriteaction_key_to_id_args(key)

	if err := c.get("rewriteaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one rewriteaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("rewriteaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteRewriteaction(key RewriteactionKey) error {
	id, qs := rewriteaction_key_to_id_args(key)

	return c.delete("rewriteaction", id, qs)
}

func (c *NitroClient) UnsetRewriteaction(unset RewriteactionUnset) error {
	payload := unset_rewriteaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("rewriteaction", "", qs, payload)
}

func (c *NitroClient) UpdateRewriteaction(resource Rewriteaction) error {
	payload := update_rewriteaction_payload{
		update_rewriteaction{
			resource.Name,
			resource.Target,
			resource.Stringbuilderexpr,
			resource.Pattern,
			resource.Search,
			resource.Bypasssafetycheck,
			resource.Refinesearch,
			resource.Comment,
		},
	}

	return c.put("rewriteaction", "", nil, payload)
}
