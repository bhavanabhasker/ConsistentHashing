package main

import (
	"assign4/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range server.Routess {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
