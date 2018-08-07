package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbmetrictableMetricBinding struct {
	Metric      string `json:"metric,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
	Snmpoid     string `json:"snmpoid,omitempty"`
}

type LbmetrictableMetricBindingKey struct {
	Metrictable string
	Metric      string
}

func (resource LbmetrictableMetricBinding) ToKey() LbmetrictableMetricBindingKey {
	key := LbmetrictableMetricBindingKey{
		resource.Metrictable,
		resource.Metric,
	}

	return key
}

func (key LbmetrictableMetricBindingKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Metrictable
	args = append(args, "metric:"+key.Metric)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
}

//      CREATE

type add_lbmetrictable_metric_binding_payload struct {
	Resource LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

func (c *NitroClient) AddLbmetrictableMetricBinding(resource LbmetrictableMetricBinding) error {
	payload := add_lbmetrictable_metric_binding_payload{
		resource,
	}

	return c.put("lbmetrictable_metric_binding", "", nil, payload)
}

//      LIST

type list_lbmetrictable_metric_binding_result struct {
	Results []LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

func (c *NitroClient) ListLbmetrictableMetricBinding() ([]LbmetrictableMetricBinding, error) {
	results := list_lbmetrictable_metric_binding_result{}

	if err := c.get("lbmetrictable_metric_binding", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbmetrictable_metric_binding_result struct {
	Results []LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

func (c *NitroClient) GetLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) (*LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding_result

	id, qs := key.to_id_args()

	if err := c.get("lbmetrictable_metric_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmetrictable_metric_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbmetrictable_metric_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

type count_lbmetrictable_metric_binding_result struct {
	Results []Count `json:"lbmetrictable_metric_binding"`
}

func (c *NitroClient) ExistsLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) (bool, error) {
	var results count_lbmetrictable_metric_binding_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbmetrictable_metric_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbmetrictable_metric_binding", id, qs)
}
