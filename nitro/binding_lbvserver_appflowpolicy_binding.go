package nitro

type LbvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppflowpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAppflowpolicyBinding(binding LbvserverAppflowpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAppflowpolicyBinding() ([]LbvserverAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) (*LbvserverAppflowpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAppflowpolicyBinding(key LbvserverAppflowpolicyBindingKey) error {
	return nil
}
