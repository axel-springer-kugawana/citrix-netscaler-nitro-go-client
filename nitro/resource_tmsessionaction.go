package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Tmsessionaction struct {
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Name                       string `json:"name,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,string,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,string,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
}

type TmsessionactionKey struct {
	Name string
}

func (resource Tmsessionaction) ToKey() TmsessionactionKey {
	key := TmsessionactionKey{
		resource.Name,
	}

	return key
}

func (key TmsessionactionKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Name

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key TmsessionactionKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key TmsessionactionKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_tmsessionaction_payload struct {
	Resource Tmsessionaction `json:"tmsessionaction"`
}

func (c *NitroClient) AddTmsessionaction(resource Tmsessionaction) error {
	payload := add_tmsessionaction_payload{
		resource,
	}

	return c.post("tmsessionaction", "", nil, payload)
}

//      LIST

type list_tmsessionaction_result struct {
	Results []Tmsessionaction `json:"tmsessionaction"`
}

func (c *NitroClient) ListTmsessionaction() ([]Tmsessionaction, error) {
	results := list_tmsessionaction_result{}

	if err := c.get("tmsessionaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_tmsessionaction_result struct {
	Results []Tmsessionaction `json:"tmsessionaction"`
}

func (c *NitroClient) GetTmsessionaction(key TmsessionactionKey) (*Tmsessionaction, error) {
	var results get_tmsessionaction_result

	id, qs := key.to_id_args()

	if err := c.get("tmsessionaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one tmsessionaction element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("tmsessionaction element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_tmsessionaction_result struct {
	Results []Count `json:"tmsessionaction"`
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

//      EXISTS

func (c *NitroClient) ExistsTmsessionaction(key TmsessionactionKey) (bool, error) {
	var results count_tmsessionaction_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("tmsessionaction", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteTmsessionaction(key TmsessionactionKey) error {
	id, qs := key.to_id_args()

	return c.delete("tmsessionaction", id, qs)
}

//      UPDATE

type TmsessionactionUpdate struct {
	Name                       string `json:"name,omitempty"`
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

func (resource Tmsessionaction) ToUpdate() TmsessionactionUpdate {
	update := TmsessionactionUpdate{
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
	}

	return update
}

type update_tmsessionaction_payload struct {
	Update TmsessionactionUpdate `json:"tmsessionaction"`
}

func (c *NitroClient) UpdateTmsessionaction(update TmsessionactionUpdate) error {
	payload := update_tmsessionaction_payload{
		update,
	}

	return c.put("tmsessionaction", "", nil, payload)
}

//      UNSET

type TmsessionactionUnset struct {
	Name                       string `json:"name,omitempty"`
	Sesstimeout                bool   `json:"sesstimeout,omitempty"`
	Defaultauthorizationaction bool   `json:"defaultauthorizationaction,omitempty"`
	Sso                        bool   `json:"sso,omitempty"`
	Ssocredential              bool   `json:"ssocredential,omitempty"`
	Ssodomain                  bool   `json:"ssodomain,omitempty"`
	Kcdaccount                 bool   `json:"kcdaccount,omitempty"`
	Httponlycookie             bool   `json:"httponlycookie,omitempty"`
	Persistentcookie           bool   `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   bool   `json:"persistentcookievalidity,omitempty"`
	Homepage                   bool   `json:"homepage,omitempty"`
}

func (resource Tmsessionaction) ToUnset() TmsessionactionUnset {
	unset := TmsessionactionUnset{
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

type unset_tmsessionaction_payload struct {
	Unset TmsessionactionUnset `json:"tmsessionaction"`
}

func (c *NitroClient) UnsetTmsessionaction(unset TmsessionactionUnset) error {
	payload := unset_tmsessionaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("tmsessionaction", "", qs, payload)
}

//      RENAME
//      TODO
