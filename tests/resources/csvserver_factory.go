package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_csvserver(t *testing.T, client *nitro.NitroClient) (*nitro.Csvserver, func()) {
	//resource := nitro.Csvserver {
	//Name: "csvserver",
	//Appflowlog: (ENABLED|DISABLED),
	//Authentication: (ON|OFF),
	//Authenticationhost: string,
	//Authn401: (ON|OFF),
	//Authnprofile: string,
	//Authnvsname: string,
	//Backupvserver: string,
	//Cacheable: (YES|NO),
	//Casesensitive: (ON|OFF),
	//Clttimeout: double,
	//Comment: string,
	//Dbprofilename: dbdbprofile.name,
	//Disableprimaryondown: (ENABLED|DISABLED),
	//Dnsprofilename: dnsprofile.dnsprofilename,
	//Downstateflush: (ENABLED|DISABLED),
	//Httpprofilename: nshttpprofile.name,
	//Icmpvsrresponse: (PASSIVE|ACTIVE),
	//Insertvserveripport: (OFF|VIPADDR|V6TOV4MAPPING),
	//Ipmask: ip_mask,
	//Ippattern: string,
	//Ipv46: ip,
	//L2conn: (ON|OFF),
	//Listenpolicy: string,
	//Listenpriority: double,
	//Mssqlserverversion: (70|2000|2000SP1|2005|2008|2008R2|2012|2014),
	//Mysqlcharacterset: double,
	//Mysqlprotocolversion: double,
	//Mysqlservercapabilities: double,
	//Mysqlserverversion: string,
	//Netprofile: string,
	//Oracleserverversion: (10G|11G),
	//Port: int,
	//Precedence: (RULE|URL),
	//Push: (ENABLED|DISABLED),
	//Pushlabel: string,
	//Pushmulticlients: (YES|NO),
	//Pushvserver: string,
	//Range: double,
	//Redirectportrewrite: (ENABLED|DISABLED),
	//Redirecturl: string,
	//Rhistate: (PASSIVE|ACTIVE),
	//Rtspnat: (ON|OFF),
	//Servicetype: (HTTP|SSL|TCP|FTP|RTSP|SSL_TCP|UDP|DNS|SIP_UDP|SIP_TCP|SIP_SSL|ANY|RADIUS|RDP|MYSQL|MSSQL|DIAMETER|SSL_DIAMETER|DNS_TCP|ORACLE|SMPP),
	//Sobackupaction: (DROP|ACCEPT|REDIRECT),
	//Somethod: (CONNECTION|DYNAMICCONNECTION|BANDWIDTH|HEALTH|NONE),
	//Sopersistence: (ENABLED|DISABLED),
	//Sopersistencetimeout: double,
	//Sothreshold: double,
	//Stateupdate: (ENABLED|DISABLED),
	//Tcpprofilename: nstcpprofile.name,
	//Td: double,
	//Vipheader: string,
	//}

	return nil, func() {
	}
}
