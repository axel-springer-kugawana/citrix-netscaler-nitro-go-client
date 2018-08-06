package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policyexpression struct {
	Name                  string `json:"name"`
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Description           string `json:"description,omitempty"`
	Value                 string `json:"value,omitempty"`
}

func policyexpression_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type PolicyexpressionUnset struct {
	Name                  string `json:"name"`
	Value                 bool   `json:"value,omitempty"`
	Description           bool   `json:"description,omitempty"`
	Comment               bool   `json:"comment,omitempty"`
	Clientsecuritymessage bool   `json:"clientsecuritymessage,omitempty"`
}

type update_policyexpression struct {
	Name                  string `json:"name"`
	Value                 string `json:"value,omitempty"`
	Description           string `json:"description,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
}

type rename_policyexpression struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_policyexpression_payload struct {
	Resource Policyexpression `json:"policyexpression"`
}

type rename_policyexpression_payload struct {
	Rename rename_policyexpression `json:"policyexpression"`
}

type unset_policyexpression_payload struct {
	Unset PolicyexpressionUnset `json:"policyexpression"`
}

type update_policyexpression_payload struct {
	Update update_policyexpression `json:"policyexpression"`
}

type get_policyexpression_result struct {
	Results []Policyexpression `json:"policyexpression"`
}

type count_policyexpression_result struct {
	Results []Count `json:"policyexpression"`
}

func (c *NitroClient) AddPolicyexpression(resource Policyexpression) error {
	payload := add_policyexpression_payload{
		resource,
	}

	return c.post("policyexpression", "", nil, payload)
}

func (c *NitroClient) RenamePolicyexpression(name string, newName string) error {
	payload := rename_policyexpression_payload{
		rename_policyexpression{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("policyexpression", "", qs, payload)
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

func (c *NitroClient) ExistsPolicyexpression(key string) (bool, error) {
	var results count_policyexpression_result

	id, qs := policyexpression_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("policyexpression", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListPolicyexpression() ([]Policyexpression, error) {
	results := get_policyexpression_result{}

	if err := c.get("policyexpression", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicyexpression(key string) (*Policyexpression, error) {
	var results get_policyexpression_result

	id, qs := policyexpression_key_to_id_args(key)

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

func (c *NitroClient) DeletePolicyexpression(key string) error {
	id, qs := policyexpression_key_to_id_args(key)

	return c.delete("policyexpression", id, qs)
}

func (c *NitroClient) UnsetPolicyexpression(unset PolicyexpressionUnset) error {
	payload := unset_policyexpression_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("policyexpression", "", qs, payload)
}

func (c *NitroClient) UpdatePolicyexpression(resource Policyexpression) error {
	payload := update_policyexpression_payload{
		update_policyexpression{
			resource.Name,
			resource.Value,
			resource.Description,
			resource.Comment,
			resource.Clientsecuritymessage,
		},
	}

	return c.put("policyexpression", "", nil, payload)
}
