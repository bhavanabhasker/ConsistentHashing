package main

import (
	"assign4/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := NewRouter1()

	log.Fatal(http.ListenAndServe(":3001", router))
}
func NewRouter1() *mux.Router {
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
