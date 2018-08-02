package nitro

type LbvserverCapolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverCapolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverCapolicyBinding(binding LbvserverCapolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverCapolicyBinding() ([]LbvserverCapolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverCapolicyBinding(key LbvserverCapolicyBindingKey) (*LbvserverCapolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverCapolicyBinding(key LbvserverCapolicyBindingKey) error {
	return nil
}
