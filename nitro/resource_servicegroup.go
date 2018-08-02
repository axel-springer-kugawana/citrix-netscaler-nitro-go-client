package nitro

type Servicegroup struct {
	Servicegroupname   string `json:"servicegroupname"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Autoscale          string `json:"autoscale,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cachetype          string `json:"cachetype,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Memberport         int    `json:"memberport,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sp                 string `json:"sp,omitempty"`
	State              string `json:"state,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Usip               string `json:"usip,omitempty"`
}

type ServicegroupKey struct {
	Servicegroupname string
}

type servicegroup_update struct {
	Servicegroupname   string `json:"servicegroupname"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
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
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
}

type servicegroup_payload struct {
	servicegroup interface{}
}

func servicegroup_key_to_args(key ServicegroupKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteServicegroup(key ServicegroupKey) error {
	return c.deleteResourceWithArgs("servicegroup", key.Servicegroupname, servicegroup_key_to_args(key))
}

func (c *NitroClient) GetServicegroup(key ServicegroupKey) (*Servicegroup, error) {
	var results struct {
		Servicegroup []Servicegroup
	}

	if err := c.getResourceWithArgs("servicegroup", key.Servicegroupname, servicegroup_key_to_args(key), &results); err != nil || len(results.Servicegroup) != 1 {
		return nil, err
	}

	return &results.Servicegroup[0], nil
}

func (c *NitroClient) ListServicegroup() ([]Servicegroup, error) {
	var results struct {
		Servicegroup []Servicegroup
	}

	if err := c.listResources("servicegroup", &results); err != nil {
		return nil, err
	}

	return results.Servicegroup, nil
}

func (c *NitroClient) AddServicegroup(resource Servicegroup) error {
	return c.addResource("servicegroup", resource)
}

func (c *NitroClient) RenameServicegroup(servicegroupname string, newName string) error {
	return c.renameResource("servicegroup", "servicegroupname", servicegroupname, newName)
}

func (c *NitroClient) UnsetServicegroup(servicegroupname string, fields ...string) error {
	return c.unsetResource("servicegroup", "servicegroupname", servicegroupname, fields)
}

func (c *NitroClient) UpdateServicegroup(resource Servicegroup) error {
	update := servicegroup_update{
		resource.Servicegroupname,
		resource.Maxclient,
		resource.Maxreq,
		resource.Healthmonitor,
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
		resource.Clttimeout,
		resource.Svrtimeout,
		resource.Cka,
		resource.Tcpb,
		resource.Cmp,
		resource.Maxbandwidth,
		resource.Monthreshold,
		resource.Downstateflush,
		resource.Tcpprofilename,
		resource.Httpprofilename,
		resource.Comment,
		resource.Appflowlog,
		resource.Netprofile,
		resource.Monconnectionclose,
	}

	return c.Put("servicegroup", update)
}
