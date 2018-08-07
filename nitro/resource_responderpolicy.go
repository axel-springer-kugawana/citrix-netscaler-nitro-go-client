package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Responderpolicy struct {
	Action        string `json:"action,omitempty"`
	Appflowaction string `json:"appflowaction,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Logaction     string `json:"logaction,omitempty"`
	Name          string `json:"name,omitempty"`
	Rule          string `json:"rule,omitempty"`
	Undefaction   string `json:"undefaction,omitempty"`
}

type ResponderpolicyKey struct {
	Name string
}

func (resource Responderpolicy) ToKey() ResponderpolicyKey {
	key := ResponderpolicyKey{
		resource.Name,
	}

	return key
}

func (key ResponderpolicyKey) to_id_args() (string, map[string]string) {
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

type add_responderpolicy_payload struct {
	Resource Responderpolicy `json:"responderpolicy"`
}

func (c *NitroClient) AddResponderpolicy(resource Responderpolicy) error {
	payload := add_responderpolicy_payload{
		resource,
	}

	return c.post("responderpolicy", "", nil, payload)
}

//      LIST

type list_responderpolicy_result struct {
	Results []Responderpolicy `json:"responderpolicy"`
}

func (c *NitroClient) ListResponderpolicy() ([]Responderpolicy, error) {
	results := list_responderpolicy_result{}

	if err := c.get("responderpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_responderpolicy_result struct {
	Results []Responderpolicy `json:"responderpolicy"`
}

func (c *NitroClient) GetResponderpolicy(key ResponderpolicyKey) (*Responderpolicy, error) {
	var results get_responderpolicy_result

	id, qs := key.to_id_args()

	if err := c.get("responderpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one responderpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("responderpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_responderpolicy_result struct {
	Results []Count `json:"responderpolicy"`
}

func (c *NitroClient) CountResponderpolicy() (int, error) {
	var results count_responderpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("responderpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsResponderpolicy(key ResponderpolicyKey) (bool, error) {
	var results count_responderpolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("responderpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteResponderpolicy(key ResponderpolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("responderpolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
