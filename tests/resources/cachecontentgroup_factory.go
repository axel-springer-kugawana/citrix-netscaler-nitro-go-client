package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_cachecontentgroup(t *testing.T, client *nitro.NitroClient) (*nitro.Cachecontentgroup, func()) {
	// resource := nitro.Cachecontentgroup{
	// 	Name: "cachecontentgroup",
	// 	//Absexpiry:              []string{""},
	// 	//Absexpirygmt:           []string{""},
	// 	Alwaysevalpolicies: "NO",
	// 	Cachecontrol:       "no-store",
	// 	Expireatlastbyte:   "NO",
	// 	Flashcache:         "NO",
	// 	Heurexpiryparam:    0,
	// 	Hitparams:          []string{"Penn", "Teller"},
	// 	//Hitselector:           "string",
	// 	//Ignoreparamvaluecase:  "YES",
	// 	Ignorereloadreq:      "YES",
	// 	Ignorereqcachinghdrs: "YES",
	// 	Insertage:            "NO",
	// 	Insertetag:           "YES",
	// 	Insertvia:            "YES",
	// 	//Invalparams:           []string{"test"},
	// 	Invalrestrictedtohost: "NO",
	// 	//Invalselector:         "string",
	// 	Lazydnsresolve:     "YES",
	// 	Matchcookies:       "YES",
	// 	Maxressize:         80,
	// 	Memlimit:           65536,
	// 	Minhits:            0,
	// 	Minressize:         0,
	// 	Persistha:          "NO",
	// 	Pinned:             "NO",
	// 	Polleverytime:      "NO",
	// 	Prefetch:           "YES",
	// 	Prefetchmaxpending: 2,
	// 	//Prefetchperiod:         2,
	// 	//Prefetchperiodmillisec: 20,
	// 	Quickabortsize: 4194303,
	// 	//Relexpiry:              2,
	// 	//Relexpirymillisec: 2,
	// 	Removecookies:    "YES",
	// 	Type:             "HTTP",
	// 	Weaknegrelexpiry: 0,
	// 	//Weakposrelexpiry: 20,
	// }

	return nil, func() {
		//client.DeleteCachecontentgroup(resource.ToKey())
	}
}
