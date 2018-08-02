package nitro

type Lbmonitor struct {
	Monitorname                    string `json:"monitorname"`
	State                          string `json:"state,omitempty"`
	Action                         string `json:"action,omitempty"`
	Alertretries                   int    `json:"alertretries,omitempty"`
	Application                    string `json:"application,omitempty"`
	Attribute                      string `json:"attribute,omitempty"`
	Basedn                         string `json:"basedn,omitempty"`
	Binddn                         string `json:"binddn,omitempty"`
	Customheaders                  string `json:"customheaders,omitempty"`
	Database                       string `json:"database,omitempty"`
	Destip                         string `json:"destip,omitempty"`
	Destport                       int    `json:"destport,omitempty"`
	Deviation                      int    `json:"deviation,string,omitempty"`
	Dispatcherip                   string `json:"dispatcherip,omitempty"`
	Dispatcherport                 int    `json:"dispatcherport,omitempty"`
	Domain                         string `json:"domain,omitempty"`
	Downtime                       int    `json:"downtime,omitempty"`
	Evalrule                       string `json:"evalrule,omitempty"`
	Failureretries                 int    `json:"failureretries,omitempty"`
	Filename                       string `json:"filename,omitempty"`
	Filter                         string `json:"filter,omitempty"`
	Firmwarerevision               int    `json:"firmwarerevision,string,omitempty"`
	Group                          string `json:"group,omitempty"`
	Hostipaddress                  string `json:"hostipaddress,omitempty"`
	Hostname                       string `json:"hostname,omitempty"`
	Httprequest                    string `json:"httprequest,omitempty"`
	Inbandsecurityid               string `json:"inbandsecurityid,omitempty"`
	Interval                       int    `json:"interval,omitempty"`
	Iptunnel                       string `json:"iptunnel,omitempty"`
	Kcdaccount                     string `json:"kcdaccount,omitempty"`
	Lasversion                     string `json:"lasversion,omitempty"`
	Logonpointname                 string `json:"logonpointname,omitempty"`
	Lrtm                           string `json:"lrtm,omitempty"`
	Maxforwards                    int    `json:"maxforwards,string,omitempty"`
	Metrictable                    string `json:"metrictable,omitempty"`
	Mssqlprotocolversion           string `json:"mssqlprotocolversion,omitempty"`
	Netprofile                     string `json:"netprofile,omitempty"`
	Oraclesid                      string `json:"oraclesid,omitempty"`
	Originhost                     string `json:"originhost,omitempty"`
	Originrealm                    string `json:"originrealm,omitempty"`
	Password                       string `json:"password,omitempty"`
	Productname                    string `json:"productname,omitempty"`
	Query                          string `json:"query,omitempty"`
	Querytype                      string `json:"querytype,omitempty"`
	Radaccountsession              string `json:"radaccountsession,omitempty"`
	Radaccounttype                 int    `json:"radaccounttype,string,omitempty"`
	Radapn                         string `json:"radapn,omitempty"`
	Radframedip                    string `json:"radframedip,omitempty"`
	Radkey                         string `json:"radkey,omitempty"`
	Radmsisdn                      string `json:"radmsisdn,omitempty"`
	Radnasid                       string `json:"radnasid,omitempty"`
	Radnasip                       string `json:"radnasip,omitempty"`
	Recv                           string `json:"recv,omitempty"`
	Resptimeout                    int    `json:"resptimeout,omitempty"`
	Resptimeoutthresh              int    `json:"resptimeoutthresh,string,omitempty"`
	Retries                        int    `json:"retries,omitempty"`
	Reverse                        string `json:"reverse,omitempty"`
	Rtsprequest                    string `json:"rtsprequest,omitempty"`
	Scriptargs                     string `json:"scriptargs,omitempty"`
	Scriptname                     string `json:"scriptname,omitempty"`
	Secondarypassword              string `json:"secondarypassword,omitempty"`
	Secure                         string `json:"secure,omitempty"`
	Send                           string `json:"send,omitempty"`
	Sipmethod                      string `json:"sipmethod,omitempty"`
	Sipreguri                      string `json:"sipreguri,omitempty"`
	Sipuri                         string `json:"sipuri,omitempty"`
	Sitepath                       string `json:"sitepath,omitempty"`
	Snmpcommunity                  string `json:"snmpcommunity,omitempty"`
	Snmpoid                        string `json:"snmpoid,omitempty"`
	Snmpthreshold                  string `json:"snmpthreshold,omitempty"`
	Snmpversion                    string `json:"snmpversion,omitempty"`
	Sqlquery                       string `json:"sqlquery,omitempty"`
	Sslprofile                     string `json:"sslprofile,omitempty"`
	Storedb                        string `json:"storedb,omitempty"`
	Storefrontacctservice          string `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices string `json:"storefrontcheckbackendservices,omitempty"`
	Storename                      string `json:"storename,omitempty"`
	Successretries                 int    `json:"successretries,omitempty"`
	Tos                            string `json:"tos,omitempty"`
	Tosid                          int    `json:"tosid,string,omitempty"`
	Transparent                    string `json:"transparent,omitempty"`
	Trofscode                      int    `json:"trofscode,string,omitempty"`
	Trofsstring                    string `json:"trofsstring,omitempty"`
	Type                           string `json:"type,omitempty"`
	Units1                         string `json:"units1,omitempty"`
	Units2                         string `json:"units2,omitempty"`
	Units3                         string `json:"units3,omitempty"`
	Units4                         string `json:"units4,omitempty"`
	Username                       string `json:"username,omitempty"`
	Validatecred                   string `json:"validatecred,omitempty"`
	Vendorid                       int    `json:"vendorid,string,omitempty"`
	Vendorspecificvendorid         int    `json:"vendorspecificvendorid,string,omitempty"`
}

type LbmonitorKey struct {
	Monitorname string
	Type        string
}

type lbmonitor_update struct {
	Monitorname                    string `json:"monitorname"`
	Type                           string `json:"type,omitempty"`
	Action                         string `json:"action,omitempty"`
	Httprequest                    string `json:"httprequest,omitempty"`
	Rtsprequest                    string `json:"rtsprequest,omitempty"`
	Customheaders                  string `json:"customheaders,omitempty"`
	Maxforwards                    int    `json:"maxforwards,string,omitempty"`
	Sipmethod                      string `json:"sipmethod,omitempty"`
	Sipreguri                      string `json:"sipreguri,omitempty"`
	Sipuri                         string `json:"sipuri,omitempty"`
	Send                           string `json:"send,omitempty"`
	Recv                           string `json:"recv,omitempty"`
	Query                          string `json:"query,omitempty"`
	Querytype                      string `json:"querytype,omitempty"`
	Username                       string `json:"username,omitempty"`
	Password                       string `json:"password,omitempty"`
	Secondarypassword              string `json:"secondarypassword,omitempty"`
	Logonpointname                 string `json:"logonpointname,omitempty"`
	Lasversion                     string `json:"lasversion,omitempty"`
	Radkey                         string `json:"radkey,omitempty"`
	Radnasid                       string `json:"radnasid,omitempty"`
	Radnasip                       string `json:"radnasip,omitempty"`
	Radaccounttype                 int    `json:"radaccounttype,string,omitempty"`
	Radframedip                    string `json:"radframedip,omitempty"`
	Radapn                         string `json:"radapn,omitempty"`
	Radmsisdn                      string `json:"radmsisdn,omitempty"`
	Radaccountsession              string `json:"radaccountsession,omitempty"`
	Lrtm                           string `json:"lrtm,omitempty"`
	Deviation                      int    `json:"deviation,string,omitempty"`
	Units1                         string `json:"units1,omitempty"`
	Scriptname                     string `json:"scriptname,omitempty"`
	Scriptargs                     string `json:"scriptargs,omitempty"`
	Validatecred                   string `json:"validatecred,omitempty"`
	Domain                         string `json:"domain,omitempty"`
	Dispatcherip                   string `json:"dispatcherip,omitempty"`
	Dispatcherport                 int    `json:"dispatcherport,omitempty"`
	Interval                       int    `json:"interval,omitempty"`
	Units3                         string `json:"units3,omitempty"`
	Resptimeout                    int    `json:"resptimeout,omitempty"`
	Units4                         string `json:"units4,omitempty"`
	Resptimeoutthresh              int    `json:"resptimeoutthresh,string,omitempty"`
	Retries                        int    `json:"retries,omitempty"`
	Failureretries                 int    `json:"failureretries,omitempty"`
	Alertretries                   int    `json:"alertretries,omitempty"`
	Successretries                 int    `json:"successretries,omitempty"`
	Downtime                       int    `json:"downtime,omitempty"`
	Units2                         string `json:"units2,omitempty"`
	Destip                         string `json:"destip,omitempty"`
	Destport                       int    `json:"destport,omitempty"`
	State                          string `json:"state,omitempty"`
	Reverse                        string `json:"reverse,omitempty"`
	Transparent                    string `json:"transparent,omitempty"`
	Iptunnel                       string `json:"iptunnel,omitempty"`
	Tos                            string `json:"tos,omitempty"`
	Tosid                          int    `json:"tosid,string,omitempty"`
	Secure                         string `json:"secure,omitempty"`
	Group                          string `json:"group,omitempty"`
	Filename                       string `json:"filename,omitempty"`
	Basedn                         string `json:"basedn,omitempty"`
	Binddn                         string `json:"binddn,omitempty"`
	Filter                         string `json:"filter,omitempty"`
	Attribute                      string `json:"attribute,omitempty"`
	Database                       string `json:"database,omitempty"`
	Oraclesid                      string `json:"oraclesid,omitempty"`
	Sqlquery                       string `json:"sqlquery,omitempty"`
	Evalrule                       string `json:"evalrule,omitempty"`
	Snmpoid                        string `json:"snmpoid,omitempty"`
	Snmpcommunity                  string `json:"snmpcommunity,omitempty"`
	Snmpthreshold                  string `json:"snmpthreshold,omitempty"`
	Snmpversion                    string `json:"snmpversion,omitempty"`
	Metrictable                    string `json:"metrictable,omitempty"`
	Application                    string `json:"application,omitempty"`
	Sitepath                       string `json:"sitepath,omitempty"`
	Storename                      string `json:"storename,omitempty"`
	Storefrontacctservice          string `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices string `json:"storefrontcheckbackendservices,omitempty"`
	Hostname                       string `json:"hostname,omitempty"`
	Netprofile                     string `json:"netprofile,omitempty"`
	Mssqlprotocolversion           string `json:"mssqlprotocolversion,omitempty"`
	Originhost                     string `json:"originhost,omitempty"`
	Originrealm                    string `json:"originrealm,omitempty"`
	Hostipaddress                  string `json:"hostipaddress,omitempty"`
	Vendorid                       int    `json:"vendorid,string,omitempty"`
	Productname                    string `json:"productname,omitempty"`
	Firmwarerevision               int    `json:"firmwarerevision,string,omitempty"`
	Vendorspecificvendorid         int    `json:"vendorspecificvendorid,string,omitempty"`
	Kcdaccount                     string `json:"kcdaccount,omitempty"`
	Storedb                        string `json:"storedb,omitempty"`
	Trofscode                      int    `json:"trofscode,string,omitempty"`
	Trofsstring                    string `json:"trofsstring,omitempty"`
	Sslprofile                     string `json:"sslprofile,omitempty"`
}

