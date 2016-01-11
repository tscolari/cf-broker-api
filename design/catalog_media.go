package design

import (
	. "github.com/raphael/goa/design/dsl"
)

var CatalogMedia = MediaType("application/cfbroker.catalog+json", func() {
	Description("Service Broker Catalog")
	Attributes(func() {
		Attribute("services", ArrayOf(ServiceMedia), "Offered Services")

		View("default", func() {
			Attribute("services")
		})
	})
})
