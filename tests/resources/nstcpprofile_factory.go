package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_nstcpprofile(t *testing.T, client *nitro.NitroClient) (*nitro.Nstcpprofile, func()) {
	//resource := nitro.Nstcpprofile {
	//Name: "nstcpprofile",
	//Ackaggregation: (ENABLED|DISABLED),
	//Ackonpush: (ENABLED|DISABLED),
	//Buffersize: double,
	//Burstratecontrol: (DISABLED|FIXED|DYNAMIC),
	//Delayedack: double,
	//Dropestconnontimeout: (ENABLED|DISABLED),
	//Drophalfclosedconnontimeout: (ENABLED|DISABLED),
	//Dsack: (ENABLED|DISABLED),
	//Dupackthresh: double,
	//Dynamicreceivebuffering: (ENABLED|DISABLED),
	//Ecn: (ENABLED|DISABLED),
	//Establishclientconn: (AUTOMATIC|CONN_ESTABLISHED|ON_FIRST_DATA),
	//Fack: (ENABLED|DISABLED),
	//Flavor: (Default|Westwood|BIC|CUBIC|Nile),
	//Frto: (ENABLED|DISABLED),
	//Hystart: (ENABLED|DISABLED),
	//Initialcwnd: double,
	//Ka: (ENABLED|DISABLED),
	//Kaconnidletime: double,
	//Kamaxprobes: double,
	//Kaprobeinterval: double,
	//Kaprobeupdatelastactivity: (ENABLED|DISABLED),
	//Maxburst: double,
	//Maxcwnd: double,
	//Maxpktpermss: double,
	//Minrto: double,
	//Mptcp: (ENABLED|DISABLED),
	//Mptcpdropdataonpreestsf: (ENABLED|DISABLED),
	//Mptcpfastopen: (ENABLED|DISABLED),
	//Mptcpsessiontimeout: double,
	//Mss: double,
	//Nagle: (ENABLED|DISABLED),
	//Oooqsize: double,
	//Pktperretx: double,
	//Rateqmax: double,
	//Rstmaxack: (ENABLED|DISABLED),
	//Rstwindowattenuate: (ENABLED|DISABLED),
	//Sack: (ENABLED|DISABLED),
	//Sendbuffsize: double,
	//Slowstartincr: double,
	//Spoofsyndrop: (ENABLED|DISABLED),
	//Syncookie: (ENABLED|DISABLED),
	//Tcpfastopen: (ENABLED|DISABLED),
	//Tcpmode: (TRANSPARENT|ENDPOINT),
	//Tcprate: double,
	//Tcpsegoffload: (AUTOMATIC|DISABLED),
	//Timestamp: (ENABLED|DISABLED),
	//Ws: (ENABLED|DISABLED),
	//Wsval: double,
	//}

	return nil, func() {
	}
}
