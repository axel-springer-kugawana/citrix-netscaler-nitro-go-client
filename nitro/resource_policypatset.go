package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policypatset struct {
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
	Name      string `json:"name,omitempty"`
}

type PolicypatsetKey struct {
	Name string
}

func (resource Policypatset) ToKey() PolicypatsetKey {
	key := PolicypatsetKey{
		resource.Name,
	}

	return key
}

func (key PolicypatsetKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key PolicypatsetKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PolicypatsetKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_policypatset_payload struct {
	Resource Policypatset `json:"policypatset"`
}

func (c *NitroClient) AddPolicypatset(resource Policypatset) error {
	payload := add_policypatset_payload{
		resource,
	}

	return c.post("policypatset", "", nil, payload)
}

//      LIST

type list_policypatset_result struct {
	Results []Policypatset `json:"policypatset"`
}

func (c *NitroClient) ListPolicypatset() ([]Policypatset, error) {
	results := list_policypatset_result{}

	if err := c.get("policypatset", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_policypatset_result struct {
	Results []Policypatset `json:"policypatset"`
}

func (c *NitroClient) GetPolicypatset(key PolicypatsetKey) (*Policypatset, error) {
	var results get_policypatset_result

	id, qs := key.to_id_args()

	if err := c.get("policypatset", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policypatset element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policypatset element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_policypatset_result struct {
	Results []Count `json:"policypatset"`
}

func (c *NitroClient) CountPolicypatset() (int, error) {
	var results count_policypatset_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policypatset", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicypatset(key PolicypatsetKey) (bool, error) {
	var results count_policypatset_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("policypatset", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicypatset(key PolicypatsetKey) error {
	id, qs := key.to_id_args()

	return c.delete("policypatset", id, qs)
}

//      RENAME
//      TODO
