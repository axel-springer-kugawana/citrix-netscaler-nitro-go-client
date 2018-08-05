package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Cachecontentgroup struct {
	Name                   string   `json:"name"`
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

func cachecontentgroup_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	return key, nil
}

type CachecontentgroupUnset struct {
	Name                   string `json:"name"`
	Weakposrelexpiry       bool   `json:"weakposrelexpiry,omitempty"`
	Heurexpiryparam        bool   `json:"heurexpiryparam,omitempty"`
	Relexpiry              bool   `json:"relexpiry,omitempty"`
	Relexpirymillisec      bool   `json:"relexpirymillisec,omitempty"`
	Absexpiry              bool   `json:"absexpiry,omitempty"`
	Absexpirygmt           bool   `json:"absexpirygmt,omitempty"`
	Weaknegrelexpiry       bool   `json:"weaknegrelexpiry,omitempty"`
	Hitparams              bool   `json:"hitparams,omitempty"`
	Invalparams            bool   `json:"invalparams,omitempty"`
	Ignoreparamvaluecase   bool   `json:"ignoreparamvaluecase,omitempty"`
	Matchcookies           bool   `json:"matchcookies,omitempty"`
	Invalrestrictedtohost  bool   `json:"invalrestrictedtohost,omitempty"`
	Polleverytime          bool   `json:"polleverytime,omitempty"`
	Ignorereloadreq        bool   `json:"ignorereloadreq,omitempty"`
	Removecookies          bool   `json:"removecookies,omitempty"`
	Prefetch               bool   `json:"prefetch,omitempty"`
	Prefetchperiod         bool   `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec bool   `json:"prefetchperiodmillisec,omitempty"`
	Prefetchmaxpending     bool   `json:"prefetchmaxpending,omitempty"`
	Flashcache             bool   `json:"flashcache,omitempty"`
	Expireatlastbyte       bool   `json:"expireatlastbyte,omitempty"`
	Insertvia              bool   `json:"insertvia,omitempty"`
	Insertage              bool   `json:"insertage,omitempty"`
	Insertetag             bool   `json:"insertetag,omitempty"`
	Cachecontrol           bool   `json:"cachecontrol,omitempty"`
	Quickabortsize         bool   `json:"quickabortsize,omitempty"`
	Minressize             bool   `json:"minressize,omitempty"`
	Maxressize             bool   `json:"maxressize,omitempty"`
	Memlimit               bool   `json:"memlimit,omitempty"`
	Ignorereqcachinghdrs   bool   `json:"ignorereqcachinghdrs,omitempty"`
	Minhits                bool   `json:"minhits,omitempty"`
	Alwaysevalpolicies     bool   `json:"alwaysevalpolicies,omitempty"`
	Persistha              bool   `json:"persistha,omitempty"`
	Pinned                 bool   `json:"pinned,omitempty"`
	Lazydnsresolve         bool   `json:"lazydnsresolve,omitempty"`
	Hitselector            bool   `json:"hitselector,omitempty"`
	Invalselector          bool   `json:"invalselector,omitempty"`
}

type update_cachecontentgroup struct {
	Name                   string   `json:"name"`
	Weakposrelexpiry       int      `json:"weakposrelexpiry,string,omitempty"`
	Heurexpiryparam        int      `json:"heurexpiryparam,string,omitempty"`
	Relexpiry              int      `json:"relexpiry,string,omitempty"`
	Relexpirymillisec      int      `json:"relexpirymillisec,string,omitempty"`
	Absexpiry              []string `json:"absexpiry,omitempty"`
	Absexpirygmt           []string `json:"absexpirygmt,omitempty"`
	Weaknegrelexpiry       int      `json:"weaknegrelexpiry,string,omitempty"`
	Hitparams              []string `json:"hitparams,omitempty"`
	Invalparams            []string `json:"invalparams,omitempty"`
	Ignoreparamvaluecase   string   `json:"ignoreparamvaluecase,omitempty"`
	Matchcookies           string   `json:"matchcookies,omitempty"`
	Invalrestrictedtohost  string   `json:"invalrestrictedtohost,omitempty"`
	Polleverytime          string   `json:"polleverytime,omitempty"`
	Ignorereloadreq        string   `json:"ignorereloadreq,omitempty"`
	Removecookies          string   `json:"removecookies,omitempty"`
	Prefetch               string   `json:"prefetch,omitempty"`
	Prefetchperiod         int      `json:"prefetchperiod,string,omitempty"`
	Prefetchperiodmillisec int      `json:"prefetchperiodmillisec,string,omitempty"`
	Prefetchmaxpending     int      `json:"prefetchmaxpending,string,omitempty"`
	Flashcache             string   `json:"flashcache,omitempty"`
	Expireatlastbyte       string   `json:"expireatlastbyte,omitempty"`
	Insertvia              string   `json:"insertvia,omitempty"`
	Insertage              string   `json:"insertage,omitempty"`
	Insertetag             string   `json:"insertetag,omitempty"`
	Cachecontrol           string   `json:"cachecontrol,omitempty"`
	Quickabortsize         int      `json:"quickabortsize,string,omitempty"`
	Minressize             int      `json:"minressize,string,omitempty"`
	Maxressize             int      `json:"maxressize,string,omitempty"`
	Memlimit               int      `json:"memlimit,string,omitempty"`
	Ignorereqcachinghdrs   string   `json:"ignorereqcachinghdrs,omitempty"`
	Minhits                int      `json:"minhits,omitempty"`
	Alwaysevalpolicies     string   `json:"alwaysevalpolicies,omitempty"`
	Persistha              string   `json:"persistha,omitempty"`
	Pinned                 string   `json:"pinned,omitempty"`
	Lazydnsresolve         string   `json:"lazydnsresolve,omitempty"`
	Hitselector            string   `json:"hitselector,omitempty"`
	Invalselector          string   `json:"invalselector,omitempty"`
}

type rename_cachecontentgroup struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_cachecontentgroup_payload struct {
	Resource Cachecontentgroup `json:"cachecontentgroup"`
}

type rename_cachecontentgroup_payload struct {
	Rename rename_cachecontentgroup `json:"cachecontentgroup"`
}

type unset_cachecontentgroup_payload struct {
	Unset CachecontentgroupUnset `json:"cachecontentgroup"`
}

type update_cachecontentgroup_payload struct {
	Update update_cachecontentgroup `json:"cachecontentgroup"`
}

type get_cachecontentgroup_result struct {
	Results []Cachecontentgroup `json:"cachecontentgroup"`
}

type count_cachecontentgroup_result struct {
	Results []Count `json:"cachecontentgroup"`
}

func (c *NitroClient) AddCachecontentgroup(resource Cachecontentgroup) error {
	payload := add_cachecontentgroup_payload{
		resource,
	}

	return c.post("cachecontentgroup", "", nil, payload)
}

func (c *NitroClient) RenameCachecontentgroup(name string, newName string) error {
	payload := rename_cachecontentgroup_payload{
		rename_cachecontentgroup{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("cachecontentgroup", "", qs, payload)
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

func (c *NitroClient) ExistsCachecontentgroup(key string) (bool, error) {
	var results count_cachecontentgroup_result

	id, qs := cachecontentgroup_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("cachecontentgroup", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCachecontentgroup() ([]Cachecontentgroup, error) {
	results := get_cachecontentgroup_result{}

	if err := c.get("cachecontentgroup", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCachecontentgroup(key string) (*Cachecontentgroup, error) {
	var results get_cachecontentgroup_result

	id, qs := cachecontentgroup_key_to_id_args(key)

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

func (c *NitroClient) DeleteCachecontentgroup(key string) error {
	id, qs := cachecontentgroup_key_to_id_args(key)

	return c.delete("cachecontentgroup", id, qs)
}

func (c *NitroClient) UnsetCachecontentgroup(unset CachecontentgroupUnset) error {
	payload := unset_cachecontentgroup_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("cachecontentgroup", "", qs, payload)
}

func (c *NitroClient) UpdateCachecontentgroup(resource Cachecontentgroup) error {
	payload := update_cachecontentgroup_payload{
		update_cachecontentgroup{
			resource.Name,
			resource.Weakposrelexpiry,
			resource.Heurexpiryparam,
			resource.Relexpiry,
			resource.Relexpirymillisec,
			resource.Absexpiry,
			resource.Absexpirygmt,
			resource.Weaknegrelexpiry,
			resource.Hitparams,
			resource.Invalparams,
			resource.Ignoreparamvaluecase,
			resource.Matchcookies,
			resource.Invalrestrictedtohost,
			resource.Polleverytime,
			resource.Ignorereloadreq,
			resource.Removecookies,
			resource.Prefetch,
			resource.Prefetchperiod,
			resource.Prefetchperiodmillisec,
			resource.Prefetchmaxpending,
			resource.Flashcache,
			resource.Expireatlastbyte,
			resource.Insertvia,
			resource.Insertage,
			resource.Insertetag,
			resource.Cachecontrol,
			resource.Quickabortsize,
			resource.Minressize,
			resource.Maxressize,
			resource.Memlimit,
			resource.Ignorereqcachinghdrs,
			resource.Minhits,
			resource.Alwaysevalpolicies,
			resource.Persistha,
			resource.Pinned,
			resource.Lazydnsresolve,
			resource.Hitselector,
			resource.Invalselector,
		},
	}

	return c.put("cachecontentgroup", "", nil, payload)
}
