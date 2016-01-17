package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var DashboardClientMedia = MediaType("application/cfbroker.dashboard-client+json", func() {
	Description("Service Plan")
	Attributes(func() {
		Attribute("id", String, "The id of the Oauth2 client that the service intends to use.")
		Attribute("secret", String, "A secret for the dashboard client")
		Attribute("redirect_url", String, "A domain for the service dashboard that will be whitelisted by the UAA to enable SSO")

		View("default", func() {
			Attribute("id")
			Attribute("secret")
			Attribute("redirect_url")
		})
	})
})
