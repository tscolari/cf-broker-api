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
