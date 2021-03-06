package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Server struct {
	Comment            string `json:"comment,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Ipv6address        string `json:"ipv6address,omitempty"`
	Name               string `json:"name,omitempty"`
	State              string `json:"state,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

type ServerKey struct {
	Name string
}

func (resource Server) ToKey() ServerKey {
	key := ServerKey{
		resource.Name,
	}

	return key
}

func (key ServerKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key ServerKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key ServerKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_server_payload struct {
	Resource Server `json:"server"`
}

func (c *NitroClient) AddServer(resource Server) error {
	payload := add_server_payload{
		resource,
	}

	return c.post("server", "", nil, payload)
}

//      LIST

type list_server_result struct {
	Results []Server `json:"server"`
}

func (c *NitroClient) ListServer() ([]Server, error) {
	results := list_server_result{}

	if err := c.get("server", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_server_result struct {
	Results []Server `json:"server"`
}

func (c *NitroClient) GetServer(key ServerKey) (*Server, error) {
	var results get_server_result

	id, qs := key.to_id_args()

	if err := c.get("server", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one server element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("server element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_server_result struct {
	Results []Count `json:"server"`
}

func (c *NitroClient) CountServer() (int, error) {
	var results count_server_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("server", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsServer(key ServerKey) (bool, error) {
	var results count_server_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("server", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteServer(key ServerKey) error {
	id, qs := key.to_id_args()

	return c.delete("server", id, qs)
}

//      UPDATE

type ServerUpdate struct {
	Name               string `json:"name,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

func (resource Server) ToUpdate() ServerUpdate {
	update := ServerUpdate{
		resource.Name,
		resource.Comment,
		resource.Domainresolveretry,
		resource.Ipaddress,
		resource.Translationip,
		resource.Translationmask,
	}

	return update
}

type update_server_payload struct {
	Update ServerUpdate `json:"server"`
}

func (c *NitroClient) UpdateServer(update ServerUpdate) error {
	payload := update_server_payload{
		update,
	}

	return c.put("server", "", nil, payload)
}

//      UNSET

type ServerUnset struct {
	Name               string `json:"name,omitempty"`
	Comment            bool   `json:"comment,omitempty"`
	Domainresolveretry bool   `json:"domainresolveretry,omitempty"`
	Ipaddress          bool   `json:"ipaddress,omitempty"`
	Translationip      bool   `json:"translationip,omitempty"`
	Translationmask    bool   `json:"translationmask,omitempty"`
}

func (resource Server) ToUnset() ServerUnset {
	unset := ServerUnset{
		resource.Name,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_server_payload struct {
	Unset ServerUnset `json:"server"`
}

func (c *NitroClient) UnsetServer(unset ServerUnset) error {
	payload := unset_server_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("server", "", qs, payload)
}

//      RENAME
//      TODO

//      ENABLE

type state_server_payload struct {
	State ServerKey `json:"server"`
}

func (c *NitroClient) EnableServer(key ServerKey) error {
	payload := state_server_payload{
		key,
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("server", "", qs, payload)
}

//      DISABLE

func (c *NitroClient) DisableServer(key ServerKey) error {
	payload := state_server_payload{
		key,
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("server", "", qs, payload)
}
