package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var DashboardMedia = MediaType("application/cfbroker.dashboard+json", func() {
	Description("Service Dashboard")
	Attributes(func() {
		Attribute("dashboard_url", String, "The URL of a web-based management user interface for the service instance; we refer to this as a service dashboard")

		View("default", func() {
			Attribute("dashboard_url")
		})
	})
})
