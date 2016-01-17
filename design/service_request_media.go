package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var ServiceRequestMedia = MediaType("application/cfbroker.service-request+json", func() {
	Description("Service Request")
	Attributes(func() {
		Attribute("instance_id", String, "id of the service being provisioned")
		Attribute("organization_id", String, "The Cloud Controller GUID of the organization under which the service is to be provisioned")
		Attribute("plan_id", String, "The ID of the plan that the user would like provisioned")
		Attribute("service_id", String, "The ID of the service within the catalog")
		Attribute("space_id", String, "Similar to organization_guid, but for the space")

		View("default", func() {
			Attribute("instance_id")
			Attribute("organization_id")
			Attribute("plan_id")
			Attribute("space_id")
			Attribute("service_id")
		})
	})
})
