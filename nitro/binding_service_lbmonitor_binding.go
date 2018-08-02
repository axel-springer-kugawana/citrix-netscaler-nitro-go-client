package nitro

type ServiceLbmonitorBinding struct {
	Monitor_name string `json:"monitor_name,omitempty"`
	Monstate     string `json:"monstate,omitempty"`
	Name         string `json:"name,omitempty"`
	Passive      bool   `json:"passive,omitempty"`
	Weight       int    `json:"weight,string,omitempty"`
}

type ServiceLbmonitorBindingKey struct {
	Name string
}

func (c *NitroClient) AddServiceLbmonitorBinding(binding ServiceLbmonitorBinding) error {
	return nil
}

func (c *NitroClient) ListServiceLbmonitorBinding() ([]ServiceLbmonitorBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) (*ServiceLbmonitorBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteServiceLbmonitorBinding(key ServiceLbmonitorBindingKey) error {
	return nil
}
