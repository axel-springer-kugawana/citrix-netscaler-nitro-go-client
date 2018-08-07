package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachecontentgroup struct {
	Absexpiry              []string `json:"absexpiry,omitempty"`
	Absexpirygmt           []string `json:"absexpirygmt,omitempty"`
	Alwaysevalpolicies     string   `json:"alwaysevalpolicies,omitempty"`
	Cachecontrol           string   `json:"cachecontrol,omitempty"`
	Expireatlastbyte       string   `json:"expireatlastbyte,omitempty"`
	Flashcache             string   `json:"flashcache,omitempty"`
	Heurexpiryparam        int      `json:"heurexpiryparam,string,omitempty"`
	Hitparams              []string `json:"hitparams,omitempty"`
	Hitselector            string   `json:"hitselector,omitempty"`
	Ignoreparamvaluecase   string   `json:"ignoreparamvaluecase,omitempty"`
	Ignorereloadreq        string   `json:"ignorereloadreq,omitempty"`
	Ignorereqcachinghdrs   string   `json:"ignorereqcachinghdrs,omitempty"`
	Insertage              string   `json:"insertage,omitempty"`
	Insertetag             string   `json:"insertetag,omitempty"`
	Insertvia              string   `json:"insertvia,omitempty"`
	Invalparams            []string `json:"invalparams,omitempty"`
	Invalrestrictedtohost  string   `json:"invalrestrictedtohost,omitempty"`
	Invalselector          string   `json:"invalselector,omitempty"`
	Lazydnsresolve         string   `json:"lazydnsresolve,omitempty"`
	Matchcookies           string   `json:"matchcookies,omitempty"`
	Maxressize             int      `json:"maxressize,string,omitempty"`
	Memlimit               int      `json:"memlimit,string,omitempty"`
	Minhits                int      `json:"minhits,omitempty"`
	Minressize             int      `json:"minressize,string,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Persistha              string   `json:"persistha,omitempty"`
	Pinned                 string   `json:"pinned,omitempty"`
	Polleverytime          string   `json:"polleverytime,omitempty"`
	Prefetch               string   `json:"prefetch,omitempty"`
	Prefetchmaxpending     int      `json:"prefetchmaxpending,string,omitempty"`
	Prefetchperiod         int      `json:"prefetchperiod,string,omitempty"`
	Prefetchperiodmillisec int      `json:"prefetchperiodmillisec,string,omitempty"`
	Quickabortsize         int      `json:"quickabortsize,string,omitempty"`
	Relexpiry              int      `json:"relexpiry,string,omitempty"`
	Relexpirymillisec      int      `json:"relexpirymillisec,string,omitempty"`
	Removecookies          string   `json:"removecookies,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Weaknegrelexpiry       int      `json:"weaknegrelexpiry,string,omitempty"`
	Weakposrelexpiry       int      `json:"weakposrelexpiry,string,omitempty"`
}

type CachecontentgroupKey struct {
	Name string
}

func (resource Cachecontentgroup) ToKey() CachecontentgroupKey {
	key := CachecontentgroupKey{
		resource.Name,
	}

	return key
}

func (key CachecontentgroupKey) to_id_params(qsKey string) (string, map[string]string) {
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

func (key CachecontentgroupKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CachecontentgroupKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_cachecontentgroup_payload struct {
	Resource Cachecontentgroup `json:"cachecontentgroup"`
}

func (c *NitroClient) AddCachecontentgroup(resource Cachecontentgroup) error {
	payload := add_cachecontentgroup_payload{
		resource,
	}

	return c.post("cachecontentgroup", "", nil, payload)
}

//      LIST

type list_cachecontentgroup_result struct {
	Results []Cachecontentgroup `json:"cachecontentgroup"`
}

func (c *NitroClient) ListCachecontentgroup() ([]Cachecontentgroup, error) {
	results := list_cachecontentgroup_result{}

	if err := c.get("cachecontentgroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_cachecontentgroup_result struct {
	Results []Cachecontentgroup `json:"cachecontentgroup"`
}

func (c *NitroClient) GetCachecontentgroup(key CachecontentgroupKey) (*Cachecontentgroup, error) {
	var results get_cachecontentgroup_result

	id, qs := key.to_id_args()

	if err := c.get("cachecontentgroup", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one cachecontentgroup element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("cachecontentgroup element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_cachecontentgroup_result struct {
	Results []Count `json:"cachecontentgroup"`
}

func (c *NitroClient) CountCachecontentgroup() (int, error) {
	var results count_cachecontentgroup_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("cachecontentgroup", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsCachecontentgroup(key CachecontentgroupKey) (bool, error) {
	var results count_cachecontentgroup_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("cachecontentgroup", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCachecontentgroup(key CachecontentgroupKey) error {
	id, qs := key.to_id_args()

	return c.delete("cachecontentgroup", id, qs)
}

//      UPDATE
//      TODO

//      UNSET
//      TODO

//      RENAME
//      TODO
