package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachepolicy struct {
	Action       string   `json:"action,omitempty"`
	Invalgroups  []string `json:"invalgroups,omitempty"`
	Invalobjects []string `json:"invalobjects,omitempty"`
	Policyname   string   `json:"policyname,omitempty"`
	Rule         string   `json:"rule,omitempty"`
	Storeingroup string   `json:"storeingroup,omitempty"`
	Undefaction  string   `json:"undefaction,omitempty"`
}

type CachepolicyKey struct {
	Policyname string
}

func (resource Cachepolicy) ToKey() CachepolicyKey {
	key := CachepolicyKey{
		resource.Policyname,
	}

	return key
}

func (key CachepolicyKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa

	var id string
	var args []string

	id = key.Policyname

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CachepolicyKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CachepolicyKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cachepolicy_payload struct {
	Resource Cachepolicy `json:"cachepolicy"`
}

func (c *NitroClient) AddCachepolicy(resource Cachepolicy) error {
	payload := add_cachepolicy_payload{
		resource,
	}

	return c.post("cachepolicy", "", nil, payload)
}

//      LIST

type list_cachepolicy_result struct {
	Results []Cachepolicy `json:"cachepolicy"`
}

func (c *NitroClient) ListCachepolicy() ([]Cachepolicy, error) {
	results := list_cachepolicy_result{}

	if err := c.get("cachepolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cachepolicy_result struct {
	Results []Cachepolicy `json:"cachepolicy"`
}

func (c *NitroClient) GetCachepolicy(key CachepolicyKey) (*Cachepolicy, error) {
	var results get_cachepolicy_result

	id, qs := key.to_id_args()

	if err := c.get("cachepolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cachepolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cachepolicy element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cachepolicy_result struct {
	Results []Count `json:"cachepolicy"`
}

func (c *NitroClient) CountCachepolicy() (int, error) {
	var results count_cachepolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cachepolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCachepolicy(key CachepolicyKey) (bool, error) {
	var results count_cachepolicy_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cachepolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCachepolicy(key CachepolicyKey) error {
	id, qs := key.to_id_args()

	return c.delete("cachepolicy", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
