package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnsprofile struct {
	Dnsprofilename         string `json:"dnsprofilename"`
	Cacheecsresponses      string `json:"cacheecsresponses,omitempty"`
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	Cacherecords           string `json:"cacherecords,omitempty"`
	Dnsanswerseclogging    string `json:"dnsanswerseclogging,omitempty"`
	Dnserrorlogging        string `json:"dnserrorlogging,omitempty"`
	Dnsextendedlogging     string `json:"dnsextendedlogging,omitempty"`
	Dnsquerylogging        string `json:"dnsquerylogging,omitempty"`
	Dropmultiqueryrequest  string `json:"dropmultiqueryrequest,omitempty"`
}

func dnsprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type DnsprofileUnset struct {
	Dnsprofilename         string `json:"dnsprofilename"`
	Dnsquerylogging        bool   `json:"dnsquerylogging,omitempty"`
	Dnsanswerseclogging    bool   `json:"dnsanswerseclogging,omitempty"`
	Dnsextendedlogging     bool   `json:"dnsextendedlogging,omitempty"`
	Dnserrorlogging        bool   `json:"dnserrorlogging,omitempty"`
	Cacherecords           bool   `json:"cacherecords,omitempty"`
	Cachenegativeresponses bool   `json:"cachenegativeresponses,omitempty"`
	Dropmultiqueryrequest  bool   `json:"dropmultiqueryrequest,omitempty"`
	Cacheecsresponses      bool   `json:"cacheecsresponses,omitempty"`
}

type update_dnsprofile struct {
	Dnsprofilename         string `json:"dnsprofilename"`
	Dnsquerylogging        string `json:"dnsquerylogging,omitempty"`
	Dnsanswerseclogging    string `json:"dnsanswerseclogging,omitempty"`
	Dnsextendedlogging     string `json:"dnsextendedlogging,omitempty"`
	Dnserrorlogging        string `json:"dnserrorlogging,omitempty"`
	Cacherecords           string `json:"cacherecords,omitempty"`
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	Dropmultiqueryrequest  string `json:"dropmultiqueryrequest,omitempty"`
	Cacheecsresponses      string `json:"cacheecsresponses,omitempty"`
}

type rename_dnsprofile struct {
	Name    string `json:"dnsprofilename"`
	Newname string `json:"newname"`
}

type add_dnsprofile_payload struct {
	Resource Dnsprofile `json:"dnsprofile"`
}

type rename_dnsprofile_payload struct {
	Rename rename_dnsprofile `json:"dnsprofile"`
}

type unset_dnsprofile_payload struct {
	Unset DnsprofileUnset `json:"dnsprofile"`
}

type update_dnsprofile_payload struct {
	Update update_dnsprofile `json:"dnsprofile"`
}

type get_dnsprofile_result struct {
	Results []Dnsprofile `json:"dnsprofile"`
}

type count_dnsprofile_result struct {
	Results []Count `json:"dnsprofile"`
}

func (c *NitroClient) AddDnsprofile(resource Dnsprofile) error {
	payload := add_dnsprofile_payload{
		resource,
	}

	return c.post("dnsprofile", "", nil, payload)
}

func (c *NitroClient) RenameDnsprofile(name string, newName string) error {
	payload := rename_dnsprofile_payload{
		rename_dnsprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("dnsprofile", "", qs, payload)
}

func (c *NitroClient) CountDnsprofile() (int, error) {
	var results count_dnsprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dnsprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsDnsprofile(key string) (bool, error) {
	var results count_dnsprofile_result

	id, qs := dnsprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("dnsprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListDnsprofile() ([]Dnsprofile, error) {
	results := get_dnsprofile_result{}

	if err := c.get("dnsprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetDnsprofile(key string) (*Dnsprofile, error) {
	var results get_dnsprofile_result

	id, qs := dnsprofile_key_to_id_args(key)

	if err := c.get("dnsprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dnsprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dnsprofile element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteDnsprofile(key string) error {
	id, qs := dnsprofile_key_to_id_args(key)

	return c.delete("dnsprofile", id, qs)
}

func (c *NitroClient) UnsetDnsprofile(unset DnsprofileUnset) error {
	payload := unset_dnsprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("dnsprofile", "", qs, payload)
}

func (c *NitroClient) UpdateDnsprofile(resource Dnsprofile) error {
	payload := update_dnsprofile_payload{
		update_dnsprofile{
			resource.Dnsprofilename,
			resource.Dnsquerylogging,
			resource.Dnsanswerseclogging,
			resource.Dnsextendedlogging,
			resource.Dnserrorlogging,
			resource.Cacherecords,
			resource.Cachenegativeresponses,
			resource.Dropmultiqueryrequest,
			resource.Cacheecsresponses,
		},
	}

	return c.put("dnsprofile", "", nil, payload)
}
