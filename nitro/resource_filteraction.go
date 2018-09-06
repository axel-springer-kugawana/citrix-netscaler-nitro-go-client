package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Filteraction struct {
	Name        string `json:"name,omitempty"`
	Page        string `json:"page,omitempty"`
	Qual        string `json:"qual,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
}

type FilteractionKey struct {
	Name string
}

func (resource Filteraction) ToKey() FilteractionKey {
	key := FilteractionKey{
		resource.Name,
	}

	return key
}

func (key FilteractionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key FilteractionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key FilteractionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_filteraction_payload struct {
	Resource Filteraction `json:"filteraction"`
}

func (c *NitroClient) AddFilteraction(resource Filteraction) error {
	payload := add_filteraction_payload{
		resource,
	}

	return c.post("filteraction", "", nil, payload)
}

//      LIST

type list_filteraction_result struct {
	Results []Filteraction `json:"filteraction"`
}

func (c *NitroClient) ListFilteraction() ([]Filteraction, error) {
	results := list_filteraction_result{}

	if err := c.get("filteraction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_filteraction_result struct {
	Results []Filteraction `json:"filteraction"`
}

func (c *NitroClient) GetFilteraction(key FilteractionKey) (*Filteraction, error) {
	var results get_filteraction_result

	id, qs := key.to_id_args()

	if err := c.get("filteraction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one filteraction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("filteraction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_filteraction_result struct {
	Results []Count `json:"filteraction"`
}

func (c *NitroClient) CountFilteraction() (int, error) {
	var results count_filteraction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("filteraction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsFilteraction(key FilteractionKey) (bool, error) {
	var results count_filteraction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("filteraction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteFilteraction(key FilteractionKey) error {
	id, qs := key.to_id_args()

	return c.delete("filteraction", id, qs)
}

//      UPDATE

type FilteractionUpdate struct {
	Name        string `json:"name,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
	Respcode    int    `json:"respcode,string,omitempty"`
	Page        string `json:"page,omitempty"`
}

func (resource Filteraction) ToUpdate() FilteractionUpdate {
	update := FilteractionUpdate{
		resource.Name,
		resource.Servicename,
		resource.Value,
		resource.Respcode,
		resource.Page,
	}

	return update
}

type update_filteraction_payload struct {
	Update FilteractionUpdate `json:"filteraction"`
}

func (c *NitroClient) UpdateFilteraction(update FilteractionUpdate) error {
	payload := update_filteraction_payload{
		update,
	}

	return c.put("filteraction", "", nil, payload)
}

//      UNSET

type FilteractionUnset struct {
	Name        string `json:"name,omitempty"`
	Servicename bool   `json:"servicename,omitempty"`
	Value       bool   `json:"value,omitempty"`
	Respcode    bool   `json:"respcode,omitempty"`
	Page        bool   `json:"page,omitempty"`
}

func (resource Filteraction) ToUnset() FilteractionUnset {
	unset := FilteractionUnset{
		resource.Name,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_filteraction_payload struct {
	Unset FilteractionUnset `json:"filteraction"`
}

func (c *NitroClient) UnsetFilteraction(unset FilteractionUnset) error {
	payload := unset_filteraction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("filteraction", "", qs, payload)
}

//      RENAME
//      TODO
