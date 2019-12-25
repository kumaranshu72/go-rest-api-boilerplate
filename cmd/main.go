package main

import (
	"fmt"
	"github.com/kumaranshu72/go-rest-api-boilerplate/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home o!")
}

// CORS MIDDLEWARE
func setUpGlobalMiddleWare(handle http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handle)
}

// Main server function
func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", setUpGlobalMiddleWare(router)))
}
