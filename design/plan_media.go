package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var PlanMedia = MediaType("application/cfbroker.plan+json", func() {
	Description("Service Plan")
	Attributes(func() {
		Attribute("id", String, "An identifier used to correlate this plan in future requests to the catalog.")
		Attribute("name", String, "CLI-friendly name of the plan that will appear in the catalog")
		Attribute("description", String, "Short description of the plan that will appear in the catalog")
		Attribute("metadata", HashOf(String, String), "A list of metadata for a service offering")
		Attribute("free", Boolean, "This field allows the plan to be limited by the non_basic_services_allowed field in a Cloud Foundry Quota")
		Attribute("dashboard_client", DashboardClientMedia, "This field allows the plan to be limited by the non_basic_services_allowed field in a Cloud Foundry Quota")

		View("default", func() {
			Attribute("id")
			Attribute("name")
			Attribute("description")
			Attribute("metadata")
			Attribute("free")
			Attribute("dashboard_client")
		})
	})
})
