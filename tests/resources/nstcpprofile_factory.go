package resources

import (
	"testing"

	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
)

func Setup_nstcpprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Nstcpprofile, func()) {
	resource := nitro.Nstcpprofile{
		Name:             "nstcpprofile",
		Ackaggregation:   "DISABLED",
		Ackonpush:        "ENABLED",
		Buffersize:       8190,
		Burstratecontrol: "DISABLED",
		Delayedack:       100,
		// Dropestconnontimeout:        "",
		// Drophalfclosedconnontimeout: "",
		Dsack:                   "ENABLED",
		Dupackthresh:            3,
		Dynamicreceivebuffering: "DISABLED",
		Ecn:                       "DISABLED",
		Establishclientconn:       "AUTOMATIC",
		Fack:                      "DISABLED",
		Flavor:                    "Default",
		Frto:                      "DISABLED",
		Hystart:                   "DISABLED",
		Initialcwnd:               4,
		Ka:                        "DISABLED",
		Kaconnidletime:            900,
		Kamaxprobes:               3,
		Kaprobeinterval:           75,
		Kaprobeupdatelastactivity: "DISABLED",
		Maxburst:                  6,
		Maxcwnd:                   8190,
		Maxpktpermss:              2,
		Minrto:                    1000,
		Mptcp:                     "DISABLED",
		Mptcpdropdataonpreestsf: "DISABLED",
		Mptcpfastopen:           "DISABLED",
		Mptcpsessiontimeout:     120,
		Mss:                     0,
		Nagle:                   "DISABLED",
		Oooqsize:                64,
		Pktperretx:              1,
		Rateqmax:                0,
		Rstmaxack:               "DISABLED",
		Rstwindowattenuate:      "DISABLED",
		Sack:                    "DISABLED",
		Sendbuffsize:            20971520,
		Slowstartincr:           1,
		Spoofsyndrop:            "DISABLED",
		Syncookie:               "DISABLED",
		Tcpfastopen:             "DISABLED",
		Tcpmode:                 "TRANSPARENT",
		Tcprate:                 1,
		Tcpsegoffload:           "AUTOMATIC",
		Timestamp:               "DISABLED",
		Ws:                      "DISABLED",
		Wsval:                   2,
	}

	return &resource, func() {
		client.DeleteNstcpprofile(resource.ToKey())
	}
}
