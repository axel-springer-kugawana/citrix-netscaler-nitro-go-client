package nitro

type LbvserverFilterpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverFilterpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverFilterpolicyBinding(binding LbvserverFilterpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverFilterpolicyBinding() ([]LbvserverFilterpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverFilterpolicyBinding(key LbvserverFilterpolicyBindingKey) (*LbvserverFilterpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverFilterpolicyBinding(key LbvserverFilterpolicyBindingKey) error {
	return nil
}
