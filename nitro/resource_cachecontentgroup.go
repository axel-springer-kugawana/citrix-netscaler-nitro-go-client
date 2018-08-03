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
	Weakposrelexpiry       bool   `json:"weakposrelexpiry,string,omitempty"`
	Heurexpiryparam        bool   `json:"heurexpiryparam,string,omitempty"`
	Relexpiry              bool   `json:"relexpiry,string,omitempty"`
	Relexpirymillisec      bool   `json:"relexpirymillisec,string,omitempty"`
	Absexpiry              bool   `json:"absexpiry,string,omitempty"`
	Absexpirygmt           bool   `json:"absexpirygmt,string,omitempty"`
	Weaknegrelexpiry       bool   `json:"weaknegrelexpiry,string,omitempty"`
	Hitparams              bool   `json:"hitparams,string,omitempty"`
	Invalparams            bool   `json:"invalparams,string,omitempty"`
	Ignoreparamvaluecase   bool   `json:"ignoreparamvaluecase,string,omitempty"`
	Matchcookies           bool   `json:"matchcookies,string,omitempty"`
	Invalrestrictedtohost  bool   `json:"invalrestrictedtohost,string,omitempty"`
	Polleverytime          bool   `json:"polleverytime,string,omitempty"`
	Ignorereloadreq        bool   `json:"ignorereloadreq,string,omitempty"`
	Removecookies          bool   `json:"removecookies,string,omitempty"`
	Prefetch               bool   `json:"prefetch,string,omitempty"`
	Prefetchperiod         bool   `json:"prefetchperiod,string,omitempty"`
	Prefetchperiodmillisec bool   `json:"prefetchperiodmillisec,string,omitempty"`
	Prefetchmaxpending     bool   `json:"prefetchmaxpending,string,omitempty"`
	Flashcache             bool   `json:"flashcache,string,omitempty"`
	Expireatlastbyte       bool   `json:"expireatlastbyte,string,omitempty"`
	Insertvia              bool   `json:"insertvia,string,omitempty"`
	Insertage              bool   `json:"insertage,string,omitempty"`
	Insertetag             bool   `json:"insertetag,string,omitempty"`
	Cachecontrol           bool   `json:"cachecontrol,string,omitempty"`
	Quickabortsize         bool   `json:"quickabortsize,string,omitempty"`
	Minressize             bool   `json:"minressize,string,omitempty"`
	Maxressize             bool   `json:"maxressize,string,omitempty"`
	Memlimit               bool   `json:"memlimit,string,omitempty"`
	Ignorereqcachinghdrs   bool   `json:"ignorereqcachinghdrs,string,omitempty"`
	Minhits                bool   `json:"minhits,string,omitempty"`
	Alwaysevalpolicies     bool   `json:"alwaysevalpolicies,string,omitempty"`
	Persistha              bool   `json:"persistha,string,omitempty"`
	Pinned                 bool   `json:"pinned,string,omitempty"`
	Lazydnsresolve         bool   `json:"lazydnsresolve,string,omitempty"`
	Hitselector            bool   `json:"hitselector,string,omitempty"`
	Invalselector          bool   `json:"invalselector,string,omitempty"`
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
			// TODO
			// return nil, fmt.Errorf("cachecontentgroup element not found")
			return nil, nil
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
