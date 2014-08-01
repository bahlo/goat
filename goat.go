package goat

import "github.com/julienschmidt/httprouter"

func New() *Router {
	r := &Router{}
	r.router = httprouter.New()
	r.router.NotFound = r.notFoundHandler

	return r
}
