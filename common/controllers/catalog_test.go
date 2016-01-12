package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/raphael/goa"
	"github.com/tscolari/cf-broker-api/common/controllers"
	"github.com/tscolari/memcached-broker/app"
	"github.com/tscolari/memcached-broker/config"
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Catalog", func() {
	var catalogConfig app.CfbrokerCatalog
	var catalogController *controllers.Catalog

	BeforeEach(func() {
		config, err := config.Load("./assets/valid.config.yml")
		Expect(err).ToNot(HaveOccurred())
		catalogConfig = config.Catalog
		catalogController = controllers.NewCatalog(catalogConfig)
	})

	Describe("#Show", func() {
		var goaContext *goa.Context
		var catalogContext *app.ShowCatalogContext
		var responseWriter *httptest.ResponseRecorder

		BeforeEach(func() {
			gctx := context.Background()
			req := http.Request{}
			responseWriter = httptest.NewRecorder()
			params := url.Values{}
			payload := map[string]string{}

			goaContext = goa.NewContext(gctx, &req, responseWriter, params, payload)

			var err error
			catalogContext, err = app.NewShowCatalogContext(goaContext)
			Expect(err).ToNot(HaveOccurred())
		})

		It("writes the correct catalog information to the context", func() {
			err := catalogController.Show(catalogContext)
			Expect(err).ToNot(HaveOccurred())
			Expect(responseWriter.Body.String()).To(Equal(`{"services":[{"bindable":true,"description":"in memory keystore","id":"my-service-id","metadata":{"what":"about"},"name":"memcached","plan_updatable":false,"plans":[{"description":"shared 100mb memory limit","free":true,"id":"first-plan-id","name":"100mb"},{"description":"shared 1024mb memory limit","free":false,"id":"second-plan-id","name":"1024mb"}]}]}`))
		})
	})
})
