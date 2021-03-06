package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Appqoeaction struct {
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Customfile        string `json:"customfile,omitempty"`
	Delay             int    `json:"delay,string,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Maxconn           int    `json:"maxconn,string,omitempty"`
	Name              string `json:"name,omitempty"`
	Polqdepth         int    `json:"polqdepth,string,omitempty"`
	Priority          string `json:"priority,omitempty"`
	Priqdepth         int    `json:"priqdepth,string,omitempty"`
	Respondwith       string `json:"respondwith,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}

type AppqoeactionKey struct {
	Name string
}

func (resource Appqoeaction) ToKey() AppqoeactionKey {
	key := AppqoeactionKey{
		resource.Name,
	}

	return key
}

func (key AppqoeactionKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key AppqoeactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AppqoeactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_appqoeaction_payload struct {
	Resource Appqoeaction `json:"appqoeaction"`
}

func (c *NitroClient) AddAppqoeaction(resource Appqoeaction) error {
	payload := add_appqoeaction_payload{
		resource,
	}

	return c.post("appqoeaction", "", nil, payload)
}

//      LIST

type list_appqoeaction_result struct {
	Results []Appqoeaction `json:"appqoeaction"`
}

func (c *NitroClient) ListAppqoeaction() ([]Appqoeaction, error) {
	results := list_appqoeaction_result{}

	if err := c.get("appqoeaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_appqoeaction_result struct {
	Results []Appqoeaction `json:"appqoeaction"`
}

func (c *NitroClient) GetAppqoeaction(key AppqoeactionKey) (*Appqoeaction, error) {
	var results get_appqoeaction_result

	id, qs := key.to_id_args()

	if err := c.get("appqoeaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one appqoeaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("appqoeaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_appqoeaction_result struct {
	Results []Count `json:"appqoeaction"`
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

//      EXISTS

func (c *NitroClient) ExistsAppqoeaction(key AppqoeactionKey) (bool, error) {
	var results count_appqoeaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("appqoeaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAppqoeaction(key AppqoeactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("appqoeaction", id, qs)
}

//      UPDATE

type AppqoeactionUpdate struct {
	Name              string `json:"name,omitempty"`
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

func (resource Appqoeaction) ToUpdate() AppqoeactionUpdate {
	update := AppqoeactionUpdate{
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
	}

	return update
}

type update_appqoeaction_payload struct {
	Update AppqoeactionUpdate `json:"appqoeaction"`
}

func (c *NitroClient) UpdateAppqoeaction(update AppqoeactionUpdate) error {
	payload := update_appqoeaction_payload{
		update,
	}

	return c.put("appqoeaction", "", nil, payload)
}

//      UNSET

type AppqoeactionUnset struct {
	Name              string `json:"name,omitempty"`
	Priority          bool   `json:"priority,omitempty"`
	Altcontentsvcname bool   `json:"altcontentsvcname,omitempty"`
	Altcontentpath    bool   `json:"altcontentpath,omitempty"`
	Polqdepth         bool   `json:"polqdepth,omitempty"`
	Priqdepth         bool   `json:"priqdepth,omitempty"`
	Maxconn           bool   `json:"maxconn,omitempty"`
	Delay             bool   `json:"delay,omitempty"`
	Dostrigexpression bool   `json:"dostrigexpression,omitempty"`
	Dosaction         bool   `json:"dosaction,omitempty"`
	Tcpprofile        bool   `json:"tcpprofile,omitempty"`
}

func (resource Appqoeaction) ToUnset() AppqoeactionUnset {
	unset := AppqoeactionUnset{
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
	}

	return unset
}

type unset_appqoeaction_payload struct {
	Unset AppqoeactionUnset `json:"appqoeaction"`
}

func (c *NitroClient) UnsetAppqoeaction(unset AppqoeactionUnset) error {
	payload := unset_appqoeaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("appqoeaction", "", qs, payload)
}

//      RENAME
//      TODO
