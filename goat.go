package goat

import "github.com/julienschmidt/httprouter"

// New creates a new Router and returns it
func New() *Router {
	r := &Router{}
	r.prefix = "/"
	r.router = httprouter.New()
	r.router.NotFound = r.notFoundHandler

	return r
}
