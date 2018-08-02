package nitro

type LbvserverResource struct {
	Name                               string `json:"name"`
	State                              string `json:"state,omitempty"`
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         int    `json:"clttimeout,string,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Datalength                         int    `json:"datalength,string,omitempty"`
	Dataoffset                         int    `json:"dataoffset,string,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	Dns64                              string `json:"dns64,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Hashlength                         int    `json:"hashlength,string,omitempty"`
	Healththreshold                    int    `json:"healththreshold,string,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipv46                              string `json:"ipv46,omitempty"`
	L2conn                             string `json:"l2conn,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,string,omitempty"`
	M                                  string `json:"m,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Maxautoscalemembers                int    `json:"maxautoscalemembers,string,omitempty"`
	Minautoscalemembers                int    `json:"minautoscalemembers,string,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlprotocolversion               int    `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlservercapabilities            int    `json:"mysqlservercapabilities,string,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Netmask                            string `json:"netmask,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Newservicerequest                  int    `json:"newservicerequest,string,omitempty"`
	Newservicerequestincrementinterval int    `json:"newservicerequestincrementinterval,string,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	Port                               int    `json:"port,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Range                              int    `json:"range,string,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Servicetype                        string `json:"servicetype,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               int    `json:"sopersistencetimeout,string,omitempty"`
	Sothreshold                        int    `json:"sothreshold,string,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Td                                 int    `json:"td,string,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Tosid                              int    `json:"tosid,string,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	V6netmasklen                       int    `json:"v6netmasklen,string,omitempty"`
	V6persistmasklen                   int    `json:"v6persistmasklen,string,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
}

type LbvserverKey struct {
	Name string
}

type lbvserver_update struct {
	Name                               string `json:"name"`
	Ipv46                              string `json:"ipv46,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Hashlength                         int    `json:"hashlength,string,omitempty"`
	Netmask                            string `json:"netmask,omitempty"`
	V6netmasklen                       int    `json:"v6netmasklen,string,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	V6persistmasklen                   int    `json:"v6persistmasklen,string,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	M                                  string `json:"m,omitempty"`
	Tosid                              int    `json:"tosid,string,omitempty"`
	Datalength                         int    `json:"datalength,string,omitempty"`
	Dataoffset                         int    `json:"dataoffset,string,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         int    `json:"clttimeout,string,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sothreshold                        int    `json:"sothreshold,string,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               int    `json:"sopersistencetimeout,string,omitempty"`
	Healththreshold                    int    `json:"healththreshold,string,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,string,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	L2conn                             string `json:"l2conn,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlprotocolversion               int    `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Mysqlcharacterset                  int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlservercapabilities            int    `json:"mysqlservercapabilities,string,omitempty"`
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Newservicerequest                  int    `json:"newservicerequest,string,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Newservicerequestincrementinterval int    `json:"newservicerequestincrementinterval,string,omitempty"`
	Minautoscalemembers                int    `json:"minautoscalemembers,string,omitempty"`
	Maxautoscalemembers                int    `json:"maxautoscalemembers,string,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Dns64                              string `json:"dns64,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
}

type lbvserver_payload struct {
	lbvserver interface{}
}

func lbvserver_key_to_args(key LbvserverKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteLbvserver(key LbvserverKey) error {
	return c.deleteResourceWithArgs("lbvserver", key.Name, lbvserver_key_to_args(key))
}

func (c *NitroClient) GetLbvserver(key LbvserverKey) (*LbvserverResource, error) {
	var results struct {
		Lbvserver []LbvserverResource
	}

	if err := c.getResourceWithArgs("lbvserver", key.Name, lbvserver_key_to_args(key), &results); err != nil || len(results.Lbvserver) != 1 {
		return nil, err
	}

	return &results.Lbvserver[0], nil
}

func (c *NitroClient) ListLbvserver() ([]LbvserverResource, error) {
	var results struct {
		Lbvserver []LbvserverResource
	}

	if err := c.listResources("lbvserver", &results); err != nil {
		return nil, err
	}

	return results.Lbvserver, nil
}

func (c *NitroClient) AddLbvserver(resource LbvserverResource) error {
	return c.addResource("lbvserver", resource)
}

func (c *NitroClient) RenameLbvserver(name string, newName string) error {
	return c.renameResource("lbvserver", "name", name, newName)
}

func (c *NitroClient) UnsetLbvserver(name string, fields ...string) error {
	return c.unsetResource("lbvserver", "name", name, fields)
}

func (c *NitroClient) UpdateLbvserver(resource LbvserverResource) error {
	update := lbvserver_update{
		resource.Name,
		resource.Ipv46,
		resource.Ippattern,
		resource.Ipmask,
		resource.Persistencetype,
		resource.Timeout,
		resource.Persistencebackup,
		resource.Backuppersistencetimeout,
		resource.Lbmethod,
		resource.Hashlength,
		resource.Netmask,
		resource.V6netmasklen,
		resource.Backuplbmethod,
		resource.Rule,
		resource.Cookiename,
		resource.Resrule,
		resource.Persistmask,
		resource.V6persistmasklen,
		resource.Pq,
		resource.Sc,
		resource.Rtspnat,
		resource.M,
		resource.Tosid,
		resource.Datalength,
		resource.Dataoffset,
		resource.Sessionless,
		resource.Trofspersistence,
		resource.Connfailover,
		resource.Backupvserver,
		resource.Redirurl,
		resource.Cacheable,
		resource.Clttimeout,
		resource.Somethod,
		resource.Sothreshold,
		resource.Sopersistence,
		resource.Sopersistencetimeout,
		resource.Healththreshold,
		resource.Sobackupaction,
		resource.Redirectportrewrite,
		resource.Downstateflush,
		resource.Insertvserveripport,
		resource.Vipheader,
		resource.Disableprimaryondown,
		resource.Authenticationhost,
		resource.Authentication,
		resource.Authn401,
		resource.Authnvsname,
		resource.Push,
		resource.Pushvserver,
		resource.Pushlabel,
		resource.Pushmulticlients,
		resource.Listenpolicy,
		resource.Listenpriority,
		resource.Tcpprofilename,
		resource.Httpprofilename,
		resource.Dbprofilename,
		resource.Comment,
		resource.L2conn,
		resource.Oracleserverversion,
		resource.Mssqlserverversion,
		resource.Mysqlprotocolversion,
		resource.Mysqlserverversion,
		resource.Mysqlcharacterset,
		resource.Mysqlservercapabilities,
		resource.Appflowlog,
		resource.Netprofile,
		resource.Icmpvsrresponse,
		resource.Rhistate,
		resource.Newservicerequest,
		resource.Newservicerequestunit,
		resource.Newservicerequestincrementinterval,
		resource.Minautoscalemembers,
		resource.Maxautoscalemembers,
		resource.Persistavpno,
		resource.Skippersistency,
		resource.Authnprofile,
		resource.Macmoderetainvlan,
		resource.Dbslb,
		resource.Dns64,
		resource.Bypassaaaa,
		resource.Recursionavailable,
		resource.Processlocal,
		resource.Dnsprofilename,
		resource.Lbprofilename,
		resource.Redirectfromport,
		resource.Httpsredirecturl,
		resource.Retainconnectionsoncluster,
	}

	return c.Put("lbvserver", update)
}
func (c *NitroClient) EnableLbvserver(name string) error {
	return c.enableResource("lbvserver", "name", name)
}

func (c *NitroClient) DisableLbvserver(name string) error {
	return c.disableResource("lbvserver", "name", name)
}
