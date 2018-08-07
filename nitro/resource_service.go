package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Service struct {
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
	Name               string `json:"name,omitempty"`
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

func (resource Service) ToKey() ServiceKey {
	key := ServiceKey{
		resource.Name,
	}

	return key
}

func (key ServiceKey) to_id_args() (string, map[string]string) {
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

type add_service_payload struct {
	Resource Service `json:"service"`
}

func (c *NitroClient) AddService(resource Service) error {
	payload := add_service_payload{
		resource,
	}

	return c.post("service", "", nil, payload)
}

//      LIST

type list_service_result struct {
	Results []Service `json:"service"`
}

func (c *NitroClient) ListService() ([]Service, error) {
	results := list_service_result{}

	if err := c.get("service", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_service_result struct {
	Results []Service `json:"service"`
}

func (c *NitroClient) GetService(key ServiceKey) (*Service, error) {
	var results get_service_result

	id, qs := key.to_id_args()

	if err := c.get("service", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("service element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_service_result struct {
	Results []Count `json:"service"`
}

func (c *NitroClient) CountService() (int, error) {
	var results count_service_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("service", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsService(key ServiceKey) (bool, error) {
	var results count_service_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("service", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteService(key ServiceKey) error {
	id, qs := key.to_id_args()

	return c.delete("service", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
