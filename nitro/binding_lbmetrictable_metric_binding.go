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

type get_lbmetrictable_metric_binding struct {
	Results []LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

type add_lbmetrictable_metric_binding_payload struct {
	lbmetrictable_metric_binding LbmetrictableMetricBinding
}

func lbmetrictable_metric_binding_key_to_id_args(key LbmetrictableMetricBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "metrictable:"+key.Metrictable)
	args = append(args, "metric:"+key.Metric)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbmetrictableMetricBinding(binding LbmetrictableMetricBinding) error {
	payload := add_lbmetrictable_metric_binding_payload{
		binding,
	}

	return c.put("lbmetrictable_metric_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) ([]LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding

	id, args := lbmetrictable_metric_binding_key_to_id_args(key)

	if err := c.get("lbmetrictable_metric_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbmetrictableMetricBinding() ([]LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding

	if err := c.get("lbmetrictable_metric_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) (*LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding

	id, args := lbmetrictable_metric_binding_key_to_id_args(key)

	if err := c.get("lbmetrictable_metric_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmetrictable_metric_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbmetrictable_metric_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) error {
	id, args := lbmetrictable_metric_binding_key_to_id_args(key)

	return c.delete("lbmetrictable_metric_binding", id, "", args)
}
