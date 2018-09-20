package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Dospolicy struct {
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
	Name          string `json:"name,omitempty"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
}

type DospolicyKey struct {
	Name string
}

func (resource Dospolicy) ToKey() DospolicyKey {
	key := DospolicyKey{
		resource.Name,
	}

	return key
}

func (key DospolicyKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key DospolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key DospolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_dospolicy_payload struct {
	Resource Dospolicy `json:"dospolicy"`
}

func (c *NitroClient) AddDospolicy(resource Dospolicy) error {
	payload := add_dospolicy_payload{
		resource,
	}

	return c.post("dospolicy", "", nil, payload)
}

//      LIST

type list_dospolicy_result struct {
	Results []Dospolicy `json:"dospolicy"`
}

func (c *NitroClient) ListDospolicy() ([]Dospolicy, error) {
	results := list_dospolicy_result{}

	if err := c.get("dospolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_dospolicy_result struct {
	Results []Dospolicy `json:"dospolicy"`
}

func (c *NitroClient) GetDospolicy(key DospolicyKey) (*Dospolicy, error) {
	var results get_dospolicy_result

	id, qs := key.to_id_args()

	if err := c.get("dospolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dospolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dospolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_dospolicy_result struct {
	Results []Count `json:"dospolicy"`
}

func (c *NitroClient) CountDospolicy() (int, error) {
	var results count_dospolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dospolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsDospolicy(key DospolicyKey) (bool, error) {
	var results count_dospolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("dospolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteDospolicy(key DospolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("dospolicy", id, qs)
}

//      UPDATE

type DospolicyUpdate struct {
	Name          string `json:"name,omitempty"`
	Qdepth        int    `json:"qdepth,string,omitempty"`
	Cltdetectrate int    `json:"cltdetectrate,string,omitempty"`
}

func (resource Dospolicy) ToUpdate() DospolicyUpdate {
	update := DospolicyUpdate{
		resource.Name,
		resource.Qdepth,
		resource.Cltdetectrate,
	}

	return update
}

type update_dospolicy_payload struct {
	Update DospolicyUpdate `json:"dospolicy"`
}

func (c *NitroClient) UpdateDospolicy(update DospolicyUpdate) error {
	payload := update_dospolicy_payload{
		update,
	}

	return c.put("dospolicy", "", nil, payload)
}

//      UNSET

type DospolicyUnset struct {
	Name          string `json:"name,omitempty"`
	Qdepth        bool   `json:"qdepth,omitempty"`
	Cltdetectrate bool   `json:"cltdetectrate,omitempty"`
}

func (resource Dospolicy) ToUnset() DospolicyUnset {
	unset := DospolicyUnset{
		resource.Name,
		false,
		false,
	}

	return unset
}

type unset_dospolicy_payload struct {
	Unset DospolicyUnset `json:"dospolicy"`
}

func (c *NitroClient) UnsetDospolicy(unset DospolicyUnset) error {
	payload := unset_dospolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("dospolicy", "", qs, payload)
}

//      RENAME
//      TODO
