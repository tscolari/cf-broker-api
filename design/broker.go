package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = API("cfbroker", func() {
	Description("CloudFoundry Service Broker API")
	BasePath("/v2")

	Host("127.0.0.1")
	Scheme("http")
	Scheme("https")
})

var _ = Resource("Catalog", func() {
	DefaultMedia(CatalogMedia)
	BasePath("/catalog")

	Action("show", func() {
		Routing(GET("/"))
		Description("Service broker catalog")
		Response(OK)
	})
})

var _ = Resource("Provisioning", func() {
	BasePath("/service_instances")

	Action("create", func() {
		Routing(GET("/:instance_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "id of the service being provisioned")
			Param("organization_guid", String, "The Cloud Controller GUID of the organization under which the service is to be provisioned")
			Param("plan_id", String, "The ID of the plan that the user would like provisioned")
			Param("service_id", String, "The ID of the service within the catalog")
			Param("space_guid", String, "Similar to organization_guid, but for the space")
			Param("accepts_incomplete", Boolean, "A value of true indicates that both the Cloud Controller and the requesting client support asynchronous provisioning")
		})
		Response(ServiceUnavailable)
		Response(Created)
		Response(Conflict)
	})
})
