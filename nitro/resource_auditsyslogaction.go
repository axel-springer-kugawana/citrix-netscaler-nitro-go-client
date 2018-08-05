package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditsyslogaction struct {
	Name                 string   `json:"name"`
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

func auditsyslogaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AuditsyslogactionUnset struct {
	Name                 string `json:"name"`
	Serverip             bool   `json:"serverip,string,omitempty"`
	Serverdomainname     bool   `json:"serverdomainname,string,omitempty"`
	Domainresolveretry   bool   `json:"domainresolveretry,string,omitempty"`
	Serverport           bool   `json:"serverport,string,omitempty"`
	Loglevel             bool   `json:"loglevel,string,omitempty"`
	Dateformat           bool   `json:"dateformat,string,omitempty"`
	Logfacility          bool   `json:"logfacility,string,omitempty"`
	Tcp                  bool   `json:"tcp,string,omitempty"`
	Acl                  bool   `json:"acl,string,omitempty"`
	Timezone             bool   `json:"timezone,string,omitempty"`
	Userdefinedauditlog  bool   `json:"userdefinedauditlog,string,omitempty"`
	Appflowexport        bool   `json:"appflowexport,string,omitempty"`
	Lsn                  bool   `json:"lsn,string,omitempty"`
	Alg                  bool   `json:"alg,string,omitempty"`
	Subscriberlog        bool   `json:"subscriberlog,string,omitempty"`
	Sslinterception      bool   `json:"sslinterception,string,omitempty"`
	Lbvservername        bool   `json:"lbvservername,string,omitempty"`
	Tcpprofilename       bool   `json:"tcpprofilename,string,omitempty"`
	Maxlogdatasizetohold bool   `json:"maxlogdatasizetohold,string,omitempty"`
	Dns                  bool   `json:"dns,string,omitempty"`
	Netprofile           bool   `json:"netprofile,string,omitempty"`
}

type update_auditsyslogaction struct {
	Name                 string   `json:"name"`
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

type rename_auditsyslogaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_auditsyslogaction_payload struct {
	Resource Auditsyslogaction `json:"auditsyslogaction"`
}

type rename_auditsyslogaction_payload struct {
	Rename rename_auditsyslogaction `json:"auditsyslogaction"`
}

type unset_auditsyslogaction_payload struct {
	Unset AuditsyslogactionUnset `json:"auditsyslogaction"`
}

type update_auditsyslogaction_payload struct {
	Update update_auditsyslogaction `json:"auditsyslogaction"`
}

type get_auditsyslogaction_result struct {
	Results []Auditsyslogaction `json:"auditsyslogaction"`
}

type count_auditsyslogaction_result struct {
	Results []Count `json:"auditsyslogaction"`
}

func (c *NitroClient) AddAuditsyslogaction(resource Auditsyslogaction) error {
	payload := add_auditsyslogaction_payload{
		resource,
	}

	return c.post("auditsyslogaction", "", nil, payload)
}

func (c *NitroClient) RenameAuditsyslogaction(name string, newName string) error {
	payload := rename_auditsyslogaction_payload{
		rename_auditsyslogaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("auditsyslogaction", "", qs, payload)
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

func (c *NitroClient) ExistsAuditsyslogaction(key string) (bool, error) {
	var results count_auditsyslogaction_result

	id, qs := auditsyslogaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("auditsyslogaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuditsyslogaction() ([]Auditsyslogaction, error) {
	results := get_auditsyslogaction_result{}

	if err := c.get("auditsyslogaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuditsyslogaction(key string) (*Auditsyslogaction, error) {
	var results get_auditsyslogaction_result

	id, qs := auditsyslogaction_key_to_id_args(key)

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

func (c *NitroClient) DeleteAuditsyslogaction(key string) error {
	id, qs := auditsyslogaction_key_to_id_args(key)

	return c.delete("auditsyslogaction", id, qs)
}

func (c *NitroClient) UnsetAuditsyslogaction(unset AuditsyslogactionUnset) error {
	payload := unset_auditsyslogaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("auditsyslogaction", "", qs, payload)
}

func (c *NitroClient) UpdateAuditsyslogaction(resource Auditsyslogaction) error {
	payload := update_auditsyslogaction_payload{
		update_auditsyslogaction{
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
		},
	}

	return c.put("auditsyslogaction", "", nil, payload)
}
