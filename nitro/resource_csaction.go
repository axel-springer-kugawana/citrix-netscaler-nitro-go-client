package nitro

import (
	"fmt"
	"net/url"
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

func (key CsactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CsactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CsactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type CsactionUpdate struct {
	Name              string `json:"name,omitempty"`
	Targetlbvserver   string `json:"targetlbvserver,omitempty"`
	Targetvserver     string `json:"targetvserver,omitempty"`
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
	Comment           string `json:"comment,omitempty"`
}

func (resource Csaction) ToUpdate() CsactionUpdate {
	update := CsactionUpdate{
		resource.Name,
		resource.Targetlbvserver,
		resource.Targetvserver,
		resource.Targetvserverexpr,
		resource.Comment,
	}

	return update
}

type update_csaction_payload struct {
	Update CsactionUpdate `json:"csaction"`
}

func (c *NitroClient) UpdateCsaction(update CsactionUpdate) error {
	payload := update_csaction_payload{
		update,
	}

	return c.put("csaction", "", nil, payload)
}

//      UNSET

type CsactionUnset struct {
	Name              string `json:"name,omitempty"`
	Targetlbvserver   bool   `json:"targetlbvserver,omitempty"`
	Targetvserver     bool   `json:"targetvserver,omitempty"`
	Targetvserverexpr bool   `json:"targetvserverexpr,omitempty"`
	Comment           bool   `json:"comment,omitempty"`
}

func (resource Csaction) ToUnset() CsactionUnset {
	unset := CsactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_csaction_payload struct {
	Unset CsactionUnset `json:"csaction"`
}

func (c *NitroClient) UnsetCsaction(unset CsactionUnset) error {
	payload := unset_csaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("csaction", "", qs, payload)
}

//      RENAME
//      TODO
