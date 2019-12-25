package router

import (
	"github.com/gorilla/mux"
)

// router function for all routes in routes.go
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		sub.HandleFunc(route.Pattern, route.HandleFunc).
			Name(route.Name).
			Methods(route.Method)
	}
	return router
}
