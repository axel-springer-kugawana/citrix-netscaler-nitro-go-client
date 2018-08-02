package nitro

type ServiceScpolicyBinding struct {
	Name       string `json:"name,omitempty"`
	Policyname string `json:"policyname,omitempty"`
}

type ServiceScpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (c *NitroClient) AddServiceScpolicyBinding(binding ServiceScpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListServiceScpolicyBinding() ([]ServiceScpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetServiceScpolicyBinding(key ServiceScpolicyBindingKey) (*ServiceScpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteServiceScpolicyBinding(key ServiceScpolicyBindingKey) error {
	return nil
}
