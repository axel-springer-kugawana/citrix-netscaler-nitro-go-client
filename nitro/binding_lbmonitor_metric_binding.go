package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbmonitorMetricBinding struct {
	Metric          string `json:"metric,omitempty"`
	Metricthreshold int    `json:"metricthreshold,string,omitempty"`
	Metricweight    int    `json:"metricweight,string,omitempty"`
	Monitorname     string `json:"monitorname,omitempty"`
}

type LbmonitorMetricBindingKey struct {
	Monitorname string
	Metric      string
}

type get_lbmonitor_metric_binding struct {
	Results []LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

type add_lbmonitor_metric_binding_payload struct {
	lbmonitor_metric_binding LbmonitorMetricBinding
}

func lbmonitor_metric_binding_key_to_id_args(key LbmonitorMetricBindingKey) (string, string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "monitorname:"+key.Monitorname)
	args = append(args, "metric:"+key.Metric)

	return "", strings.Join(args, ",")
}

// TODO : Exists
// TODO : Count

func (c *NitroClient) AddLbmonitorMetricBinding(binding LbmonitorMetricBinding) error {
	payload := add_lbmonitor_metric_binding_payload{
		binding,
	}

	return c.put("lbmonitor_metric_binding", "", "", "", payload)
}

func (c *NitroClient) ListLbmonitorMetricBinding(key LbmonitorMetricBindingKey) ([]LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding

	id, args := lbmonitor_metric_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_metric_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) BulkListLbmonitorMetricBinding() ([]LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding

	if err := c.get("lbmonitor_metric_binding", "", "", "", &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (*LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding

	id, args := lbmonitor_metric_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_metric_binding", id, "", args, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor_metric_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbmonitor_metric_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmonitorMetricBinding(key LbmonitorMetricBindingKey) error {
	id, args := lbmonitor_metric_binding_key_to_id_args(key)

	return c.delete("lbmonitor_metric_binding", id, "", args)
}
