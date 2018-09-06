package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Caaction struct {
	Accumressize int    `json:"accumressize,string,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
}

type CaactionKey struct {
	Name string
}

func (resource Caaction) ToKey() CaactionKey {
	key := CaactionKey{
		resource.Name,
	}

	return key
}

func (key CaactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CaactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CaactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_caaction_payload struct {
	Resource Caaction `json:"caaction"`
}

func (c *NitroClient) AddCaaction(resource Caaction) error {
	payload := add_caaction_payload{
		resource,
	}

	return c.post("caaction", "", nil, payload)
}

//      LIST

type list_caaction_result struct {
	Results []Caaction `json:"caaction"`
}

func (c *NitroClient) ListCaaction() ([]Caaction, error) {
	results := list_caaction_result{}

	if err := c.get("caaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_caaction_result struct {
	Results []Caaction `json:"caaction"`
}

func (c *NitroClient) GetCaaction(key CaactionKey) (*Caaction, error) {
	var results get_caaction_result

	id, qs := key.to_id_args()

	if err := c.get("caaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one caaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("caaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_caaction_result struct {
	Results []Count `json:"caaction"`
}

func (c *NitroClient) CountCaaction() (int, error) {
	var results count_caaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("caaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCaaction(key CaactionKey) (bool, error) {
	var results count_caaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("caaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCaaction(key CaactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("caaction", id, qs)
}

//      UPDATE

type CaactionUpdate struct {
	Name         string `json:"name,omitempty"`
	Accumressize int    `json:"accumressize,string,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Type         string `json:"type,omitempty"`
}

func (resource Caaction) ToUpdate() CaactionUpdate {
	update := CaactionUpdate{
		resource.Name,
		resource.Accumressize,
		resource.Lbvserver,
		resource.Comment,
		resource.Type,
	}

	return update
}

type update_caaction_payload struct {
	Update CaactionUpdate `json:"caaction"`
}

func (c *NitroClient) UpdateCaaction(update CaactionUpdate) error {
	payload := update_caaction_payload{
		update,
	}

	return c.put("caaction", "", nil, payload)
}

//      UNSET

type CaactionUnset struct {
	Name         string `json:"name,omitempty"`
	Accumressize bool   `json:"accumressize,omitempty"`
	Lbvserver    bool   `json:"lbvserver,omitempty"`
	Comment      bool   `json:"comment,omitempty"`
	Type         bool   `json:"type,omitempty"`
}

func (resource Caaction) ToUnset() CaactionUnset {
	unset := CaactionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_caaction_payload struct {
	Unset CaactionUnset `json:"caaction"`
}

func (c *NitroClient) UnsetCaaction(unset CaactionUnset) error {
	payload := unset_caaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("caaction", "", qs, payload)
}

//      RENAME
//      TODO
