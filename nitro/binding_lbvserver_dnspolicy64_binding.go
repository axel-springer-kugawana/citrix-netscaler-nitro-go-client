package nitro

type LbvserverDnspolicy64Binding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverDnspolicy64BindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverDnspolicy64Binding(binding LbvserverDnspolicy64Binding) error {
	return nil
}

func (c *NitroClient) ListLbvserverDnspolicy64Binding() ([]LbvserverDnspolicy64Binding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) (*LbvserverDnspolicy64Binding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverDnspolicy64Binding(key LbvserverDnspolicy64BindingKey) error {
	return nil
}
