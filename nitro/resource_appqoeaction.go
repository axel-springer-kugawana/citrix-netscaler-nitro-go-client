package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Appqoeaction struct {
	Name              string `json:"name"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Customfile        string `json:"customfile,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Polqdepth         int    `json:"polqdepth,string,omitempty"`
	Priority          string `json:"priority,omitempty"`
	Priqdepth         int    `json:"priqdepth,string,omitempty"`
	Respondwith       string `json:"respondwith,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}

func appqoeaction_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type AppqoeactionUnset struct {
	Name              string `json:"name"`
	Priority          bool   `json:"priority,string,omitempty"`
	Altcontentsvcname bool   `json:"altcontentsvcname,string,omitempty"`
	Altcontentpath    bool   `json:"altcontentpath,string,omitempty"`
	Polqdepth         bool   `json:"polqdepth,string,omitempty"`
	Priqdepth         bool   `json:"priqdepth,string,omitempty"`
	Maxconn           bool   `json:"maxconn,string,omitempty"`
	Delay             bool   `json:"delay,string,omitempty"`
	Dostrigexpression bool   `json:"dostrigexpression,string,omitempty"`
	Dosaction         bool   `json:"dosaction,string,omitempty"`
	Tcpprofile        bool   `json:"tcpprofile,string,omitempty"`
}

type update_appqoeaction struct {
	Name              string `json:"name"`
	Priority          string `json:"priority,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Polqdepth         int    `json:"polqdepth,string,omitempty"`
	Priqdepth         int    `json:"priqdepth,string,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}

type rename_appqoeaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_appqoeaction_payload struct {
	Resource Appqoeaction `json:"appqoeaction"`
}

type rename_appqoeaction_payload struct {
	Rename rename_appqoeaction `json:"appqoeaction"`
}

type unset_appqoeaction_payload struct {
	Unset AppqoeactionUnset `json:"appqoeaction"`
}

type update_appqoeaction_payload struct {
	Update update_appqoeaction `json:"appqoeaction"`
}

type get_appqoeaction_result struct {
	Results []Appqoeaction `json:"appqoeaction"`
}

type count_appqoeaction_result struct {
	Results []Count `json:"appqoeaction"`
}

func (c *NitroClient) AddAppqoeaction(resource Appqoeaction) error {
	payload := add_appqoeaction_payload{
		resource,
	}

	return c.post("appqoeaction", "", nil, payload)
}

func (c *NitroClient) RenameAppqoeaction(name string, newName string) error {
	payload := rename_appqoeaction_payload{
		rename_appqoeaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("appqoeaction", "", qs, payload)
}

func (c *NitroClient) CountAppqoeaction() (int, error) {
	var results count_appqoeaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("appqoeaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsAppqoeaction(key string) (bool, error) {
	var results count_appqoeaction_result

	id, qs := appqoeaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("appqoeaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListAppqoeaction() ([]Appqoeaction, error) {
	results := get_appqoeaction_result{}

	if err := c.get("appqoeaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetAppqoeaction(key string) (*Appqoeaction, error) {
	var results get_appqoeaction_result

	id, qs := appqoeaction_key_to_id_args(key)

	if err := c.get("appqoeaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appqoeaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("appqoeaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteAppqoeaction(key string) error {
	id, qs := appqoeaction_key_to_id_args(key)

	return c.delete("appqoeaction", id, qs)
}

func (c *NitroClient) UnsetAppqoeaction(unset AppqoeactionUnset) error {
	payload := unset_appqoeaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("appqoeaction", "", qs, payload)
}

func (c *NitroClient) UpdateAppqoeaction(resource Appqoeaction) error {
	payload := update_appqoeaction_payload{
		update_appqoeaction{
			resource.Name,
			resource.Priority,
			resource.Altcontentsvcname,
			resource.Altcontentpath,
			resource.Polqdepth,
			resource.Priqdepth,
			resource.Maxconn,
			resource.Delay,
			resource.Dostrigexpression,
			resource.Dosaction,
			resource.Tcpprofile,
		},
	}

	return c.put("appqoeaction", "", nil, payload)
}
