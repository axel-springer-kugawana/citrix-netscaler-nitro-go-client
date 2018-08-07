package nitro

import (
	"fmt"
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

func (key DospolicyKey) to_id_args() (string, map[string]string) {
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
