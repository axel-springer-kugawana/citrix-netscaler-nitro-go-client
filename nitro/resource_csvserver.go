package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Csvserver struct {
	Name                    string `json:"name"`
	State                   string `json:"state,omitempty"`
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
	Stateupdate             string `json:"stateupdate,omitempty"`
	Tcpprofilename          string `json:"tcpprofilename,omitempty"`
	Td                      int    `json:"td,string,omitempty"`
	Vipheader               string `json:"vipheader,omitempty"`
}

func csvserver_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type CsvserverUnset struct {
	Name                    string `json:"name"`
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

type update_csvserver struct {
	Name                    string `json:"name"`
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

type rename_csvserver struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_csvserver_payload struct {
	Resource Csvserver `json:"csvserver"`
}

type rename_csvserver_payload struct {
	Rename rename_csvserver `json:"csvserver"`
}

type state_csvserver struct {
	Key string `json:"name"`
}

type state_csvserver_payload struct {
	Sate state_csvserver `json:"csvserver"`
}

type unset_csvserver_payload struct {
	Unset CsvserverUnset `json:"csvserver"`
}

type update_csvserver_payload struct {
	Update update_csvserver `json:"csvserver"`
}

type get_csvserver_result struct {
	Results []Csvserver `json:"csvserver"`
}

type count_csvserver_result struct {
	Results []Count `json:"csvserver"`
}

func (c *NitroClient) AddCsvserver(resource Csvserver) error {
	payload := add_csvserver_payload{
		resource,
	}

	return c.post("csvserver", "", nil, payload)
}

func (c *NitroClient) RenameCsvserver(name string, newName string) error {
	payload := rename_csvserver_payload{
		rename_csvserver{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("csvserver", "", qs, payload)
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

func (c *NitroClient) ExistsCsvserver(key string) (bool, error) {
	var results count_csvserver_result

	id, qs := csvserver_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("csvserver", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCsvserver() ([]Csvserver, error) {
	results := get_csvserver_result{}

	if err := c.get("csvserver", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCsvserver(key string) (*Csvserver, error) {
	var results get_csvserver_result

	id, qs := csvserver_key_to_id_args(key)

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

func (c *NitroClient) DeleteCsvserver(key string) error {
	id, qs := csvserver_key_to_id_args(key)

	return c.delete("csvserver", id, qs)
}

func (c *NitroClient) UnsetCsvserver(unset CsvserverUnset) error {
	payload := unset_csvserver_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("csvserver", "", qs, payload)
}

func (c *NitroClient) UpdateCsvserver(resource Csvserver) error {
	payload := update_csvserver_payload{
		update_csvserver{
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
		},
	}

	return c.put("csvserver", "", nil, payload)
}

func (c *NitroClient) EnableCsvserver(key string) error {
	payload := state_csvserver_payload{
		state_csvserver{
			key,
		},
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("csvserver", "", qs, payload)
}

func (c *NitroClient) DisableCsvserver(key string) error {
	payload := state_csvserver_payload{
		state_csvserver{
			key,
		},
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("csvserver", "", qs, payload)
}
