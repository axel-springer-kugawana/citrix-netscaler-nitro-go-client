package main

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"log"
)

func main() {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	test_Server(client)
	// test_Appflowcollector(client)
}

func test_Server(client *nitro.NitroClient) {
	server := nitro.Server{
		Name:      "blibli",
		Ipaddress: "1.3.5.7",
		// Domain: "abc.com",
		Comment: "comment",
		State:   "ENABLED",
		// Ipv6address: "NO",
		// Td: 0,
		// Domainresolveretry: 5,
		// Translationip: "1.2.3.4",
		// Translationmask: "5.6.7.8",
		// Graceful: "YES",
		// Delay: 123,
		// Domainresolvenow: true,
	}

	err := client.AddServer(server)

	log.Print(err)

	res, err := client.GetServer(server.Name)

	log.Print(err)
	log.Print(res)

	list, err := client.ListServer()

	log.Print(err)
	log.Print(list)

	server.Ipaddress = "2.4.6.8"

	err = client.UpdateServer(server)

	log.Print(err)

	err = client.EnableServer(server.Name)

	log.Print(err)

	err = client.DisableServer(server.Name)

	log.Print(err)

	err = client.RenameServer(server.Name, server.Name+"-rename")

	log.Print(err)

	err = client.DeleteServer(server.Name + "-rename")

	log.Print(err)
}
