package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Policydataset struct {
	Comment   string `json:"comment,omitempty"`
	Indextype string `json:"indextype,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
}

type PolicydatasetKey struct {
	Name string
}

func (resource Policydataset) ToKey() PolicydatasetKey {
	key := PolicydatasetKey{
		resource.Name,
	}

	return key
}

func (key PolicydatasetKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key PolicydatasetKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PolicydatasetKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_policydataset_payload struct {
	Resource Policydataset `json:"policydataset"`
}

func (c *NitroClient) AddPolicydataset(resource Policydataset) error {
	payload := add_policydataset_payload{
		resource,
	}

	return c.post("policydataset", "", nil, payload)
}

//      LIST

type list_policydataset_result struct {
	Results []Policydataset `json:"policydataset"`
}

func (c *NitroClient) ListPolicydataset() ([]Policydataset, error) {
	results := list_policydataset_result{}

	if err := c.get("policydataset", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_policydataset_result struct {
	Results []Policydataset `json:"policydataset"`
}

func (c *NitroClient) GetPolicydataset(key PolicydatasetKey) (*Policydataset, error) {
	var results get_policydataset_result

	id, qs := key.to_id_args()

	if err := c.get("policydataset", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policydataset element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policydataset element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_policydataset_result struct {
	Results []Count `json:"policydataset"`
}

func (c *NitroClient) CountPolicydataset() (int, error) {
	var results count_policydataset_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policydataset", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicydataset(key PolicydatasetKey) (bool, error) {
	var results count_policydataset_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("policydataset", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicydataset(key PolicydatasetKey) error {
	id, qs := key.to_id_args()

	return c.delete("policydataset", id, qs)
}

//      RENAME
//      TODO
