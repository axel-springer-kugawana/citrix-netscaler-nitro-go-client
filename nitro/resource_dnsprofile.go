package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dnsprofile struct {
	Cacheecsresponses      string `json:"cacheecsresponses,omitempty"`
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	Cacherecords           string `json:"cacherecords,omitempty"`
	Dnsanswerseclogging    string `json:"dnsanswerseclogging,omitempty"`
	Dnserrorlogging        string `json:"dnserrorlogging,omitempty"`
	Dnsextendedlogging     string `json:"dnsextendedlogging,omitempty"`
	Dnsprofilename         string `json:"dnsprofilename,omitempty"`
	Dnsquerylogging        string `json:"dnsquerylogging,omitempty"`
	Dropmultiqueryrequest  string `json:"dropmultiqueryrequest,omitempty"`
}

type DnsprofileKey struct {
	Dnsprofilename string
}

func (resource Dnsprofile) ToKey() DnsprofileKey {
	key := DnsprofileKey{
		resource.Dnsprofilename,
	}

	return key
}

func (key DnsprofileKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Dnsprofilename

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key DnsprofileKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key DnsprofileKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_dnsprofile_payload struct {
	Resource Dnsprofile `json:"dnsprofile"`
}

func (c *NitroClient) AddDnsprofile(resource Dnsprofile) error {
	payload := add_dnsprofile_payload{
		resource,
	}

	return c.post("dnsprofile", "", nil, payload)
}

//      LIST

type list_dnsprofile_result struct {
	Results []Dnsprofile `json:"dnsprofile"`
}

func (c *NitroClient) ListDnsprofile() ([]Dnsprofile, error) {
	results := list_dnsprofile_result{}

	if err := c.get("dnsprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_dnsprofile_result struct {
	Results []Dnsprofile `json:"dnsprofile"`
}

func (c *NitroClient) GetDnsprofile(key DnsprofileKey) (*Dnsprofile, error) {
	var results get_dnsprofile_result

	id, qs := key.to_id_args()

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

//      COUNT

type count_dnsprofile_result struct {
	Results []Count `json:"dnsprofile"`
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

//      EXISTS

func (c *NitroClient) ExistsDnsprofile(key DnsprofileKey) (bool, error) {
	var results count_dnsprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("dnsprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteDnsprofile(key DnsprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("dnsprofile", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
