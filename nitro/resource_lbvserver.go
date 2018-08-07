package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbvserver struct {
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
	Name                               string `json:"name,omitempty"`
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

func (resource Lbvserver) ToKey() LbvserverKey {
	key := LbvserverKey{
		resource.Name,
	}

	return key
}

func (key LbvserverKey) to_id_args() (string, map[string]string) {
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

type add_lbvserver_payload struct {
	Resource Lbvserver `json:"lbvserver"`
}

func (c *NitroClient) AddLbvserver(resource Lbvserver) error {
	payload := add_lbvserver_payload{
		resource,
	}

	return c.post("lbvserver", "", nil, payload)
}

//      LIST

type list_lbvserver_result struct {
	Results []Lbvserver `json:"lbvserver"`
}

func (c *NitroClient) ListLbvserver() ([]Lbvserver, error) {
	results := list_lbvserver_result{}

	if err := c.get("lbvserver", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbvserver_result struct {
	Results []Lbvserver `json:"lbvserver"`
}

func (c *NitroClient) GetLbvserver(key LbvserverKey) (*Lbvserver, error) {
	var results get_lbvserver_result

	id, qs := key.to_id_args()

	if err := c.get("lbvserver", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbvserver element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbvserver_result struct {
	Results []Count `json:"lbvserver"`
}

func (c *NitroClient) CountLbvserver() (int, error) {
	var results count_lbvserver_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbvserver(key LbvserverKey) (bool, error) {
	var results count_lbvserver_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbvserver", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbvserver(key LbvserverKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbvserver", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
