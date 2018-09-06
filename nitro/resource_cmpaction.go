package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmpaction struct {
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Cmptype         string `json:"cmptype,omitempty"`
	Deltatype       string `json:"deltatype,omitempty"`
	Name            string `json:"name,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

type CmpactionKey struct {
	Name string
}

func (resource Cmpaction) ToKey() CmpactionKey {
	key := CmpactionKey{
		resource.Name,
	}

	return key
}

func (key CmpactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CmpactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CmpactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cmpaction_payload struct {
	Resource Cmpaction `json:"cmpaction"`
}

func (c *NitroClient) AddCmpaction(resource Cmpaction) error {
	payload := add_cmpaction_payload{
		resource,
	}

	return c.post("cmpaction", "", nil, payload)
}

//      LIST

type list_cmpaction_result struct {
	Results []Cmpaction `json:"cmpaction"`
}

func (c *NitroClient) ListCmpaction() ([]Cmpaction, error) {
	results := list_cmpaction_result{}

	if err := c.get("cmpaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cmpaction_result struct {
	Results []Cmpaction `json:"cmpaction"`
}

func (c *NitroClient) GetCmpaction(key CmpactionKey) (*Cmpaction, error) {
	var results get_cmpaction_result

	id, qs := key.to_id_args()

	if err := c.get("cmpaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cmpaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cmpaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cmpaction_result struct {
	Results []Count `json:"cmpaction"`
}

func (c *NitroClient) CountCmpaction() (int, error) {
	var results count_cmpaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cmpaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCmpaction(key CmpactionKey) (bool, error) {
	var results count_cmpaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cmpaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCmpaction(key CmpactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("cmpaction", id, qs)
}

//      UPDATE

type CmpactionUpdate struct {
	Name            string `json:"name,omitempty"`
	Cmptype         string `json:"cmptype,omitempty"`
	Addvaryheader   string `json:"addvaryheader,omitempty"`
	Varyheadervalue string `json:"varyheadervalue,omitempty"`
}

func (resource Cmpaction) ToUpdate() CmpactionUpdate {
	update := CmpactionUpdate{
		resource.Name,
		resource.Cmptype,
		resource.Addvaryheader,
		resource.Varyheadervalue,
	}

	return update
}

type update_cmpaction_payload struct {
	Update CmpactionUpdate `json:"cmpaction"`
}

func (c *NitroClient) UpdateCmpaction(update CmpactionUpdate) error {
	payload := update_cmpaction_payload{
		update,
	}

	return c.put("cmpaction", "", nil, payload)
}

//      UNSET

type CmpactionUnset struct {
	Name            string `json:"name,omitempty"`
	Cmptype         bool   `json:"cmptype,omitempty"`
	Addvaryheader   bool   `json:"addvaryheader,omitempty"`
	Varyheadervalue bool   `json:"varyheadervalue,omitempty"`
}

func (resource Cmpaction) ToUnset() CmpactionUnset {
	unset := CmpactionUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_cmpaction_payload struct {
	Unset CmpactionUnset `json:"cmpaction"`
}

func (c *NitroClient) UnsetCmpaction(unset CmpactionUnset) error {
	payload := unset_cmpaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("cmpaction", "", qs, payload)
}

//      RENAME
//      TODO
