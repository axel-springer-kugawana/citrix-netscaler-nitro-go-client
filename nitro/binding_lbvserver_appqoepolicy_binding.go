package nitro

type LbvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAppqoepolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAppqoepolicyBinding(binding LbvserverAppqoepolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAppqoepolicyBinding() ([]LbvserverAppqoepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) (*LbvserverAppqoepolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAppqoepolicyBinding(key LbvserverAppqoepolicyBindingKey) error {
	return nil
}
