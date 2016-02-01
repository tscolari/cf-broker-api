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
	DefaultMedia(DashboardMedia)
	BasePath("/service_instances")

	Action("create", func() {
		Routing(GET("/:instance_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "id of the service being provisioned")
			Param("organization_id", String, "The Cloud Controller GUID of the organization under which the service is to be provisioned")
			Param("plan_id", String, "The ID of the plan that the user would like provisioned")
			Param("service_id", String, "The ID of the service within the catalog")
			Param("space_id", String, "Similar to organization_guid, but for the space")
			Param("accepts_incomplete", Boolean, "A value of true indicates that both the Cloud Controller and the requesting client support asynchronous provisioning")
		})
		Response(ServiceUnavailable)
		Response(Created)
		Response(Conflict)
	})

	Action("update", func() {
		Routing(PUT("/:instance_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "id of the service being provisioned")
			Param("plan_id", String, "The ID of the plan that the user would like provisioned")
			Param("service_id", String, "The ID of the service within the catalog")
			Param("previous_values", ServiceRequestMedia, "Information about the instance prior to the update")
			Param("accepts_incomplete", Boolean, "A value of true indicates that both the Cloud Controller and the requesting client support asynchronous provisioning")
		})
		Response(OK)
		Response(Accepted)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(DELETE("/:instance_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "ID of the service being deprovisioned")
			Param("service_id", String, "ID of the service from the catalog. While not strictly necessary, some brokers might make use of this ID")
			Param("plan_id", String, "ID of the plan from the catalog. While not strictly necessary, some brokers might make use of this ID")
			Param("accepts_incomplete", Boolean, "A value of true indicates that both the Cloud Controller and the requesting client support asynchronous provisioning")
		})
		Response(OK)
		Response(Accepted)
		Response(Gone)
	})
})

var _ = Resource("Binding", func() {
	DefaultMedia(DashboardMedia)
	BasePath("/service_instances/:instance_id/service_bindings")

	Action("update", func() {
		Routing(GET("/:binding_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "id of the service being provisioned")
			Param("binding_id", String, "id of the binding being provisioned")

			Param("app_guid", String, "GUID of the application that you want to bind your service to")
			Param("service_id", String, "The ID of the service within the catalog")
			Param("plan_id", String, "The ID of the plan that the user would like provisioned")
		})
		Response(Created)
		Response(OK)
		Response(NotFound)
		Response(Conflict)
		Response(InternalServerError)
	})

	Action("delete", func() {
		Routing(DELETE("/:binding_id"))
		Description("Provision a new service instance")
		Params(func() {
			Param("instance_id", String, "id of the service being provisioned")
			Param("binding_id", String, "id of the binding being provisioned")

			Param("service_id", String, "ID of the service from the catalog. While not strictly necessary, some brokers might make use of this ID")
			Param("plan_id", String, "ID of the plan from the catalog. While not strictly necessary, some brokers might make use of this ID")
		})
		Response(OK)
		Response(Gone)
		Response(InternalServerError)
	})
})
