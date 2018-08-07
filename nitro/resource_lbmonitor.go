package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbmonitor struct {
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
	Monitorname                    string `json:"monitorname,omitempty"`
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

func (resource Lbmonitor) ToKey() LbmonitorKey {
	key := LbmonitorKey{
		resource.Monitorname,
		resource.Type,
	}

	return key
}

func (key LbmonitorKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Monitorname
	args = append(args, "type:"+key.Type)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbmonitorKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbmonitorKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbmonitor_payload struct {
	Resource Lbmonitor `json:"lbmonitor"`
}

func (c *NitroClient) AddLbmonitor(resource Lbmonitor) error {
	payload := add_lbmonitor_payload{
		resource,
	}

	return c.post("lbmonitor", "", nil, payload)
}

//      LIST

type list_lbmonitor_result struct {
	Results []Lbmonitor `json:"lbmonitor"`
}

func (c *NitroClient) ListLbmonitor() ([]Lbmonitor, error) {
	results := list_lbmonitor_result{}

	if err := c.get("lbmonitor", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbmonitor_result struct {
	Results []Lbmonitor `json:"lbmonitor"`
}

func (c *NitroClient) GetLbmonitor(key LbmonitorKey) (*Lbmonitor, error) {
	var results get_lbmonitor_result

	id, qs := key.to_id_args()

	if err := c.get("lbmonitor", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbmonitor element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbmonitor_result struct {
	Results []Count `json:"lbmonitor"`
}

func (c *NitroClient) CountLbmonitor() (int, error) {
	var results count_lbmonitor_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmonitor", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbmonitor(key LbmonitorKey) (bool, error) {
	var results count_lbmonitor_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbmonitor", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbmonitor(key LbmonitorKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbmonitor", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      ENABLE

type state_lbmonitor_payload struct {
	State LbmonitorKey `json:"lbmonitor"`
}

func (c *NitroClient) EnableLbmonitor(key LbmonitorKey) error {
	payload := state_lbmonitor_payload{
		key,
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("lbmonitor", "", qs, payload)
}

//      DISABLE

func (c *NitroClient) DisableLbmonitor(key LbmonitorKey) error {
	payload := state_lbmonitor_payload{
		key,
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("lbmonitor", "", qs, payload)
}
