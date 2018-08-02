package nitro

type LbvserverAuditnslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
}

type LbvserverAuditnslogpolicyBindingKey struct {
	Name       string
	Policyname string
	Bindpoint  string
}

func (c *NitroClient) AddLbvserverAuditnslogpolicyBinding(binding LbvserverAuditnslogpolicyBinding) error {
	return nil
}

func (c *NitroClient) ListLbvserverAuditnslogpolicyBinding() ([]LbvserverAuditnslogpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) GetLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) (*LbvserverAuditnslogpolicyBinding, error) {
	return nil, nil
}

func (c *NitroClient) DeleteLbvserverAuditnslogpolicyBinding(key LbvserverAuditnslogpolicyBindingKey) error {
	return nil
}
