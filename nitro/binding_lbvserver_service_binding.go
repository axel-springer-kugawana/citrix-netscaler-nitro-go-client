package nitro

type LbvserverServiceBinding struct {
	Name        string `json:"name,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight      int    `json:"weight,string,omitempty"`
}

type LbvserverServiceBindingKey struct {
	Name        string
	Servicename string
}

func (c *NitroClient) AddLbvserverServiceBinding(binding LbvserverServiceBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverServiceBinding() ([]LbvserverServiceBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverServiceBinding(key LbvserverServiceBindingKey) (*LbvserverServiceBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverServiceBinding(key LbvserverServiceBindingKey) error {
	return nil
}
