package nitro

type ServicegroupServicegroupmemberBinding struct {
	Port             int    `json:"port,omitempty"`
	Servername       string `json:"servername,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupServicegroupmemberBindingKey struct {
	Servicegroupname string
	Servername       string
	Port             int
}

func (c *NitroClient) AddServicegroupServicegroupmemberBinding(binding ServicegroupServicegroupmemberBinding) error {
	return nil
}

func (c *NitroClient) ListServicegroupServicegroupmemberBinding() ([]ServicegroupServicegroupmemberBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) (*ServicegroupServicegroupmemberBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteServicegroupServicegroupmemberBinding(key ServicegroupServicegroupmemberBindingKey) error {
	return nil
}
