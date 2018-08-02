package nitro

type LbvserverAuditsyslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuditsyslogpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAuditsyslogpolicyBinding(binding LbvserverAuditsyslogpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAuditsyslogpolicyBinding() ([]LbvserverAuditsyslogpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) (*LbvserverAuditsyslogpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAuditsyslogpolicyBinding(key LbvserverAuditsyslogpolicyBindingKey) error {
	return nil
}
