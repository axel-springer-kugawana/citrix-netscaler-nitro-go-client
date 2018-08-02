package nitro

type LbvserverServicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type LbvserverServicegroupBindingKey struct {
	Name             string
	Servicegroupname string
}

func (c *NitroClient) AddLbvserverServicegroupBinding(binding LbvserverServicegroupBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverServicegroupBinding() ([]LbvserverServicegroupBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) (*LbvserverServicegroupBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverServicegroupBinding(key LbvserverServicegroupBindingKey) error {
	return nil
}
