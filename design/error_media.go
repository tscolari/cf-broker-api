package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var ErrorMedia = MediaType("application/cfbroker.error+json", func() {
	Description("Service Broker Catalog")
	Attributes(func() {
		Attribute("description", String, "Error description")

		View("default", func() {
			Attribute("description")
		})
	})
})