type lbmonitor_payload struct {
	lbmonitor interface{}
}

func lbmonitor_key_to_args(key LbmonitorKey) string {
	result := ""

	result = result + ",type:" + key.Type
	return result
}

func (c *NitroClient) DeleteLbmonitor(key LbmonitorKey) error {
	return c.deleteResourceWithArgs("lbmonitor", key.Monitorname, lbmonitor_key_to_args(key))
}

func (c *NitroClient) GetLbmonitor(key LbmonitorKey) (*Lbmonitor, error) {
	var results struct {
		Lbmonitor []Lbmonitor
	}

	if err := c.getResourceWithArgs("lbmonitor", key.Monitorname, lbmonitor_key_to_args(key), &results); err != nil || len(results.Lbmonitor) != 1 {
		return nil, err
	}

	return &results.Lbmonitor[0], nil
}

func (c *NitroClient) ListLbmonitor() ([]Lbmonitor, error) {
	var results struct {
		Lbmonitor []Lbmonitor
	}

	if err := c.listResources("lbmonitor", &results); err != nil {
		return nil, err
	}

	return results.Lbmonitor, nil
}

func (c *NitroClient) AddLbmonitor(resource Lbmonitor) error {
	return c.addResource("lbmonitor", resource)
}

func (c *NitroClient) RenameLbmonitor(monitorname string, newName string) error {
	return c.renameResource("lbmonitor", "monitorname", monitorname, newName)
}

