package nitro

type PolicydatasetValueBinding struct {
	Index int    `json:"index,string,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PolicydatasetValueBindingKey struct {
	Name  string
	Value string
}

func (c *NitroClient) AddPolicydatasetValueBinding(binding PolicydatasetValueBinding) error {
	return nil
}

func (c *NitroClient) ListPolicydatasetValueBinding() ([]PolicydatasetValueBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetPolicydatasetValueBinding(key PolicydatasetValueBindingKey) (*PolicydatasetValueBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeletePolicydatasetValueBinding(key PolicydatasetValueBindingKey) error {
	return nil
}
