package nitro

type LbmetrictableMetricBinding struct {
	Metric      string `json:"metric,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
	Snmpoid     string `json:"snmpoid,omitempty"`
}

type LbmetrictableMetricBindingKey struct {
	Metrictable string
	Metric      string
}

func (c *NitroClient) AddLbmetrictableMetricBinding(binding LbmetrictableMetricBinding) error {
	return nil
}

func (c *NitroClient) ListLbmetrictableMetricBinding() ([]LbmetrictableMetricBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) (*LbmetrictableMetricBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) error {
	return nil
}
