package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_nshttpprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Nshttpprofile, func()) {
	resource := nitro.Nshttpprofile{
		Name:        "nshttpprofile2",
		Adpttimeout: "DISABLED",
		//Altsvc:      "DISABLED",
		Apdexcltresptimethreshold: 0,
		//Clientiphdrexpr:           "",
		Cmponpush:     "DISABLED",
		Conmultiplex:  "ENABLED",
		Dropextracrlf: "ENABLED",
		Dropextradata: "DISABLED",
		Dropinvalreqs: "ENABLED",
		Http2:         "DISABLED",
		//Http2direct:               "DISABLED",
		Http2headertablesize:      4096,
		Http2initialwindowsize:    65535,
		Http2maxconcurrentstreams: 100,
		Http2maxframesize:         16384,
		Http2maxheaderlistsize:    24576,
		// Http2minseverconn:         0,
		Incomphdrdelay:   7000,
		Markconnreqinval: "DISABLED",
		Markhttp09inval:  "DISABLED",
		Maxheaderlen:     0,
		Maxreq:           0,
		Maxreusepool:     0,
		Minreusepool:     0,
		Persistentetag:   "DISABLED",
		Reqtimeout:       0,
		Reqtimeoutaction: "RESET",
		Reusepooltimeout: 0,
		Rtsptunnel:       "DISABLED",
		Spdy:             "DISABLED",
		Weblog:           "ENABLED",
		Websocket:        "ENABLED",
	}

	return &resource, func() {
		client.DeleteNshttpprofile(resource.ToKey())
	}
}
