package nitro

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

type CachecontentgroupKey struct {
	Name string
}

type cachecontentgroup_update struct {
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

type cachecontentgroup_payload struct {
	cachecontentgroup interface{}
}

func cachecontentgroup_key_to_args(key CachecontentgroupKey) string {
	result := ""

	return result
}

func (c *NitroClient) DeleteCachecontentgroup(key CachecontentgroupKey) error {
	return c.deleteResourceWithArgs("cachecontentgroup", key.Name, cachecontentgroup_key_to_args(key))
}

func (c *NitroClient) GetCachecontentgroup(key CachecontentgroupKey) (*Cachecontentgroup, error) {
	var results struct {
		Cachecontentgroup []Cachecontentgroup
	}

	if err := c.getResourceWithArgs("cachecontentgroup", key.Name, cachecontentgroup_key_to_args(key), &results); err != nil || len(results.Cachecontentgroup) != 1 {
		return nil, err
	}

	return &results.Cachecontentgroup[0], nil
}

func (c *NitroClient) ListCachecontentgroup() ([]Cachecontentgroup, error) {
	var results struct {
		Cachecontentgroup []Cachecontentgroup
	}

	if err := c.listResources("cachecontentgroup", &results); err != nil {
		return nil, err
	}

	return results.Cachecontentgroup, nil
}

func (c *NitroClient) AddCachecontentgroup(resource Cachecontentgroup) error {
	return c.addResource("cachecontentgroup", resource)
}

func (c *NitroClient) RenameCachecontentgroup(name string, newName string) error {
	return c.renameResource("cachecontentgroup", "name", name, newName)
}

func (c *NitroClient) UnsetCachecontentgroup(name string, fields ...string) error {
	return c.unsetResource("cachecontentgroup", "name", name, fields)
}

func (c *NitroClient) UpdateCachecontentgroup(resource Cachecontentgroup) error {
	update := cachecontentgroup_update{
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
	}

	return c.Put("cachecontentgroup", update)
}
