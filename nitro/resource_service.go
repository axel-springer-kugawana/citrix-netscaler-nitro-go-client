package nitro

type ServiceResource struct {
	Name               string `json:"name"`
	Accessdown         string `json:"accessdown,omitempty"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cachetype          string `json:"cachetype,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Cleartextport      int    `json:"cleartextport,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Customserverid     string `json:"customserverid,omitempty"`
	Dnsprofilename     string `json:"dnsprofilename,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Hashid             int    `json:"hashid,string,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Port               int    `json:"port,omitempty"`
	Processlocal       string `json:"processlocal,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Serverid           int    `json:"serverid,string,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sp                 string `json:"sp,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Usip               string `json:"usip,omitempty"`
}

type ServiceKey struct {
	Name string
}

type service_update struct {
	Name               string `json:"name"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Usip               string `json:"usip,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Sp                 string `json:"sp,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Customserverid     string `json:"customserverid,omitempty"`
	Serverid           int    `json:"serverid,string,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Accessdown         string `json:"accessdown,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Hashid             int    `json:"hashid,string,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Processlocal       string `json:"processlocal,omitempty"`
	Dnsprofilename     string `json:"dnsprofilename,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
}

type service_payload struct {
	service interface{}
}

func service_key_to_args(key ServiceKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteService(key ServiceKey) error {
	return c.deleteResourceWithArgs("service", key.Name, service_key_to_args(key))
}

func (c *NitroClient) GetService(key ServiceKey) (*ServiceResource, error) {
	var results struct {
		Service []ServiceResource
	}

	if err := c.getResourceWithArgs("service", key.Name, service_key_to_args(key), &results); err != nil || len(results.Service) != 1 {
		return nil, err
	}

	return &results.Service[0], nil
}

func (c *NitroClient) ListService() ([]ServiceResource, error) {
	var results struct {
		Service []ServiceResource
	}

	if err := c.listResources("service", &results); err != nil {
		return nil, err
	}

	return results.Service, nil
}

func (c *NitroClient) AddService(resource ServiceResource) error {
	return c.addResource("service", resource)
}

func (c *NitroClient) RenameService(name string, newName string) error {
	return c.renameResource("service", "name", name, newName)
}

func (c *NitroClient) UnsetService(name string, fields ...string) error {
	return c.unsetResource("service", "name", name, fields)
}

func (c *NitroClient) UpdateService(resource ServiceResource) error {
	update := service_update{
		resource.Name,
		resource.Maxclient,
		resource.Maxreq,
		resource.Cacheable,
		resource.Cip,
		resource.Cipheader,
		resource.Usip,
		resource.Pathmonitor,
		resource.Pathmonitorindv,
		resource.Useproxyport,
		resource.Sc,
		resource.Sp,
		resource.Rtspsessionidremap,
		resource.Healthmonitor,
		resource.Clttimeout,
		resource.Svrtimeout,
		resource.Customserverid,
		resource.Serverid,
		resource.Cka,
		resource.Tcpb,
		resource.Cmp,
		resource.Maxbandwidth,
		resource.Accessdown,
		resource.Monthreshold,
		resource.Downstateflush,
		resource.Tcpprofilename,
		resource.Httpprofilename,
		resource.Hashid,
		resource.Comment,
		resource.Appflowlog,
		resource.Netprofile,
		resource.Processlocal,
		resource.Dnsprofilename,
		resource.Monconnectionclose,
	}

	return c.Put("service", update)
}
