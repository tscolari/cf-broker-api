package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var ServiceMedia = MediaType("application/cfbroker.service+json", func() {
	Description("Service")
	Attributes(func() {
		Attribute("id", Integer, "An identifier used to correlate this service in future requests to the catalog")
		Attribute("name", String, "CLI-friendly name of the service that will appear in the catalog")
		Attribute("description", String, "Short description of the service that will appear in the catalog")
		Attribute("bindable", Boolean, "Whether the service can be bound to applications")
		Attribute("tags", ArrayOf(String), "Service tags")
		Attribute("metadata", HashOf(String, String), "A list of metadata for a service offering")
		Attribute("requires", ArrayOf(String), "A list of permissions that the user would have to give the service, if they provision it")
		Attribute("plan_updatable", Boolean, "Whether the service supports upgrade/downgrade for some plans")
		Attribute("plans", ArrayOf(PlanMedia), "A list of plans for this service")

		View("default", func() {
			Attribute("id")
			Attribute("name")
			Attribute("description")
			Attribute("bindable")
			Attribute("tags")
			Attribute("metadata")
			Attribute("requires")
			Attribute("plan_updatable")
			Attribute("plans")
		})
	})
})
