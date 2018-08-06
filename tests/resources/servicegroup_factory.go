package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_servicegroup(t *testing.T, client *nitro.NitroClient) (*nitro.Servicegroup, func()) {
	//resource := nitro.Servicegroup {
	//Servicegroupname: "servicegroup",
	//Appflowlog: (ENABLED|DISABLED),
	//Autoscale: (DISABLED|DNS|POLICY),
	//Cacheable: (YES|NO),
	//Cachetype: (TRANSPARENT|REVERSE|FORWARD),
	//Cip: (ENABLED|DISABLED),
	//Cipheader: string,
	//Cka: (YES|NO),
	//Clttimeout: int,
	//Cmp: (YES|NO),
	//Comment: string,
	//Downstateflush: (ENABLED|DISABLED),
	//Healthmonitor: (YES|NO),
	//Httpprofilename: nshttpprofile.name,
	//Maxbandwidth: double,
	//Maxclient: double,
	//Maxreq: double,
	//Memberport: int,
	//Monconnectionclose: (RESET|FIN),
	//Monthreshold: double,
	//Netprofile: netprofile.name,
	//Pathmonitor: (YES|NO),
	//Pathmonitorindv: (YES|NO),
	//Rtspsessionidremap: (ON|OFF),
	//Sc: (ON|OFF),
	//Servicetype: (HTTP|FTP|TCP|UDP|SSL|SSL_BRIDGE|SSL_TCP|DTLS|NNTP|RPCSVR|DNS|ADNS|SNMP|RTSP|DHCPRA|ANY|SIP_UDP|SIP_TCP|SIP_SSL|DNS_TCP|ADNS_TCP|MYSQL|MSSQL|ORACLE|RADIUS|RADIUSListener|RDP|DIAMETER|SSL_DIAMETER|TFTP|SMPP|PPTP|GRE|SYSLOGTCP|SYSLOGUDP|FIX|SSL_FIX|USER_TCP|USER_SSL_TCP),
	//Sp: (ON|OFF),
	//State: (ENABLED|DISABLED),
	//Svrtimeout: int,
	//Tcpb: (YES|NO),
	//Tcpprofilename: nstcpprofile.name,
	//Td: double,
	//Useproxyport: (YES|NO),
	//Usip: (YES|NO),
	//}

	return nil, func() {
	}
}
