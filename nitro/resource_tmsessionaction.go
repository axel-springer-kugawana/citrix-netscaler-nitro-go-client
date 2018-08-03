package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmsessionaction struct {
	Name                       string `json:"name"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,string,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,string,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
}

type TmsessionactionKey struct {
	Name string `json:"name"`
}

type TmsessionactionUnset struct {
	Name                       string `json:"name"`
	Sesstimeout                bool   `json:"sesstimeout,string,omitempty"`
	Defaultauthorizationaction bool   `json:"defaultauthorizationaction,string,omitempty"`
	Sso                        bool   `json:"sso,string,omitempty"`
	Ssocredential              bool   `json:"ssocredential,string,omitempty"`
	Ssodomain                  bool   `json:"ssodomain,string,omitempty"`
	Kcdaccount                 bool   `json:"kcdaccount,string,omitempty"`
	Httponlycookie             bool   `json:"httponlycookie,string,omitempty"`
	Persistentcookie           bool   `json:"persistentcookie,string,omitempty"`
	Persistentcookievalidity   bool   `json:"persistentcookievalidity,string,omitempty"`
	Homepage                   bool   `json:"homepage,string,omitempty"`
}

type update_tmsessionaction struct {
	Name                       string `json:"name"`
	Sesstimeout                int    `json:"sesstimeout,string,omitempty"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,string,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
}

type rename_tmsessionaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_tmsessionaction_payload struct {
	Resource Tmsessionaction `json:"tmsessionaction"`
}

type rename_tmsessionaction_payload struct {
	Rename rename_tmsessionaction `json:"tmsessionaction"`
}

type unset_tmsessionaction_payload struct {
	Unset TmsessionactionUnset `json:"tmsessionaction"`
}

type update_tmsessionaction_payload struct {
	Update update_tmsessionaction `json:"tmsessionaction"`
}

type get_tmsessionaction_result struct {
	Results []Tmsessionaction `json:"tmsessionaction"`
}

type count_tmsessionaction_result struct {
	Results []Count `json:"tmsessionaction"`
}

func tmsessionaction_key_to_id_args(key TmsessionactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddTmsessionaction(resource Tmsessionaction) error {
	payload := add_tmsessionaction_payload{
		resource,
	}

	return c.post("tmsessionaction", "", nil, payload)
}

func (c *NitroClient) RenameTmsessionaction(name string, newName string) error {
	payload := rename_tmsessionaction_payload{
		rename_tmsessionaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("tmsessionaction", "", qs, payload)
}

func (c *NitroClient) CountTmsessionaction() (int, error) {
	var results count_tmsessionaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("tmsessionaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsTmsessionaction(key TmsessionactionKey) (bool, error) {
	var results count_tmsessionaction_result

	id, qs := tmsessionaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("tmsessionaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListTmsessionaction() ([]Tmsessionaction, error) {
	var results get_tmsessionaction_result

	if err := c.get("tmsessionaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetTmsessionaction(key TmsessionactionKey) (*Tmsessionaction, error) {
	var results get_tmsessionaction_result

	id, qs := tmsessionaction_key_to_id_args(key)

	if err := c.get("tmsessionaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmsessionaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("tmsessionaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteTmsessionaction(key TmsessionactionKey) error {
	id, qs := tmsessionaction_key_to_id_args(key)

	return c.delete("tmsessionaction", id, qs)
}

func (c *NitroClient) UnsetTmsessionaction(unset TmsessionactionUnset) error {
	payload := unset_tmsessionaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("tmsessionaction", "", qs, payload)
}

func (c *NitroClient) UpdateTmsessionaction(resource Tmsessionaction) error {
	payload := update_tmsessionaction_payload{
		update_tmsessionaction{
			resource.Name,
			resource.Sesstimeout,
			resource.Defaultauthorizationaction,
			resource.Sso,
			resource.Ssocredential,
			resource.Ssodomain,
			resource.Kcdaccount,
			resource.Httponlycookie,
			resource.Persistentcookie,
			resource.Persistentcookievalidity,
			resource.Homepage,
		},
	}

	return c.put("tmsessionaction", "", nil, payload)
}
