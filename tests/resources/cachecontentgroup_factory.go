package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_cachecontentgroup(t *testing.T, client *nitro.NitroClient) (*nitro.Cachecontentgroup, func()) {
	//resource := nitro.Cachecontentgroup {
	//Name: "cachecontentgroup",
	//Absexpiry: string[],
	//Absexpirygmt: string[],
	//Alwaysevalpolicies: (YES|NO),
	//Cachecontrol: string,
	//Expireatlastbyte: (YES|NO),
	//Flashcache: (YES|NO),
	//Heurexpiryparam: double,
	//Hitparams: string[],
	//Hitselector: string,
	//Ignoreparamvaluecase: (YES|NO),
	//Ignorereloadreq: (YES|NO),
	//Ignorereqcachinghdrs: (YES|NO),
	//Insertage: (YES|NO),
	//Insertetag: (YES|NO),
	//Insertvia: (YES|NO),
	//Invalparams: string[],
	//Invalrestrictedtohost: (YES|NO),
	//Invalselector: string,
	//Lazydnsresolve: (YES|NO),
	//Matchcookies: (YES|NO),
	//Maxressize: double,
	//Memlimit: double,
	//Minhits: int,
	//Minressize: double,
	//Persistha: (YES|NO),
	//Pinned: (YES|NO),
	//Polleverytime: (YES|NO),
	//Prefetch: (YES|NO),
	//Prefetchmaxpending: double,
	//Prefetchperiod: double,
	//Prefetchperiodmillisec: double,
	//Quickabortsize: double,
	//Relexpiry: double,
	//Relexpirymillisec: double,
	//Removecookies: (YES|NO),
	//Type: (HTTP|MYSQL|MSSQL),
	//Weaknegrelexpiry: double,
	//Weakposrelexpiry: double,
	//}

	return nil, func() {
	}
}
