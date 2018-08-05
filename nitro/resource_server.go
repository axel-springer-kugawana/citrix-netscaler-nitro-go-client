package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Server struct {
	Name               string `json:"name"`
	State              string `json:"state,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Ipv6address        string `json:"ipv6address,omitempty"`
	Td                 int    `json:"td,string,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

func server_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type ServerUnset struct {
	Name               string `json:"name"`
	Comment            bool   `json:"comment,string,omitempty"`
	Domainresolveretry bool   `json:"domainresolveretry,string,omitempty"`
	Ipaddress          bool   `json:"ipaddress,string,omitempty"`
	Translationip      bool   `json:"translationip,string,omitempty"`
	Translationmask    bool   `json:"translationmask,string,omitempty"`
}

type update_server struct {
	Name               string `json:"name"`
	Comment            string `json:"comment,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Translationip      string `json:"translationip,omitempty"`
	Translationmask    string `json:"translationmask,omitempty"`
}

type rename_server struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_server_payload struct {
	Resource Server `json:"server"`
}

type rename_server_payload struct {
	Rename rename_server `json:"server"`
}

type state_server struct {
	Key string `json:"name"`
}

type state_server_payload struct {
	Sate state_server `json:"server"`
}

type unset_server_payload struct {
	Unset ServerUnset `json:"server"`
}

type update_server_payload struct {
	Update update_server `json:"server"`
}

type get_server_result struct {
	Results []Server `json:"server"`
}

type count_server_result struct {
	Results []Count `json:"server"`
}

func (c *NitroClient) AddServer(resource Server) error {
	payload := add_server_payload{
		resource,
	}

	return c.post("server", "", nil, payload)
}

func (c *NitroClient) RenameServer(name string, newName string) error {
	payload := rename_server_payload{
		rename_server{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("server", "", qs, payload)
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

func (c *NitroClient) ExistsServer(key string) (bool, error) {
	var results count_server_result

	id, qs := server_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("server", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListServer() ([]Server, error) {
	results := get_server_result{}

	if err := c.get("server", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetServer(key string) (*Server, error) {
	var results get_server_result

	id, qs := server_key_to_id_args(key)

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

func (c *NitroClient) DeleteServer(key string) error {
	id, qs := server_key_to_id_args(key)

	return c.delete("server", id, qs)
}

func (c *NitroClient) UnsetServer(unset ServerUnset) error {
	payload := unset_server_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("server", "", qs, payload)
}

func (c *NitroClient) UpdateServer(resource Server) error {
	payload := update_server_payload{
		update_server{
			resource.Name,
			resource.Comment,
			resource.Domainresolveretry,
			resource.Ipaddress,
			resource.Translationip,
			resource.Translationmask,
		},
	}

	return c.put("server", "", nil, payload)
}

func (c *NitroClient) EnableServer(key string) error {
	payload := state_server_payload{
		state_server{
			key,
		},
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("server", "", qs, payload)
}

func (c *NitroClient) DisableServer(key string) error {
	payload := state_server_payload{
		state_server{
			key,
		},
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("server", "", qs, payload)
}
