package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Filterpolicy struct {
	Name      string `json:"name,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}

type FilterpolicyKey struct {
	Name string
}

func (resource Filterpolicy) ToKey() FilterpolicyKey {
	key := FilterpolicyKey{
		resource.Name,
	}

	return key
}

func (key FilterpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key FilterpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key FilterpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_filterpolicy_payload struct {
	Resource Filterpolicy `json:"filterpolicy"`
}

func (c *NitroClient) AddFilterpolicy(resource Filterpolicy) error {
	payload := add_filterpolicy_payload{
		resource,
	}

	return c.post("filterpolicy", "", nil, payload)
}

//      LIST

type list_filterpolicy_result struct {
	Results []Filterpolicy `json:"filterpolicy"`
}

func (c *NitroClient) ListFilterpolicy() ([]Filterpolicy, error) {
	results := list_filterpolicy_result{}

	if err := c.get("filterpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_filterpolicy_result struct {
	Results []Filterpolicy `json:"filterpolicy"`
}

func (c *NitroClient) GetFilterpolicy(key FilterpolicyKey) (*Filterpolicy, error) {
	var results get_filterpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("filterpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one filterpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("filterpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_filterpolicy_result struct {
	Results []Count `json:"filterpolicy"`
}

func (c *NitroClient) CountFilterpolicy() (int, error) {
	var results count_filterpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("filterpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsFilterpolicy(key FilterpolicyKey) (bool, error) {
	var results count_filterpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("filterpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteFilterpolicy(key FilterpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("filterpolicy", id, qs)
}

//      UPDATE

type FilterpolicyUpdate struct {
	Name      string `json:"name,omitempty"`
	Rule      string `json:"rule,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
}

func (resource Filterpolicy) ToUpdate() FilterpolicyUpdate {
	update := FilterpolicyUpdate{
		resource.Name,
		resource.Rule,
		resource.Reqaction,
		resource.Resaction,
	}

	return update
}

type update_filterpolicy_payload struct {
	Update FilterpolicyUpdate `json:"filterpolicy"`
}

func (c *NitroClient) UpdateFilterpolicy(update FilterpolicyUpdate) error {
	payload := update_filterpolicy_payload{
		update,
	}

	return c.put("filterpolicy", "", nil, payload)
}

//      UNSET

type FilterpolicyUnset struct {
	Name      string `json:"name,omitempty"`
	Rule      bool   `json:"rule,omitempty"`
	Reqaction bool   `json:"reqaction,omitempty"`
	Resaction bool   `json:"resaction,omitempty"`
}

func (resource Filterpolicy) ToUnset() FilterpolicyUnset {
	unset := FilterpolicyUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_filterpolicy_payload struct {
	Unset FilterpolicyUnset `json:"filterpolicy"`
}

func (c *NitroClient) UnsetFilterpolicy(unset FilterpolicyUnset) error {
	payload := unset_filterpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("filterpolicy", "", qs, payload)
}

//      RENAME
//      TODO
