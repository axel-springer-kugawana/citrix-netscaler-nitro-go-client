package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-nitro-go-client/tests/resources"
	"testing"
)

func Setup_lbmetrictable_metric_binding(t *testing.T, client *nitro.NitroClient) (*nitro.LbmetrictableMetricBinding, func()) {
	metrictable, metrictableTearDown := resources.Setup_lbmetrictable(t, client)

	client.AddLbmetrictable(*metrictable)

	resource := nitro.LbmetrictableMetricBinding{
		Metrictable: metrictable.Metrictable,
		Metric:      "test",
		Snmpoid:     "test",
	}

	return &resource, func() {
		client.DeleteLbmetrictable(metrictable.ToKey())

		if metrictableTearDown != nil {
			metrictableTearDown()
		}
	}
}
