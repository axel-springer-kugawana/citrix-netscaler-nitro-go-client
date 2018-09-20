package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Auditnslogaction struct {
	Acl                 string   `json:"acl,omitempty"`
	Alg                 string   `json:"alg,omitempty"`
	Appflowexport       string   `json:"appflowexport,omitempty"`
	Dateformat          string   `json:"dateformat,omitempty"`
	Domainresolveretry  int      `json:"domainresolveretry,omitempty"`
	Logfacility         string   `json:"logfacility,omitempty"`
	Loglevel            []string `json:"loglevel,omitempty"`
	Lsn                 string   `json:"lsn,omitempty"`
	Name                string   `json:"name,omitempty"`
	Serverdomainname    string   `json:"serverdomainname,omitempty"`
	Serverip            string   `json:"serverip,omitempty"`
	Serverport          int      `json:"serverport,omitempty"`
	Sslinterception     string   `json:"sslinterception,omitempty"`
	Subscriberlog       string   `json:"subscriberlog,omitempty"`
	Tcp                 string   `json:"tcp,omitempty"`
	Timezone            string   `json:"timezone,omitempty"`
	Userdefinedauditlog string   `json:"userdefinedauditlog,omitempty"`
}

type AuditnslogactionKey struct {
	Name string
}

func (resource Auditnslogaction) ToKey() AuditnslogactionKey {
	key := AuditnslogactionKey{
		resource.Name,
	}

	return key
}

func (key AuditnslogactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AuditnslogactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuditnslogactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_auditnslogaction_payload struct {
	Resource Auditnslogaction `json:"auditnslogaction"`
}

func (c *NitroClient) AddAuditnslogaction(resource Auditnslogaction) error {
	payload := add_auditnslogaction_payload{
		resource,
	}

	return c.post("auditnslogaction", "", nil, payload)
}

//      LIST

type list_auditnslogaction_result struct {
	Results []Auditnslogaction `json:"auditnslogaction"`
}

func (c *NitroClient) ListAuditnslogaction() ([]Auditnslogaction, error) {
	results := list_auditnslogaction_result{}

	if err := c.get("auditnslogaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_auditnslogaction_result struct {
	Results []Auditnslogaction `json:"auditnslogaction"`
}

func (c *NitroClient) GetAuditnslogaction(key AuditnslogactionKey) (*Auditnslogaction, error) {
	var results get_auditnslogaction_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_auditnslogaction_result struct {
	Results []Count `json:"auditnslogaction"`
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

//      EXISTS

func (c *NitroClient) ExistsAuditnslogaction(key AuditnslogactionKey) (bool, error) {
	var results count_auditnslogaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("auditnslogaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuditnslogaction(key AuditnslogactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("auditnslogaction", id, qs)
}

//      UPDATE

type AuditnslogactionUpdate struct {
	Name                string   `json:"name,omitempty"`
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

func (resource Auditnslogaction) ToUpdate() AuditnslogactionUpdate {
	update := AuditnslogactionUpdate{
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
	}

	return update
}

type update_auditnslogaction_payload struct {
	Update AuditnslogactionUpdate `json:"auditnslogaction"`
}

func (c *NitroClient) UpdateAuditnslogaction(update AuditnslogactionUpdate) error {
	payload := update_auditnslogaction_payload{
		update,
	}

	return c.put("auditnslogaction", "", nil, payload)
}

//      UNSET

type AuditnslogactionUnset struct {
	Name                string `json:"name,omitempty"`
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

func (resource Auditnslogaction) ToUnset() AuditnslogactionUnset {
	unset := AuditnslogactionUnset{
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
	}

	return unset
}

type unset_auditnslogaction_payload struct {
	Unset AuditnslogactionUnset `json:"auditnslogaction"`
}

func (c *NitroClient) UnsetAuditnslogaction(unset AuditnslogactionUnset) error {
	payload := unset_auditnslogaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("auditnslogaction", "", qs, payload)
}

//      RENAME
//      TODO
