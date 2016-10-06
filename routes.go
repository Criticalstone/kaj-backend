package main

import (
	"net/http"

	"github.com/criticalstone/kaj-backend/modules/products"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured     bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Index page
	Route{"Index", "GET", "/", false, IndexSite},

	// Products
	Route{"Products", "GET", "/products/", false, products.Index},
	Route{"Products", "GET", "/products/{id:[0-9]+}", false, products.Show},
}
