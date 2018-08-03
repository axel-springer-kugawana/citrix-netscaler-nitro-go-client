package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Service struct {
	Name               string `json:"name"`
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

func service_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type ServiceUnset struct {
	Name               string `json:"name"`
	Maxclient          bool   `json:"maxclient,string,omitempty"`
	Maxreq             bool   `json:"maxreq,string,omitempty"`
	Cacheable          bool   `json:"cacheable,string,omitempty"`
	Cip                bool   `json:"cip,string,omitempty"`
	Cipheader          bool   `json:"cipheader,string,omitempty"`
	Usip               bool   `json:"usip,string,omitempty"`
	Pathmonitor        bool   `json:"pathmonitor,string,omitempty"`
	Pathmonitorindv    bool   `json:"pathmonitorindv,string,omitempty"`
	Useproxyport       bool   `json:"useproxyport,string,omitempty"`
	Sc                 bool   `json:"sc,string,omitempty"`
	Sp                 bool   `json:"sp,string,omitempty"`
	Rtspsessionidremap bool   `json:"rtspsessionidremap,string,omitempty"`
	Healthmonitor      bool   `json:"healthmonitor,string,omitempty"`
	Clttimeout         bool   `json:"clttimeout,string,omitempty"`
	Svrtimeout         bool   `json:"svrtimeout,string,omitempty"`
	Customserverid     bool   `json:"customserverid,string,omitempty"`
	Serverid           bool   `json:"serverid,string,omitempty"`
	Cka                bool   `json:"cka,string,omitempty"`
	Tcpb               bool   `json:"tcpb,string,omitempty"`
	Cmp                bool   `json:"cmp,string,omitempty"`
	Maxbandwidth       bool   `json:"maxbandwidth,string,omitempty"`
	Accessdown         bool   `json:"accessdown,string,omitempty"`
	Monthreshold       bool   `json:"monthreshold,string,omitempty"`
	Downstateflush     bool   `json:"downstateflush,string,omitempty"`
	Tcpprofilename     bool   `json:"tcpprofilename,string,omitempty"`
	Httpprofilename    bool   `json:"httpprofilename,string,omitempty"`
	Hashid             bool   `json:"hashid,string,omitempty"`
	Comment            bool   `json:"comment,string,omitempty"`
	Appflowlog         bool   `json:"appflowlog,string,omitempty"`
	Netprofile         bool   `json:"netprofile,string,omitempty"`
	Processlocal       bool   `json:"processlocal,string,omitempty"`
	Dnsprofilename     bool   `json:"dnsprofilename,string,omitempty"`
	Monconnectionclose bool   `json:"monconnectionclose,string,omitempty"`
}

type update_service struct {
	Name               string `json:"name"`
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

type rename_service struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_service_payload struct {
	Resource Service `json:"service"`
}

type rename_service_payload struct {
	Rename rename_service `json:"service"`
}

type unset_service_payload struct {
	Unset ServiceUnset `json:"service"`
}

type update_service_payload struct {
	Update update_service `json:"service"`
}

type get_service_result struct {
	Results []Service `json:"service"`
}

type count_service_result struct {
	Results []Count `json:"service"`
}

func (c *NitroClient) AddService(resource Service) error {
	payload := add_service_payload{
		resource,
	}

	return c.post("service", "", nil, payload)
}

func (c *NitroClient) RenameService(name string, newName string) error {
	payload := rename_service_payload{
		rename_service{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("service", "", qs, payload)
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

func (c *NitroClient) ExistsService(key string) (bool, error) {
	var results count_service_result

	id, qs := service_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("service", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListService() ([]Service, error) {
	results := get_service_result{}

	if err := c.get("service", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetService(key string) (*Service, error) {
	var results get_service_result

	id, qs := service_key_to_id_args(key)

	if err := c.get("service", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one service element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("service element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteService(key string) error {
	id, qs := service_key_to_id_args(key)

	return c.delete("service", id, qs)
}

func (c *NitroClient) UnsetService(unset ServiceUnset) error {
	payload := unset_service_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("service", "", qs, payload)
}

func (c *NitroClient) UpdateService(resource Service) error {
	payload := update_service_payload{
		update_service{
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
		},
	}

	return c.put("service", "", nil, payload)
}
