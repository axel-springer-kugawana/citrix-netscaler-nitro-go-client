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

func (key FilteractionKey) to_id_args() (string, map[string]string) {
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
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
