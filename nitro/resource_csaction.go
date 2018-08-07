package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Csaction struct {
	Comment           string `json:"comment,omitempty"`
	Name              string `json:"name,omitempty"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
}

type CsactionKey struct {
	Name string
}

func (resource Csaction) ToKey() CsactionKey {
	key := CsactionKey{
		resource.Name,
	}

	return key
}

func (key CsactionKey) to_id_args() (string, map[string]string) {
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

type add_csaction_payload struct {
	Resource Csaction `json:"csaction"`
}

func (c *NitroClient) AddCsaction(resource Csaction) error {
	payload := add_csaction_payload{
		resource,
	}

	return c.post("csaction", "", nil, payload)
}

//      LIST

type list_csaction_result struct {
	Results []Csaction `json:"csaction"`
}

func (c *NitroClient) ListCsaction() ([]Csaction, error) {
	results := list_csaction_result{}

	if err := c.get("csaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_csaction_result struct {
	Results []Csaction `json:"csaction"`
}

func (c *NitroClient) GetCsaction(key CsactionKey) (*Csaction, error) {
	var results get_csaction_result

	id, qs := key.to_id_args()

	if err := c.get("csaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_csaction_result struct {
	Results []Count `json:"csaction"`
}

func (c *NitroClient) CountCsaction() (int, error) {
	var results count_csaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("csaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCsaction(key CsactionKey) (bool, error) {
	var results count_csaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("csaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCsaction(key CsactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("csaction", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
