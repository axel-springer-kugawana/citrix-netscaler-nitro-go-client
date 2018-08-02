package nitro

type ServiceDospolicyBinding struct {
	Name       string `json:"name,omitempty"`
	Policyname string `json:"policyname,omitempty"`
}

type ServiceDospolicyBindingKey struct {
	Name       string
	Policyname string
}

func (c *NitroClient) AddServiceDospolicyBinding(binding ServiceDospolicyBinding) error {
	return nil
}

func (c *NitroClient) ListServiceDospolicyBinding() ([]ServiceDospolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetServiceDospolicyBinding(key ServiceDospolicyBindingKey) (*ServiceDospolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteServiceDospolicyBinding(key ServiceDospolicyBindingKey) error {
	return nil
}
