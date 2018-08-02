package nitro

type LbvserverTransformpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverTransformpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverTransformpolicyBinding(binding LbvserverTransformpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverTransformpolicyBinding() ([]LbvserverTransformpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverTransformpolicyBinding(key LbvserverTransformpolicyBindingKey) (*LbvserverTransformpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverTransformpolicyBinding(key LbvserverTransformpolicyBindingKey) error {
	return nil
}
