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

type get_policydataset_value_binding struct {
	Results []PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

type add_policydataset_value_binding_payload struct {
	policydataset_value_binding PolicydatasetValueBinding
}

func policydataset_value_binding_key_to_id_args(key PolicydatasetValueBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "value:"+key.Value)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddPolicydatasetValueBinding(binding PolicydatasetValueBinding) error {
	payload := add_policydataset_value_binding_payload{
		binding,
	}

	return c.put("policydataset_value_binding", "", "", "", payload)
}

func (c *NitroClient) ListPolicydatasetValueBinding(key PolicydatasetValueBindingKey) ([]PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding

	id, args := policydataset_value_binding_key_to_id_args(key)

	if err := c.get("policydataset_value_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListPolicydatasetValueBinding() ([]PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding

	if err := c.get("policydataset_value_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicydatasetValueBinding(key PolicydatasetValueBindingKey) (*PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding

	id, args := policydataset_value_binding_key_to_id_args(key)

	if err := c.get("policydataset_value_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policydataset_value_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policydataset_value_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicydatasetValueBinding(key PolicydatasetValueBindingKey) error {
	id, args := policydataset_value_binding_key_to_id_args(key)

	return c.delete("policydataset_value_binding", id, "", args)
}
