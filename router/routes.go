package router

import (
	"github.com/kumaranshu72/go-rest-api-boilerplate/controllers"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Health Check",
		"GET",
		"/health",
		controllers.Health,
	},
}
