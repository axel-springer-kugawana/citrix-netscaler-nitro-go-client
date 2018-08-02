package nitro

type PolicypatsetPatternBinding struct {
	Charset string `json:"charset,omitempty"`
	Index   int    `json:"index,string,omitempty"`
	Name    string `json:"name,omitempty"`
	String  string `json:"string,omitempty"`
}

type PolicypatsetPatternBindingKey struct {
	Name   string
	String string
}

func (c *NitroClient) AddPolicypatsetPatternBinding(binding PolicypatsetPatternBinding) error {
	return nil
}

func (c *NitroClient) ListPolicypatsetPatternBinding() ([]PolicypatsetPatternBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetPolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) (*PolicypatsetPatternBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeletePolicypatsetPatternBinding(key PolicypatsetPatternBindingKey) error {
	return nil
}
