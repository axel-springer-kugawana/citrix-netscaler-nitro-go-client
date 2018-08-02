package nitro

type PolicystringmapPatternBinding struct {
	Key   string `json:"key,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PolicystringmapPatternBindingKey struct {
	Name string
	Key  string
}

func (c *NitroClient) AddPolicystringmapPatternBinding(binding PolicystringmapPatternBinding) error {
	return nil
}

func (c *NitroClient) ListPolicystringmapPatternBinding() ([]PolicystringmapPatternBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetPolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) (*PolicystringmapPatternBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeletePolicystringmapPatternBinding(key PolicystringmapPatternBindingKey) error {
	return nil
}
