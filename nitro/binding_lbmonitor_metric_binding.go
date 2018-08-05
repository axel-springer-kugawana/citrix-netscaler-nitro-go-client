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

type add_lbmonitor_metric_binding_payload struct {
	Resource LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

type get_lbmonitor_metric_binding_result struct {
	Results []LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

type count_lbmonitor_metric_binding_result struct {
	Results []Count `json:"lbmonitor_metric_binding"`
}

func lbmonitor_metric_binding_key_to_id_qs(key LbmonitorMetricBindingKey, arg string) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "monitorname:"+key.Monitorname)
	args = append(args, "metric:"+key.Metric)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[arg] = strings.Join(args, ",")
	}

	return "", qs
}

func lbmonitor_metric_binding_key_to_id_args(key LbmonitorMetricBindingKey) (string, map[string]string) {
	return lbmonitor_metric_binding_key_to_id_qs(key, "args")
}

func lbmonitor_metric_binding_key_to_id_filter(key LbmonitorMetricBindingKey) (string, map[string]string) {
	return lbmonitor_metric_binding_key_to_id_qs(key, "filter")
}

func (c *NitroClient) AddLbmonitorMetricBinding(binding LbmonitorMetricBinding) error {
	payload := add_lbmonitor_metric_binding_payload{
		binding,
	}

	return c.put("lbmonitor_metric_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbmonitorMetricBinding() (int, error) {
	var results count_lbmonitor_metric_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbmonitor_metric_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbmonitorMetricBinding(id string) (int, error) {
	var results count_lbmonitor_metric_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmonitor_metric_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (bool, error) {
	var results count_lbmonitor_metric_binding_result

	id, qs := lbmonitor_metric_binding_key_to_id_filter(key)

	qs["count"] = "yes"

	if err := c.get("lbmonitor_metric_binding", id, qs, &results); err != nil {
		return false, err
	} else {
		if len(results.Results) > 1 {
			return false, fmt.Errorf("More than one lbmonitor_metric_binding element found")
		}

		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) BulkListLbmonitorMetricBinding() ([]LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbmonitor_metric_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbmonitorMetricBinding(id string) ([]LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding_result

	if err := c.get("lbmonitor_metric_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (*LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding_result

	id, qs := lbmonitor_metric_binding_key_to_id_args(key)

	if err := c.get("lbmonitor_metric_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor_metric_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbmonitor_metric_binding element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmonitorMetricBinding(key LbmonitorMetricBindingKey) error {
	id, qs := lbmonitor_metric_binding_key_to_id_args(key)

	return c.delete("lbmonitor_metric_binding", id, qs)
}
