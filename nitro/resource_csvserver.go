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

func (key CsvserverKey) to_id_args() (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return id, qs
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
//      TODO

//      UNSET
//      TODO

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
