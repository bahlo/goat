package goat

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New creates a new Router and returns it
func New() *Router {
	r := &Router{}
	r.index = make(map[string]string)
	r.prefix = "/"
	r.router = httprouter.New()
	r.router.NotFound = r.notFoundHandler

	return r
}

// Run starts the server
func (r *Router) Run(address string) error {
	return http.ListenAndServe(address, r.chain())
}
