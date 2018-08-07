package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Authorizationpolicylabel struct {
	Labelname string `json:"labelname,omitempty"`
}

type AuthorizationpolicylabelKey struct {
	Labelname string
}

func (resource Authorizationpolicylabel) ToKey() AuthorizationpolicylabelKey {
	key := AuthorizationpolicylabelKey{
		resource.Labelname,
	}

	return key
}

func (key AuthorizationpolicylabelKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Labelname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key AuthorizationpolicylabelKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key AuthorizationpolicylabelKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_authorizationpolicylabel_payload struct {
	Resource Authorizationpolicylabel `json:"authorizationpolicylabel"`
}

func (c *NitroClient) AddAuthorizationpolicylabel(resource Authorizationpolicylabel) error {
	payload := add_authorizationpolicylabel_payload{
		resource,
	}

	return c.post("authorizationpolicylabel", "", nil, payload)
}

//      LIST

type list_authorizationpolicylabel_result struct {
	Results []Authorizationpolicylabel `json:"authorizationpolicylabel"`
}

func (c *NitroClient) ListAuthorizationpolicylabel() ([]Authorizationpolicylabel, error) {
	results := list_authorizationpolicylabel_result{}

	if err := c.get("authorizationpolicylabel", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_authorizationpolicylabel_result struct {
	Results []Authorizationpolicylabel `json:"authorizationpolicylabel"`
}

func (c *NitroClient) GetAuthorizationpolicylabel(key AuthorizationpolicylabelKey) (*Authorizationpolicylabel, error) {
	var results get_authorizationpolicylabel_result

	id, qs := key.to_id_args()

	if err := c.get("authorizationpolicylabel", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one authorizationpolicylabel element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("authorizationpolicylabel element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_authorizationpolicylabel_result struct {
	Results []Count `json:"authorizationpolicylabel"`
}

func (c *NitroClient) CountAuthorizationpolicylabel() (int, error) {
	var results count_authorizationpolicylabel_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("authorizationpolicylabel", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsAuthorizationpolicylabel(key AuthorizationpolicylabelKey) (bool, error) {
	var results count_authorizationpolicylabel_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("authorizationpolicylabel", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteAuthorizationpolicylabel(key AuthorizationpolicylabelKey) error {
	id, qs := key.to_id_args()

	return c.delete("authorizationpolicylabel", id, qs)
}

//      RENAME
//      TODO
