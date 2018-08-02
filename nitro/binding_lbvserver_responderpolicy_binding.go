package nitro

type LbvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverResponderpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverResponderpolicyBinding(binding LbvserverResponderpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverResponderpolicyBinding() ([]LbvserverResponderpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) (*LbvserverResponderpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverResponderpolicyBinding(key LbvserverResponderpolicyBindingKey) error {
	return nil
}
