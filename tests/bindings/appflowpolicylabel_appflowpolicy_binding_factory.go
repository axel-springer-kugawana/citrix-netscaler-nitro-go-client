package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-nitro-go-client/tests/resources"
	"testing"
)

func Setup_appflowpolicylabel_appflowpolicy_binding(t *testing.T, client *nitro.NitroClient) (*nitro.AppflowpolicylabelAppflowpolicyBinding, func()) {
	appflowpolicylabel, appflowpolicylabelTearDown := resources.Setup_appflowpolicylabel(t, client)
	appflowpolicy, appflowpolicyTearDown := resources.Setup_appflowpolicy(t, client)

	client.AddAppflowpolicylabel(*appflowpolicylabel)
	client.AddAppflowpolicy(*appflowpolicy)

	resource := nitro.AppflowpolicylabelAppflowpolicyBinding{
		Labelname:  appflowpolicylabel.Labelname,
		Policyname: appflowpolicy.Name,
		// Gotopriorityexpression: string,
		Invoke:           true,
		Invoke_labelname: appflowpolicylabel.Labelname,
		Labeltype:        "policylabel",
		Priority:         1,
	}

	return &resource, func() {
		client.DeleteAppflowpolicy(appflowpolicy.ToKey())
		client.DeleteAppflowpolicylabel(appflowpolicylabel.ToKey())

		if appflowpolicyTearDown != nil {
			appflowpolicyTearDown()
		}

		if appflowpolicylabelTearDown != nil {
			appflowpolicylabelTearDown()
		}
	}
}
