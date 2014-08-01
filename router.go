package goat

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router represents a router
type Router struct {
	prefix string
	router *httprouter.Router
}

// notFoundHandler handles (as you already know) the 404 error
func (r *Router) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	WriteError(w, "404 Not Found")
}

// ServeHTTP calls the same method on the router
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
