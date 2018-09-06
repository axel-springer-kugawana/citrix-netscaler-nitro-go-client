package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Csvserver struct {
	Appflowlog              string `json:"appflowlog,omitempty"`
	Authentication          string `json:"authentication,omitempty"`
	Authenticationhost      string `json:"authenticationhost,omitempty"`
	Authn401                string `json:"authn401,omitempty"`
	Authnprofile            string `json:"authnprofile,omitempty"`
	Authnvsname             string `json:"authnvsname,omitempty"`
	Backupvserver           string `json:"backupvserver,omitempty"`
	Cacheable               string `json:"cacheable,omitempty"`
	Casesensitive           string `json:"casesensitive,omitempty"`
	Clttimeout              int    `json:"clttimeout,string,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	Dbprofilename           string `json:"dbprofilename,omitempty"`
	Disableprimaryondown    string `json:"disableprimaryondown,omitempty"`
	Dnsprofilename          string `json:"dnsprofilename,omitempty"`
	Downstateflush          string `json:"downstateflush,omitempty"`
	Httpprofilename         string `json:"httpprofilename,omitempty"`
	Icmpvsrresponse         string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport     string `json:"insertvserveripport,omitempty"`
	Ipmask                  string `json:"ipmask,omitempty"`
	Ippattern               string `json:"ippattern,omitempty"`
	Ipv46                   string `json:"ipv46,omitempty"`
	L2conn                  string `json:"l2conn,omitempty"`
	Listenpolicy            string `json:"listenpolicy,omitempty"`
	Listenpriority          int    `json:"listenpriority,string,omitempty"`
	Mssqlserverversion      string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset       int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlprotocolversion    int    `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlservercapabilities int    `json:"mysqlservercapabilities,string,omitempty"`
	Mysqlserverversion      string `json:"mysqlserverversion,omitempty"`
	Name                    string `json:"name,omitempty"`
	Netprofile              string `json:"netprofile,omitempty"`
	Oracleserverversion     string `json:"oracleserverversion,omitempty"`
	Port                    int    `json:"port,omitempty"`
	Precedence              string `json:"precedence,omitempty"`
	Push                    string `json:"push,omitempty"`
	Pushlabel               string `json:"pushlabel,omitempty"`
	Pushmulticlients        string `json:"pushmulticlients,omitempty"`
	Pushvserver             string `json:"pushvserver,omitempty"`
	Range                   int    `json:"range,string,omitempty"`
	Redirectportrewrite     string `json:"redirectportrewrite,omitempty"`
	Redirecturl             string `json:"redirecturl,omitempty"`
	Rhistate                string `json:"rhistate,omitempty"`
	Rtspnat                 string `json:"rtspnat,omitempty"`
	Servicetype             string `json:"servicetype,omitempty"`
	Sobackupaction          string `json:"sobackupaction,omitempty"`
	Somethod                string `json:"somethod,omitempty"`
	Sopersistence           string `json:"sopersistence,omitempty"`
	Sopersistencetimeout    int    `json:"sopersistencetimeout,string,omitempty"`
	Sothreshold             int    `json:"sothreshold,string,omitempty"`
	State                   string `json:"state,omitempty"`
	Stateupdate             string `json:"stateupdate,omitempty"`
	Tcpprofilename          string `json:"tcpprofilename,omitempty"`
	Td                      int    `json:"td,string,omitempty"`
	Vipheader               string `json:"vipheader,omitempty"`
}

type CsvserverKey struct {
	Name string
}

func (resource Csvserver) ToKey() CsvserverKey {
	key := CsvserverKey{
		resource.Name,
	}

	return key
}

