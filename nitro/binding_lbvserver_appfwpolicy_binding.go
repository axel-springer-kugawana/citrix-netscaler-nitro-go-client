package nitro

type LbvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppfwpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAppfwpolicyBinding(binding LbvserverAppfwpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAppfwpolicyBinding() ([]LbvserverAppfwpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) (*LbvserverAppfwpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAppfwpolicyBinding(key LbvserverAppfwpolicyBindingKey) error {
	return nil
}
