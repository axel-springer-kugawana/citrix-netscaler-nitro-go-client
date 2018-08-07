package nitro

import (
	"fmt"
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

func (key RewriteactionKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
