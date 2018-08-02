package nitro

type LbvserverSpilloverpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverSpilloverpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverSpilloverpolicyBinding(binding LbvserverSpilloverpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverSpilloverpolicyBinding() ([]LbvserverSpilloverpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) (*LbvserverSpilloverpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverSpilloverpolicyBinding(key LbvserverSpilloverpolicyBindingKey) error {
	return nil
}
