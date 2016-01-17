package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var BindingMedia = MediaType("application/cfbroker.service-binding+json", func() {
	Description("Service Binding")
	Attributes(func() {
		Attribute("credentials", String, "A free-form hash of credentials that the bound application can use to access the service")
		Attribute("syslog_drain_url", String, "A URL to which Cloud Foundry should drain logs for the bound application")
		Attribute("route_service_url", String, "A URL to which Cloud Foundry should proxy requests for the bound route")

		View("default", func() {
			Attribute("credentials")
			Attribute("syslog_drain_url")
			Attribute("route_service_url")
		})
	})
})
