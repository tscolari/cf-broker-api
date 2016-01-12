package controllers

import (
	"github.com/raphael/goa"
	"github.com/tscolari/memcached-broker/app"
)

type Catalog struct {
	goa.Controller
	catalog app.CfbrokerCatalog
}

type ShowCatalogContext struct {
	*goa.Context
}

func NewCatalog(config app.CfbrokerCatalog) *Catalog {
	return &Catalog{
		catalog: config,
	}
}

func (c *Catalog) Show(ctx *app.ShowCatalogContext) error {
	return ctx.OK(&c.catalog)
}