func (c *NitroClient) UnsetLbmonitor(monitorname string, fields ...string) error {
	return c.unsetResource("lbmonitor", "monitorname", monitorname, fields)
}

func (c *NitroClient) UpdateLbmonitor(resource Lbmonitor) error {
	update := lbmonitor_update{
		resource.Monitorname,
		resource.Type,
		resource.Action,
		resource.Httprequest,
		resource.Rtsprequest,
		resource.Customheaders,
		resource.Maxforwards,
		resource.Sipmethod,
		resource.Sipreguri,
		resource.Sipuri,
		resource.Send,
		resource.Recv,
		resource.Query,
		resource.Querytype,
		resource.Username,
		resource.Password,
		resource.Secondarypassword,
		resource.Logonpointname,
		resource.Lasversion,
		resource.Radkey,
		resource.Radnasid,
		resource.Radnasip,
		resource.Radaccounttype,
		resource.Radframedip,
		resource.Radapn,
		resource.Radmsisdn,
		resource.Radaccountsession,
		resource.Lrtm,
		resource.Deviation,
		resource.Units1,
		resource.Scriptname,
		resource.Scriptargs,
		resource.Validatecred,
		resource.Domain,
		resource.Dispatcherip,
		resource.Dispatcherport,
		resource.Interval,
		resource.Units3,
		resource.Resptimeout,
		resource.Units4,
		resource.Resptimeoutthresh,
		resource.Retries,
		resource.Failureretries,
		resource.Alertretries,
		resource.Successretries,
		resource.Downtime,
		resource.Units2,
		resource.Destip,
		resource.Destport,
		resource.State,
		resource.Reverse,
		resource.Transparent,
		resource.Iptunnel,
		resource.Tos,
		resource.Tosid,
		resource.Secure,
		resource.Group,
		resource.Filename,
		resource.Basedn,
		resource.Binddn,
		resource.Filter,
		resource.Attribute,
		resource.Database,
		resource.Oraclesid,
		resource.Sqlquery,
		resource.Evalrule,
		resource.Snmpoid,
		resource.Snmpcommunity,
		resource.Snmpthreshold,
		resource.Snmpversion,
		resource.Metrictable,
		resource.Application,
		resource.Sitepath,
		resource.Storename,
		resource.Storefrontacctservice,
		resource.Storefrontcheckbackendservices,
		resource.Hostname,
		resource.Netprofile,
		resource.Mssqlprotocolversion,
		resource.Originhost,
		resource.Originrealm,
		resource.Hostipaddress,
		resource.Vendorid,
		resource.Productname,
		resource.Firmwarerevision,
		resource.Vendorspecificvendorid,
		resource.Kcdaccount,
		resource.Storedb,
		resource.Trofscode,
		resource.Trofsstring,
		resource.Sslprofile,
	}

	return c.Put("lbmonitor", update)
}
func (c *NitroClient) EnableLbmonitor(monitorname string) error {
	return c.enableResource("lbmonitor", "monitorname", monitorname)
}

func (c *NitroClient) DisableLbmonitor(monitorname string) error {
	return c.disableResource("lbmonitor", "monitorname", monitorname)
}
