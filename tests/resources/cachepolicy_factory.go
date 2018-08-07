package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func Setup_cachepolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Cachepolicy, func()) {
	//resource := nitro.Cachepolicy {
	//Policyname: "cachepolicy",
	//Action: (CACHE|NOCACHE|MAY_CACHE|MAY_NOCACHE|INVAL),
	//Invalgroups: cachecontentgroup.name[],
	//Invalobjects: cachecontentgroup.name[],
	//Rule: string,
	//Storeingroup: cachecontentgroup.name,
	//Undefaction: (NOCACHE|RESET),
	//}

	return nil, func() {
	}
}
