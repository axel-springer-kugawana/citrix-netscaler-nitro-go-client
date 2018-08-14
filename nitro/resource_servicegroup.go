package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Servicegroup struct {
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
	Servicegroupname   string `json:"servicegroupname,omitempty"`
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

type ServicegroupKey struct {
	Servicegroupname string
}

func (resource Servicegroup) ToKey() ServicegroupKey {
	key := ServicegroupKey{
		resource.Servicegroupname,
	}

	return key
}

func (key ServicegroupKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Servicegroupname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key ServicegroupKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServicegroupKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_servicegroup_payload struct {
	Resource Servicegroup `json:"servicegroup"`
}

func (c *NitroClient) AddServicegroup(resource Servicegroup) error {
	payload := add_servicegroup_payload{
		resource,
	}

	return c.post("servicegroup", "", nil, payload)
}

//      LIST

type list_servicegroup_result struct {
	Results []Servicegroup `json:"servicegroup"`
}

func (c *NitroClient) ListServicegroup() ([]Servicegroup, error) {
	results := list_servicegroup_result{}

	if err := c.get("servicegroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_servicegroup_result struct {
	Results []Servicegroup `json:"servicegroup"`
}

func (c *NitroClient) GetServicegroup(key ServicegroupKey) (*Servicegroup, error) {
	var results get_servicegroup_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_servicegroup_result struct {
	Results []Count `json:"servicegroup"`
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

//      EXISTS

func (c *NitroClient) ExistsServicegroup(key ServicegroupKey) (bool, error) {
	var results count_servicegroup_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("servicegroup", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServicegroup(key ServicegroupKey) error {
	id, qs := key.to_id_args()

	return c.delete("servicegroup", id, qs)
}

//      UPDATE

type ServicegroupUpdate struct {
	Servicegroupname   string `json:"servicegroupname,omitempty"`
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

func (resource Servicegroup) ToUpdate() ServicegroupUpdate {
	update := ServicegroupUpdate{
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
	}

	return update
}

type update_servicegroup_payload struct {
	Update ServicegroupUpdate `json:"servicegroup"`
}

func (c *NitroClient) UpdateServicegroup(update ServicegroupUpdate) error {
	payload := update_servicegroup_payload{
		update,
	}

	return c.put("servicegroup", "", nil, payload)
}

//      UNSET

type ServicegroupUnset struct {
	Servicegroupname   string `json:"servicegroupname,omitempty"`
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

func (resource Servicegroup) ToUnset() ServicegroupUnset {
	unset := ServicegroupUnset{
		resource.Servicegroupname,
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

type unset_servicegroup_payload struct {
	Unset ServicegroupUnset `json:"servicegroup"`
}

func (c *NitroClient) UnsetServicegroup(unset ServicegroupUnset) error {
	payload := unset_servicegroup_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("servicegroup", "", qs, payload)
}

//      RENAME
//      TODO
