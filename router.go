package goat

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router represents a router
type Router struct {
	middleware []Middleware
	prefix     string
	router     *httprouter.Router
}

// notFoundHandler handles (as you already know) the 404 error
func (r *Router) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement notFoundHandler
}

// Use adds middleware(s) to the router
func (r *Router) Use(m ...Middleware) {
	r.middleware = append(r.middleware, m...)
}
