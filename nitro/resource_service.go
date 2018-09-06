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

func (key ServiceKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key ServiceKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServiceKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
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

type ServiceUpdate struct {
	Name               string `json:"name,omitempty"`
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

func (resource Service) ToUpdate() ServiceUpdate {
	update := ServiceUpdate{
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

	return update
}

type update_service_payload struct {
	Update ServiceUpdate `json:"service"`
}

func (c *NitroClient) UpdateService(update ServiceUpdate) error {
	payload := update_service_payload{
		update,
	}

	return c.put("service", "", nil, payload)
}

//      UNSET

type ServiceUnset struct {
	Name               string `json:"name,omitempty"`
	Maxclient          bool   `json:"maxclient,omitempty"`
	Maxreq             bool   `json:"maxreq,omitempty"`
	Cacheable          bool   `json:"cacheable,omitempty"`
	Cip                bool   `json:"cip,omitempty"`
	Cipheader          bool   `json:"cipheader,omitempty"`
	Usip               bool   `json:"usip,omitempty"`
	Pathmonitor        bool   `json:"pathmonitor,omitempty"`
	Pathmonitorindv    bool   `json:"pathmonitorindv,omitempty"`
	Useproxyport       bool   `json:"useproxyport,omitempty"`
	Sc                 bool   `json:"sc,omitempty"`
	Sp                 bool   `json:"sp,omitempty"`
	Rtspsessionidremap bool   `json:"rtspsessionidremap,omitempty"`
	Healthmonitor      bool   `json:"healthmonitor,omitempty"`
	Clttimeout         bool   `json:"clttimeout,omitempty"`
	Svrtimeout         bool   `json:"svrtimeout,omitempty"`
	Customserverid     bool   `json:"customserverid,omitempty"`
	Serverid           bool   `json:"serverid,omitempty"`
	Cka                bool   `json:"cka,omitempty"`
	Tcpb               bool   `json:"tcpb,omitempty"`
	Cmp                bool   `json:"cmp,omitempty"`
	Maxbandwidth       bool   `json:"maxbandwidth,omitempty"`
	Accessdown         bool   `json:"accessdown,omitempty"`
	Monthreshold       bool   `json:"monthreshold,omitempty"`
	Downstateflush     bool   `json:"downstateflush,omitempty"`
	Tcpprofilename     bool   `json:"tcpprofilename,omitempty"`
	Httpprofilename    bool   `json:"httpprofilename,omitempty"`
	Hashid             bool   `json:"hashid,omitempty"`
	Comment            bool   `json:"comment,omitempty"`
	Appflowlog         bool   `json:"appflowlog,omitempty"`
	Netprofile         bool   `json:"netprofile,omitempty"`
	Processlocal       bool   `json:"processlocal,omitempty"`
	Dnsprofilename     bool   `json:"dnsprofilename,omitempty"`
	Monconnectionclose bool   `json:"monconnectionclose,omitempty"`
}

func (resource Service) ToUnset() ServiceUnset {
	unset := ServiceUnset{
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
	}

	return unset
}

type unset_service_payload struct {
	Unset ServiceUnset `json:"service"`
}

func (c *NitroClient) UnsetService(unset ServiceUnset) error {
	payload := unset_service_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("service", "", qs, payload)
}

//      RENAME
//      TODO
