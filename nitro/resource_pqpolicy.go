package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Pqpolicy struct {
	Policyname string `json:"policyname,omitempty"`
	Polqdepth  int    `json:"polqdepth,string,omitempty"`
	Priority   int    `json:"priority,string,omitempty"`
	Qdepth     int    `json:"qdepth,string,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Weight     int    `json:"weight,string,omitempty"`
}

type PqpolicyKey struct {
	Policyname string
}

func (resource Pqpolicy) ToKey() PqpolicyKey {
	key := PqpolicyKey{
		resource.Policyname,
	}

	return key
}

func (key PqpolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Policyname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key PqpolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PqpolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_pqpolicy_payload struct {
	Resource Pqpolicy `json:"pqpolicy"`
}

func (c *NitroClient) AddPqpolicy(resource Pqpolicy) error {
	payload := add_pqpolicy_payload{
		resource,
	}

	return c.post("pqpolicy", "", nil, payload)
}

//      LIST

type list_pqpolicy_result struct {
	Results []Pqpolicy `json:"pqpolicy"`
}

func (c *NitroClient) ListPqpolicy() ([]Pqpolicy, error) {
	results := list_pqpolicy_result{}

	if err := c.get("pqpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_pqpolicy_result struct {
	Results []Pqpolicy `json:"pqpolicy"`
}

func (c *NitroClient) GetPqpolicy(key PqpolicyKey) (*Pqpolicy, error) {
	var results get_pqpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("pqpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one pqpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("pqpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_pqpolicy_result struct {
	Results []Count `json:"pqpolicy"`
}

func (c *NitroClient) CountPqpolicy() (int, error) {
	var results count_pqpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("pqpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsPqpolicy(key PqpolicyKey) (bool, error) {
	var results count_pqpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("pqpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePqpolicy(key PqpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("pqpolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
