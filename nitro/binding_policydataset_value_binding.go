package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type PolicydatasetValueBinding struct {
	Index int    `json:"index,string,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PolicydatasetValueBindingKey struct {
	Name  string
	Value string
}

func (resource PolicydatasetValueBinding) ToKey() PolicydatasetValueBindingKey {
	key := PolicydatasetValueBindingKey{
		resource.Name,
		resource.Value,
	}

	return key
}

func (key PolicydatasetValueBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "value:"+key.Value)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key PolicydatasetValueBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key PolicydatasetValueBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_policydataset_value_binding_payload struct {
	Resource PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

func (c *NitroClient) AddPolicydatasetValueBinding(resource PolicydatasetValueBinding) error {
	payload := add_policydataset_value_binding_payload{
		resource,
	}

	return c.put("policydataset_value_binding", "", nil, payload)
}

//      LIST

type list_policydataset_value_binding_result struct {
	Results []PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

func (c *NitroClient) ListPolicydatasetValueBinding() ([]PolicydatasetValueBinding, error) {
	results := list_policydataset_value_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("policydataset_value_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_policydataset_value_binding_result struct {
	Results []Count `json:"policydataset_value_binding"`
}

func (c *NitroClient) CountPolicydatasetValueBinding() (int, error) {
	results := count_policydataset_value_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("policydataset_value_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_policydataset_value_binding_result struct {
	Results []PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

func (c *NitroClient) GetPolicydatasetValueBinding(key PolicydatasetValueBindingKey) (*PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("policydataset_value_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policydataset_value_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("policydataset_value_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsPolicydatasetValueBinding(key PolicydatasetValueBindingKey) (bool, error) {
	var results count_policydataset_value_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("policydataset_value_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeletePolicydatasetValueBinding(key PolicydatasetValueBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("policydataset_value_binding", id, qs)
}
