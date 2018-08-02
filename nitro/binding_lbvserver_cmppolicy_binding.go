package nitro

type LbvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCmppolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverCmppolicyBinding(binding LbvserverCmppolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverCmppolicyBinding() ([]LbvserverCmppolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) (*LbvserverCmppolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverCmppolicyBinding(key LbvserverCmppolicyBindingKey) error {
	return nil
}
