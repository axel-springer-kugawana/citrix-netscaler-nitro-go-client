package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cspolicy struct {
	Action     string `json:"action,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Logaction  string `json:"logaction,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Url        string `json:"url,omitempty"`
}

type CspolicyKey struct {
	Policyname string
}

func (resource Cspolicy) ToKey() CspolicyKey {
	key := CspolicyKey{
		resource.Policyname,
	}

	return key
}

func (key CspolicyKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Policyname

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_cspolicy_payload struct {
	Resource Cspolicy `json:"cspolicy"`
}

func (c *NitroClient) AddCspolicy(resource Cspolicy) error {
	payload := add_cspolicy_payload{
		resource,
	}

	return c.post("cspolicy", "", nil, payload)
}

//      LIST

type list_cspolicy_result struct {
	Results []Cspolicy `json:"cspolicy"`
}

func (c *NitroClient) ListCspolicy() ([]Cspolicy, error) {
	results := list_cspolicy_result{}

	if err := c.get("cspolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cspolicy_result struct {
	Results []Cspolicy `json:"cspolicy"`
}

func (c *NitroClient) GetCspolicy(key CspolicyKey) (*Cspolicy, error) {
	var results get_cspolicy_result

	id, qs := key.to_id_args()

	if err := c.get("cspolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cspolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cspolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cspolicy_result struct {
	Results []Count `json:"cspolicy"`
}

func (c *NitroClient) CountCspolicy() (int, error) {
	var results count_cspolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cspolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCspolicy(key CspolicyKey) (bool, error) {
	var results count_cspolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cspolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCspolicy(key CspolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("cspolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
