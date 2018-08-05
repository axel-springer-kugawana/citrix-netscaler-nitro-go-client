package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Auditnslogaction struct {
	Name                string   `json:"name"`
	Acl                 string   `json:"acl,omitempty"`
	Alg                 string   `json:"alg,omitempty"`
	Appflowexport       string   `json:"appflowexport,omitempty"`
	Dateformat          string   `json:"dateformat,omitempty"`
	Domainresolveretry  int      `json:"domainresolveretry,omitempty"`
	Logfacility         string   `json:"logfacility,omitempty"`
	Loglevel            []string `json:"loglevel,omitempty"`
	Lsn                 string   `json:"lsn,omitempty"`
	Serverdomainname    string   `json:"serverdomainname,omitempty"`
	Serverip            string   `json:"serverip,omitempty"`
	Serverport          int      `json:"serverport,omitempty"`
	Sslinterception     string   `json:"sslinterception,omitempty"`
	Subscriberlog       string   `json:"subscriberlog,omitempty"`
	Tcp                 string   `json:"tcp,omitempty"`
	Timezone            string   `json:"timezone,omitempty"`
	Userdefinedauditlog string   `json:"userdefinedauditlog,omitempty"`
}

func auditnslogaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AuditnslogactionUnset struct {
	Name                string `json:"name"`
	Serverip            bool   `json:"serverip,omitempty"`
	Serverdomainname    bool   `json:"serverdomainname,omitempty"`
	Domainresolveretry  bool   `json:"domainresolveretry,omitempty"`
	Serverport          bool   `json:"serverport,omitempty"`
	Loglevel            bool   `json:"loglevel,omitempty"`
	Dateformat          bool   `json:"dateformat,omitempty"`
	Logfacility         bool   `json:"logfacility,omitempty"`
	Tcp                 bool   `json:"tcp,omitempty"`
	Acl                 bool   `json:"acl,omitempty"`
	Timezone            bool   `json:"timezone,omitempty"`
	Userdefinedauditlog bool   `json:"userdefinedauditlog,omitempty"`
	Appflowexport       bool   `json:"appflowexport,omitempty"`
	Lsn                 bool   `json:"lsn,omitempty"`
	Alg                 bool   `json:"alg,omitempty"`
	Subscriberlog       bool   `json:"subscriberlog,omitempty"`
	Sslinterception     bool   `json:"sslinterception,omitempty"`
}

type update_auditnslogaction struct {
	Name                string   `json:"name"`
	Serverip            string   `json:"serverip,omitempty"`
	Serverdomainname    string   `json:"serverdomainname,omitempty"`
	Domainresolveretry  int      `json:"domainresolveretry,omitempty"`
	Serverport          int      `json:"serverport,omitempty"`
	Loglevel            []string `json:"loglevel,omitempty"`
	Dateformat          string   `json:"dateformat,omitempty"`
	Logfacility         string   `json:"logfacility,omitempty"`
	Tcp                 string   `json:"tcp,omitempty"`
	Acl                 string   `json:"acl,omitempty"`
	Timezone            string   `json:"timezone,omitempty"`
	Userdefinedauditlog string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport       string   `json:"appflowexport,omitempty"`
	Lsn                 string   `json:"lsn,omitempty"`
	Alg                 string   `json:"alg,omitempty"`
	Subscriberlog       string   `json:"subscriberlog,omitempty"`
	Sslinterception     string   `json:"sslinterception,omitempty"`
}

type rename_auditnslogaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_auditnslogaction_payload struct {
	Resource Auditnslogaction `json:"auditnslogaction"`
}

type rename_auditnslogaction_payload struct {
	Rename rename_auditnslogaction `json:"auditnslogaction"`
}

type unset_auditnslogaction_payload struct {
	Unset AuditnslogactionUnset `json:"auditnslogaction"`
}

type update_auditnslogaction_payload struct {
	Update update_auditnslogaction `json:"auditnslogaction"`
}

type get_auditnslogaction_result struct {
	Results []Auditnslogaction `json:"auditnslogaction"`
}

type count_auditnslogaction_result struct {
	Results []Count `json:"auditnslogaction"`
}

func (c *NitroClient) AddAuditnslogaction(resource Auditnslogaction) error {
	payload := add_auditnslogaction_payload{
		resource,
	}

	return c.post("auditnslogaction", "", nil, payload)
}

func (c *NitroClient) RenameAuditnslogaction(name string, newName string) error {
	payload := rename_auditnslogaction_payload{
		rename_auditnslogaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("auditnslogaction", "", qs, payload)
}

func (c *NitroClient) CountAuditnslogaction() (int, error) {
	var results count_auditnslogaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("auditnslogaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAuditnslogaction(key string) (bool, error) {
	var results count_auditnslogaction_result

	id, qs := auditnslogaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("auditnslogaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAuditnslogaction() ([]Auditnslogaction, error) {
	results := get_auditnslogaction_result{}

	if err := c.get("auditnslogaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAuditnslogaction(key string) (*Auditnslogaction, error) {
	var results get_auditnslogaction_result

	id, qs := auditnslogaction_key_to_id_args(key)

	if err := c.get("auditnslogaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one auditnslogaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("auditnslogaction element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAuditnslogaction(key string) error {
	id, qs := auditnslogaction_key_to_id_args(key)

	return c.delete("auditnslogaction", id, qs)
}

func (c *NitroClient) UnsetAuditnslogaction(unset AuditnslogactionUnset) error {
	payload := unset_auditnslogaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("auditnslogaction", "", qs, payload)
}

func (c *NitroClient) UpdateAuditnslogaction(resource Auditnslogaction) error {
	payload := update_auditnslogaction_payload{
		update_auditnslogaction{
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
		},
	}

	return c.put("auditnslogaction", "", nil, payload)
}
