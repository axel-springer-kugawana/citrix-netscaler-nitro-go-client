package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Rewriteaction struct {
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Name              string `json:"name,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Search            string `json:"search,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Target            string `json:"target,omitempty"`
	Type              string `json:"type,omitempty"`
}

type RewriteactionKey struct {
	Name string
}

func (resource Rewriteaction) ToKey() RewriteactionKey {
	key := RewriteactionKey{
		resource.Name,
	}

	return key
}

func (key RewriteactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key RewriteactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key RewriteactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_rewriteaction_payload struct {
	Resource Rewriteaction `json:"rewriteaction"`
}

func (c *NitroClient) AddRewriteaction(resource Rewriteaction) error {
	payload := add_rewriteaction_payload{
		resource,
	}

	return c.post("rewriteaction", "", nil, payload)
}

//      LIST

type list_rewriteaction_result struct {
	Results []Rewriteaction `json:"rewriteaction"`
}

func (c *NitroClient) ListRewriteaction() ([]Rewriteaction, error) {
	results := list_rewriteaction_result{}

	if err := c.get("rewriteaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_rewriteaction_result struct {
	Results []Rewriteaction `json:"rewriteaction"`
}

func (c *NitroClient) GetRewriteaction(key RewriteactionKey) (*Rewriteaction, error) {
	var results get_rewriteaction_result

	id, qs := key.to_id_args()

	if err := c.get("rewriteaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one rewriteaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("rewriteaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_rewriteaction_result struct {
	Results []Count `json:"rewriteaction"`
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

//      EXISTS

func (c *NitroClient) ExistsRewriteaction(key RewriteactionKey) (bool, error) {
	var results count_rewriteaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("rewriteaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteRewriteaction(key RewriteactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("rewriteaction", id, qs)
}

//      UPDATE

type RewriteactionUpdate struct {
	Name              string `json:"name,omitempty"`
	Target            string `json:"target,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Pattern           string `json:"pattern,omitempty"`
	Search            string `json:"search,omitempty"`
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Refinesearch      string `json:"refinesearch,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

func (resource Rewriteaction) ToUpdate() RewriteactionUpdate {
	update := RewriteactionUpdate{
		resource.Name,
		resource.Target,
		resource.Stringbuilderexpr,
		resource.Pattern,
		resource.Search,
		resource.Bypasssafetycheck,
		resource.Refinesearch,
		resource.Comment,
	}

	return update
}

type update_rewriteaction_payload struct {
	Update RewriteactionUpdate `json:"rewriteaction"`
}

func (c *NitroClient) UpdateRewriteaction(update RewriteactionUpdate) error {
	payload := update_rewriteaction_payload{
		update,
	}

	return c.put("rewriteaction", "", nil, payload)
}

//      UNSET

type RewriteactionUnset struct {
	Name              string `json:"name,omitempty"`
	Target            bool   `json:"target,omitempty"`
	Stringbuilderexpr bool   `json:"stringbuilderexpr,omitempty"`
	Pattern           bool   `json:"pattern,omitempty"`
	Search            bool   `json:"search,omitempty"`
	Bypasssafetycheck bool   `json:"bypasssafetycheck,omitempty"`
	Refinesearch      bool   `json:"refinesearch,omitempty"`
	Comment           bool   `json:"comment,omitempty"`
}

func (resource Rewriteaction) ToUnset() RewriteactionUnset {
	unset := RewriteactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_rewriteaction_payload struct {
	Unset RewriteactionUnset `json:"rewriteaction"`
}

func (c *NitroClient) UnsetRewriteaction(unset RewriteactionUnset) error {
	payload := unset_rewriteaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("rewriteaction", "", qs, payload)
}

//      RENAME
//      TODO
