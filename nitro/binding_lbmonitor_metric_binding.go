package nitro

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

func (c *NitroClient) AddLbmonitorMetricBinding(binding LbmonitorMetricBinding) error {
	return nil
}

func (c *NitroClient) ListLbmonitorMetricBinding() ([]LbmonitorMetricBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbmonitorMetricBinding(key LbmonitorMetricBindingKey) (*LbmonitorMetricBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbmonitorMetricBinding(key LbmonitorMetricBindingKey) error {
	return nil
}