func (key CsvserverKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CsvserverKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CsvserverKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_csvserver_payload struct {
	Resource Csvserver `json:"csvserver"`
}

func (c *NitroClient) AddCsvserver(resource Csvserver) error {
	payload := add_csvserver_payload{
		resource,
	}

	return c.post("csvserver", "", nil, payload)
}

//      LIST

type list_csvserver_result struct {
	Results []Csvserver `json:"csvserver"`
}

func (c *NitroClient) ListCsvserver() ([]Csvserver, error) {
	results := list_csvserver_result{}

	if err := c.get("csvserver", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_csvserver_result struct {
	Results []Csvserver `json:"csvserver"`
}

func (c *NitroClient) GetCsvserver(key CsvserverKey) (*Csvserver, error) {
	var results get_csvserver_result

	id, qs := key.to_id_args()

	if err := c.get("csvserver", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csvserver element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csvserver element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_csvserver_result struct {
	Results []Count `json:"csvserver"`
}

func (c *NitroClient) CountCsvserver() (int, error) {
	var results count_csvserver_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("csvserver", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCsvserver(key CsvserverKey) (bool, error) {
	var results count_csvserver_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("csvserver", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCsvserver(key CsvserverKey) error {
	id, qs := key.to_id_args()

	return c.delete("csvserver", id, qs)
}

//      UPDATE

type CsvserverUpdate struct {
	Name                    string `json:"name,omitempty"`
	Ipv46                   string `json:"ipv46,omitempty"`
	Ippattern               string `json:"ippattern,omitempty"`
	Ipmask                  string `json:"ipmask,omitempty"`
	Stateupdate             string `json:"stateupdate,omitempty"`
	Precedence              string `json:"precedence,omitempty"`
	Casesensitive           string `json:"casesensitive,omitempty"`
	Backupvserver           string `json:"backupvserver,omitempty"`
	Redirecturl             string `json:"redirecturl,omitempty"`
	Cacheable               string `json:"cacheable,omitempty"`
	Clttimeout              int    `json:"clttimeout,string,omitempty"`
	Somethod                string `json:"somethod,omitempty"`
	Sopersistence           string `json:"sopersistence,omitempty"`
	Sopersistencetimeout    int    `json:"sopersistencetimeout,string,omitempty"`
	Sothreshold             int    `json:"sothreshold,string,omitempty"`
	Sobackupaction          string `json:"sobackupaction,omitempty"`
	Redirectportrewrite     string `json:"redirectportrewrite,omitempty"`
	Downstateflush          string `json:"downstateflush,omitempty"`
	Disableprimaryondown    string `json:"disableprimaryondown,omitempty"`
	Insertvserveripport     string `json:"insertvserveripport,omitempty"`
	Vipheader               string `json:"vipheader,omitempty"`
	Rtspnat                 string `json:"rtspnat,omitempty"`
	Authenticationhost      string `json:"authenticationhost,omitempty"`
	Authentication          string `json:"authentication,omitempty"`
	Listenpolicy            string `json:"listenpolicy,omitempty"`
	Listenpriority          int    `json:"listenpriority,string,omitempty"`
	Authn401                string `json:"authn401,omitempty"`
	Authnvsname             string `json:"authnvsname,omitempty"`
	Push                    string `json:"push,omitempty"`
	Pushvserver             string `json:"pushvserver,omitempty"`
	Pushlabel               string `json:"pushlabel,omitempty"`
	Pushmulticlients        string `json:"pushmulticlients,omitempty"`
	Tcpprofilename          string `json:"tcpprofilename,omitempty"`
	Httpprofilename         string `json:"httpprofilename,omitempty"`
	Dbprofilename           string `json:"dbprofilename,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	L2conn                  string `json:"l2conn,omitempty"`
	Mssqlserverversion      string `json:"mssqlserverversion,omitempty"`
	Mysqlprotocolversion    int    `json:"mysqlprotocolversion,string,omitempty"`
	Oracleserverversion     string `json:"oracleserverversion,omitempty"`
	Mysqlserverversion      string `json:"mysqlserverversion,omitempty"`
	Mysqlcharacterset       int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlservercapabilities int    `json:"mysqlservercapabilities,string,omitempty"`
	Appflowlog              string `json:"appflowlog,omitempty"`
	Netprofile              string `json:"netprofile,omitempty"`
	Authnprofile            string `json:"authnprofile,omitempty"`
	Icmpvsrresponse         string `json:"icmpvsrresponse,omitempty"`
	Rhistate                string `json:"rhistate,omitempty"`
	Dnsprofilename          string `json:"dnsprofilename,omitempty"`
}

func (resource Csvserver) ToUpdate() CsvserverUpdate {
	update := CsvserverUpdate{
		resource.Name,
		resource.Ipv46,
		resource.Ippattern,
		resource.Ipmask,
		resource.Stateupdate,
		resource.Precedence,
		resource.Casesensitive,
		resource.Backupvserver,
		resource.Redirecturl,
		resource.Cacheable,
		resource.Clttimeout,
		resource.Somethod,
		resource.Sopersistence,
		resource.Sopersistencetimeout,
		resource.Sothreshold,
		resource.Sobackupaction,
		resource.Redirectportrewrite,
		resource.Downstateflush,
		resource.Disableprimaryondown,
		resource.Insertvserveripport,
		resource.Vipheader,
		resource.Rtspnat,
		resource.Authenticationhost,
		resource.Authentication,
		resource.Listenpolicy,
		resource.Listenpriority,
		resource.Authn401,
		resource.Authnvsname,
		resource.Push,
		resource.Pushvserver,
		resource.Pushlabel,
		resource.Pushmulticlients,
		resource.Tcpprofilename,
		resource.Httpprofilename,
		resource.Dbprofilename,
		resource.Comment,
		resource.L2conn,
		resource.Mssqlserverversion,
		resource.Mysqlprotocolversion,
		resource.Oracleserverversion,
		resource.Mysqlserverversion,
		resource.Mysqlcharacterset,
		resource.Mysqlservercapabilities,
		resource.Appflowlog,
		resource.Netprofile,
		resource.Authnprofile,
		resource.Icmpvsrresponse,
		resource.Rhistate,
		resource.Dnsprofilename,
	}

	return update
}

type update_csvserver_payload struct {
	Update CsvserverUpdate `json:"csvserver"`
}

func (c *NitroClient) UpdateCsvserver(update CsvserverUpdate) error {
	payload := update_csvserver_payload{
		update,
	}

	return c.put("csvserver", "", nil, payload)
}

//      UNSET

type CsvserverUnset struct {
	Name                    string `json:"name,omitempty"`
	Ipv46                   bool   `json:"ipv46,omitempty"`
	Ippattern               bool   `json:"ippattern,omitempty"`
	Ipmask                  bool   `json:"ipmask,omitempty"`
	Stateupdate             bool   `json:"stateupdate,omitempty"`
	Precedence              bool   `json:"precedence,omitempty"`
	Casesensitive           bool   `json:"casesensitive,omitempty"`
	Backupvserver           bool   `json:"backupvserver,omitempty"`
	Redirecturl             bool   `json:"redirecturl,omitempty"`
	Cacheable               bool   `json:"cacheable,omitempty"`
	Clttimeout              bool   `json:"clttimeout,omitempty"`
	Somethod                bool   `json:"somethod,omitempty"`
	Sopersistence           bool   `json:"sopersistence,omitempty"`
	Sopersistencetimeout    bool   `json:"sopersistencetimeout,omitempty"`
	Sothreshold             bool   `json:"sothreshold,omitempty"`
	Sobackupaction          bool   `json:"sobackupaction,omitempty"`
	Redirectportrewrite     bool   `json:"redirectportrewrite,omitempty"`
	Downstateflush          bool   `json:"downstateflush,omitempty"`
	Disableprimaryondown    bool   `json:"disableprimaryondown,omitempty"`
	Insertvserveripport     bool   `json:"insertvserveripport,omitempty"`
	Vipheader               bool   `json:"vipheader,omitempty"`
	Rtspnat                 bool   `json:"rtspnat,omitempty"`
	Authenticationhost      bool   `json:"authenticationhost,omitempty"`
	Authentication          bool   `json:"authentication,omitempty"`
	Listenpolicy            bool   `json:"listenpolicy,omitempty"`
	Listenpriority          bool   `json:"listenpriority,omitempty"`
	Authn401                bool   `json:"authn401,omitempty"`
	Authnvsname             bool   `json:"authnvsname,omitempty"`
	Push                    bool   `json:"push,omitempty"`
	Pushvserver             bool   `json:"pushvserver,omitempty"`
	Pushlabel               bool   `json:"pushlabel,omitempty"`
	Pushmulticlients        bool   `json:"pushmulticlients,omitempty"`
	Tcpprofilename          bool   `json:"tcpprofilename,omitempty"`
	Httpprofilename         bool   `json:"httpprofilename,omitempty"`
	Dbprofilename           bool   `json:"dbprofilename,omitempty"`
	Comment                 bool   `json:"comment,omitempty"`
	L2conn                  bool   `json:"l2conn,omitempty"`
	Mssqlserverversion      bool   `json:"mssqlserverversion,omitempty"`
	Mysqlprotocolversion    bool   `json:"mysqlprotocolversion,omitempty"`
	Oracleserverversion     bool   `json:"oracleserverversion,omitempty"`
	Mysqlserverversion      bool   `json:"mysqlserverversion,omitempty"`
	Mysqlcharacterset       bool   `json:"mysqlcharacterset,omitempty"`
	Mysqlservercapabilities bool   `json:"mysqlservercapabilities,omitempty"`
	Appflowlog              bool   `json:"appflowlog,omitempty"`
	Netprofile              bool   `json:"netprofile,omitempty"`
	Authnprofile            bool   `json:"authnprofile,omitempty"`
	Icmpvsrresponse         bool   `json:"icmpvsrresponse,omitempty"`
	Rhistate                bool   `json:"rhistate,omitempty"`
	Dnsprofilename          bool   `json:"dnsprofilename,omitempty"`
}

func (resource Csvserver) ToUnset() CsvserverUnset {
	unset := CsvserverUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_csvserver_payload struct {
	Unset CsvserverUnset `json:"csvserver"`
}

func (c *NitroClient) UnsetCsvserver(unset CsvserverUnset) error {
	payload := unset_csvserver_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("csvserver", "", qs, payload)
}

//      RENAME
//      TODO

//      ENABLE

type state_csvserver_payload struct {
	State CsvserverKey `json:"csvserver"`
}

func (c *NitroClient) EnableCsvserver(key CsvserverKey) error {
	payload := state_csvserver_payload{
		key,
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("csvserver", "", qs, payload)
}

//      DISABLE

func (c *NitroClient) DisableCsvserver(key CsvserverKey) error {
	payload := state_csvserver_payload{
		key,
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("csvserver", "", qs, payload)
}
