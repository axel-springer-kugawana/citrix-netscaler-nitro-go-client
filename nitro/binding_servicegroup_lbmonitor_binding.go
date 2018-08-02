package nitro

type ServicegroupLbmonitorBinding struct {
	Monitor_name     string `json:"monitor_name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Weight           int    `json:"weight,string,omitempty"`
}

type ServicegroupLbmonitorBindingKey struct {
	Servicegroupname string
	Monitor_name     string
}

func (c *NitroClient) AddServicegroupLbmonitorBinding(binding ServicegroupLbmonitorBinding) error {
	return nil
}

func (c *NitroClient) ListServicegroupLbmonitorBinding() ([]ServicegroupLbmonitorBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) (*ServicegroupLbmonitorBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteServicegroupLbmonitorBinding(key ServicegroupLbmonitorBindingKey) error {
	return nil
}
