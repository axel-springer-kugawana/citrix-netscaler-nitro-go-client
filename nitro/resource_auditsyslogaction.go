package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditsyslogaction struct {
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,string,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Transport            string   `json:"transport,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditsyslogactionKey struct {
	Name string
}

func (resource Auditsyslogaction) ToKey() AuditsyslogactionKey {
	key := AuditsyslogactionKey{
		resource.Name,
	}

	return key
}

func (key AuditsyslogactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AuditsyslogactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuditsyslogactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_auditsyslogaction_payload struct {
	Resource Auditsyslogaction `json:"auditsyslogaction"`
}

func (c *NitroClient) AddAuditsyslogaction(resource Auditsyslogaction) error {
	payload := add_auditsyslogaction_payload{
		resource,
	}

	return c.post("auditsyslogaction", "", nil, payload)
}

//      LIST

type list_auditsyslogaction_result struct {
	Results []Auditsyslogaction `json:"auditsyslogaction"`
}

func (c *NitroClient) ListAuditsyslogaction() ([]Auditsyslogaction, error) {
	results := list_auditsyslogaction_result{}

	if err := c.get("auditsyslogaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_auditsyslogaction_result struct {
	Results []Auditsyslogaction `json:"auditsyslogaction"`
}

func (c *NitroClient) GetAuditsyslogaction(key AuditsyslogactionKey) (*Auditsyslogaction, error) {
	var results get_auditsyslogaction_result

	id, qs := key.to_id_args()

	if err := c.get("auditsyslogaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one auditsyslogaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("auditsyslogaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_auditsyslogaction_result struct {
	Results []Count `json:"auditsyslogaction"`
}

func (c *NitroClient) CountAuditsyslogaction() (int, error) {
	var results count_auditsyslogaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("auditsyslogaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAuditsyslogaction(key AuditsyslogactionKey) (bool, error) {
	var results count_auditsyslogaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("auditsyslogaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuditsyslogaction(key AuditsyslogactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("auditsyslogaction", id, qs)
}

//      UPDATE

type AuditsyslogactionUpdate struct {
	Name                 string   `json:"name,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,string,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
}

func (resource Auditsyslogaction) ToUpdate() AuditsyslogactionUpdate {
	update := AuditsyslogactionUpdate{
		resource.Name,
		resource.Serverip,
		resource.Serverdomainname,
		resource.Domainresolveretry,
		resource.Serverport,
		resource.Loglevel,
		resource.Dateformat,
		resource.Logfacility,
		resource.Tcp,
		resource.Acl,
		resource.Timezone,
		resource.Userdefinedauditlog,
		resource.Appflowexport,
		resource.Lsn,
		resource.Alg,
		resource.Subscriberlog,
		resource.Sslinterception,
		resource.Lbvservername,
		resource.Tcpprofilename,
		resource.Maxlogdatasizetohold,
		resource.Dns,
		resource.Netprofile,
	}

	return update
}

type update_auditsyslogaction_payload struct {
	Update AuditsyslogactionUpdate `json:"auditsyslogaction"`
}

func (c *NitroClient) UpdateAuditsyslogaction(update AuditsyslogactionUpdate) error {
	payload := update_auditsyslogaction_payload{
		update,
	}

	return c.put("auditsyslogaction", "", nil, payload)
}

//      UNSET

type AuditsyslogactionUnset struct {
	Name                 string `json:"name,omitempty"`
	Serverip             bool   `json:"serverip,omitempty"`
	Serverdomainname     bool   `json:"serverdomainname,omitempty"`
	Domainresolveretry   bool   `json:"domainresolveretry,omitempty"`
	Serverport           bool   `json:"serverport,omitempty"`
	Loglevel             bool   `json:"loglevel,omitempty"`
	Dateformat           bool   `json:"dateformat,omitempty"`
	Logfacility          bool   `json:"logfacility,omitempty"`
	Tcp                  bool   `json:"tcp,omitempty"`
	Acl                  bool   `json:"acl,omitempty"`
	Timezone             bool   `json:"timezone,omitempty"`
	Userdefinedauditlog  bool   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        bool   `json:"appflowexport,omitempty"`
	Lsn                  bool   `json:"lsn,omitempty"`
	Alg                  bool   `json:"alg,omitempty"`
	Subscriberlog        bool   `json:"subscriberlog,omitempty"`
	Sslinterception      bool   `json:"sslinterception,omitempty"`
	Lbvservername        bool   `json:"lbvservername,omitempty"`
	Tcpprofilename       bool   `json:"tcpprofilename,omitempty"`
	Maxlogdatasizetohold bool   `json:"maxlogdatasizetohold,omitempty"`
	Dns                  bool   `json:"dns,omitempty"`
	Netprofile           bool   `json:"netprofile,omitempty"`
}

func (resource Auditsyslogaction) ToUnset() AuditsyslogactionUnset {
	unset := AuditsyslogactionUnset{
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
	}

	return unset
}

type unset_auditsyslogaction_payload struct {
	Unset AuditsyslogactionUnset `json:"auditsyslogaction"`
}

func (c *NitroClient) UnsetAuditsyslogaction(unset AuditsyslogactionUnset) error {
	payload := unset_auditsyslogaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("auditsyslogaction", "", qs, payload)
}

//      RENAME
//      TODO
