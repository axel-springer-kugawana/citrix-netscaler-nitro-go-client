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

func (resource LbmonitorMetricBinding) ToKey() LbmonitorMetricBindingKey {
	key := LbmonitorMetricBindingKey{
		resource.Monitorname,
		resource.Metric,
	}

	return key
}

func (key LbmonitorMetricBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Monitorname
	args = append(args, "metric:"+key.Metric)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbmonitorMetricBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbmonitorMetricBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbmonitor_metric_binding_payload struct {
	Resource LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

func (c *NitroClient) AddLbmonitorMetricBinding(resource LbmonitorMetricBinding) error {
	payload := add_lbmonitor_metric_binding_payload{
		resource,
	}

	return c.put("lbmonitor_metric_binding", "", nil, payload)
}

//      LIST

type list_lbmonitor_metric_binding_result struct {
	Results []LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

func (c *NitroClient) ListLbmonitorMetricBinding() ([]LbmonitorMetricBinding, error) {
	results := list_lbmonitor_metric_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbmonitor_metric_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_lbmonitor_metric_binding_result struct {
	Results []Count `json:"lbmonitor_metric_binding"`
}

func (c *NitroClient) CountLbmonitorMetricBinding() (int, error) {
	results := count_lbmonitor_metric_binding_result{}

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

//      READ

type get_lbmonitor_metric_binding_result struct {
	Results []LbmonitorMetricBinding `json:"lbmonitor_metric_binding"`
}

func (c *NitroClient) GetLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (*LbmonitorMetricBinding, error) {
	var results get_lbmonitor_metric_binding_result

	id, qs := key.to_id_filter()

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

//      EXISTS

func (c *NitroClient) ExistsLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (bool, error) {
	var results count_lbmonitor_metric_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("lbmonitor_metric_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbmonitorMetricBinding(key LbmonitorMetricBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbmonitor_metric_binding", id, qs)
}
