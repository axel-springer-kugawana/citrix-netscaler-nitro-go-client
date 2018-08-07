package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policyexpression struct {
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Description           string `json:"description,omitempty"`
	Name                  string `json:"name,omitempty"`
	Value                 string `json:"value,omitempty"`
}

type PolicyexpressionKey struct {
	Name string
}

func (resource Policyexpression) ToKey() PolicyexpressionKey {
	key := PolicyexpressionKey{
		resource.Name,
	}

	return key
}

func (key PolicyexpressionKey) to_id_args() (string, map[string]string) {
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

type add_policyexpression_payload struct {
	Resource Policyexpression `json:"policyexpression"`
}

func (c *NitroClient) AddPolicyexpression(resource Policyexpression) error {
	payload := add_policyexpression_payload{
		resource,
	}

	return c.post("policyexpression", "", nil, payload)
}

//      LIST

type list_policyexpression_result struct {
	Results []Policyexpression `json:"policyexpression"`
}

func (c *NitroClient) ListPolicyexpression() ([]Policyexpression, error) {
	results := list_policyexpression_result{}

	if err := c.get("policyexpression", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_policyexpression_result struct {
	Results []Policyexpression `json:"policyexpression"`
}

func (c *NitroClient) GetPolicyexpression(key PolicyexpressionKey) (*Policyexpression, error) {
	var results get_policyexpression_result

	id, qs := key.to_id_args()

	if err := c.get("policyexpression", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policyexpression element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policyexpression element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_policyexpression_result struct {
	Results []Count `json:"policyexpression"`
}

func (c *NitroClient) CountPolicyexpression() (int, error) {
	var results count_policyexpression_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policyexpression", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicyexpression(key PolicyexpressionKey) (bool, error) {
	var results count_policyexpression_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("policyexpression", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicyexpression(key PolicyexpressionKey) error {
	id, qs := key.to_id_args()

	return c.delete("policyexpression", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
