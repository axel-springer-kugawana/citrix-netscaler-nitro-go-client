package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Servicegroup struct {
	Servicegroupname   string `json:"servicegroupname"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Autoscale          string `json:"autoscale,omitempty"`
	Cacheable          string `json:"cacheable,omitempty"`
	Cachetype          string `json:"cachetype,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Cipheader          string `json:"cipheader,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Memberport         int    `json:"memberport,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Pathmonitor        string `json:"pathmonitor,omitempty"`
	Pathmonitorindv    string `json:"pathmonitorindv,omitempty"`
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	Sc                 string `json:"sc,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sp                 string `json:"sp,omitempty"`
	State              string `json:"state,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Usip               string `json:"usip,omitempty"`
}

func servicegroup_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type ServicegroupUnset struct {
	Servicegroupname   string `json:"servicegroupname"`
	Maxclient          bool   `json:"maxclient,omitempty"`
	Maxreq             bool   `json:"maxreq,omitempty"`
	Healthmonitor      bool   `json:"healthmonitor,omitempty"`
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
	Clttimeout         bool   `json:"clttimeout,omitempty"`
	Svrtimeout         bool   `json:"svrtimeout,omitempty"`
	Cka                bool   `json:"cka,omitempty"`
	Tcpb               bool   `json:"tcpb,omitempty"`
	Cmp                bool   `json:"cmp,omitempty"`
	Maxbandwidth       bool   `json:"maxbandwidth,omitempty"`
	Monthreshold       bool   `json:"monthreshold,omitempty"`
	Downstateflush     bool   `json:"downstateflush,omitempty"`
	Tcpprofilename     bool   `json:"tcpprofilename,omitempty"`
	Httpprofilename    bool   `json:"httpprofilename,omitempty"`
	Comment            bool   `json:"comment,omitempty"`
	Appflowlog         bool   `json:"appflowlog,omitempty"`
	Netprofile         bool   `json:"netprofile,omitempty"`
	Monconnectionclose bool   `json:"monconnectionclose,omitempty"`
}

type update_servicegroup struct {
	Servicegroupname   string `json:"servicegroupname"`
	Maxclient          int    `json:"maxclient,string,omitempty"`
	Maxreq             int    `json:"maxreq,string,omitempty"`
	Healthmonitor      string `json:"healthmonitor,omitempty"`
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
	Clttimeout         int    `json:"clttimeout,omitempty"`
	Svrtimeout         int    `json:"svrtimeout,omitempty"`
	Cka                string `json:"cka,omitempty"`
	Tcpb               string `json:"tcpb,omitempty"`
	Cmp                string `json:"cmp,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,string,omitempty"`
	Monthreshold       int    `json:"monthreshold,string,omitempty"`
	Downstateflush     string `json:"downstateflush,omitempty"`
	Tcpprofilename     string `json:"tcpprofilename,omitempty"`
	Httpprofilename    string `json:"httpprofilename,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Appflowlog         string `json:"appflowlog,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
}

type rename_servicegroup struct {
	Name    string `json:"servicegroupname"`
	Newname string `json:"newname"`
}

type add_servicegroup_payload struct {
	Resource Servicegroup `json:"servicegroup"`
}

type rename_servicegroup_payload struct {
	Rename rename_servicegroup `json:"servicegroup"`
}

type unset_servicegroup_payload struct {
	Unset ServicegroupUnset `json:"servicegroup"`
}

type update_servicegroup_payload struct {
	Update update_servicegroup `json:"servicegroup"`
}

type get_servicegroup_result struct {
	Results []Servicegroup `json:"servicegroup"`
}

type count_servicegroup_result struct {
	Results []Count `json:"servicegroup"`
}

func (c *NitroClient) AddServicegroup(resource Servicegroup) error {
	payload := add_servicegroup_payload{
		resource,
	}

	return c.post("servicegroup", "", nil, payload)
}

func (c *NitroClient) RenameServicegroup(name string, newName string) error {
	payload := rename_servicegroup_payload{
		rename_servicegroup{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("servicegroup", "", qs, payload)
}

func (c *NitroClient) CountServicegroup() (int, error) {
	var results count_servicegroup_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("servicegroup", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsServicegroup(key string) (bool, error) {
	var results count_servicegroup_result

	id, qs := servicegroup_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("servicegroup", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListServicegroup() ([]Servicegroup, error) {
	results := get_servicegroup_result{}

	if err := c.get("servicegroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServicegroup(key string) (*Servicegroup, error) {
	var results get_servicegroup_result

	id, qs := servicegroup_key_to_id_args(key)

	if err := c.get("servicegroup", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one servicegroup element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("servicegroup element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteServicegroup(key string) error {
	id, qs := servicegroup_key_to_id_args(key)

	return c.delete("servicegroup", id, qs)
}

func (c *NitroClient) UnsetServicegroup(unset ServicegroupUnset) error {
	payload := unset_servicegroup_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("servicegroup", "", qs, payload)
}

func (c *NitroClient) UpdateServicegroup(resource Servicegroup) error {
	payload := update_servicegroup_payload{
		update_servicegroup{
			resource.Servicegroupname,
			resource.Maxclient,
			resource.Maxreq,
			resource.Healthmonitor,
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
			resource.Clttimeout,
			resource.Svrtimeout,
			resource.Cka,
			resource.Tcpb,
			resource.Cmp,
			resource.Maxbandwidth,
			resource.Monthreshold,
			resource.Downstateflush,
			resource.Tcpprofilename,
			resource.Httpprofilename,
			resource.Comment,
			resource.Appflowlog,
			resource.Netprofile,
			resource.Monconnectionclose,
		},
	}

	return c.put("servicegroup", "", nil, payload)
}
