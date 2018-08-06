package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func create_appflowaction(t *testing.T) *nitro.Appflowaction {
	collector := create_appflowcollector()

	resource := nitro.Appflowaction {
		Name: "appflowaction",
		Clientsidemeasurements: "ENABLED",
		Comment: "comment",
		Pagetracking: "ENABLED",
		Securityinsight: "ENABLED",
		Webinsight: "ENABLED",
	}

	resource.Collectors = append(resource.Collectors, collector.Name)

	return &resource
}
