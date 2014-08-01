package goat

import "net/http"

// Middleware reprents a default middleware function
type Middleware func(http.Handler) http.Handler

// chain calls all middlewares and returns the final handler
func (r *Router) chain() http.Handler {
	var final http.Handler
	final = r.router

	for i := len(r.middleware) - 1; i >= 0; i-- {
		final = r.middleware[i](final)
	}

	return final
}

// Use adds middleware to the router
func (r *Router) Use(middleware ...Middleware) {
	if r.parent != nil {
		panic("subrouters can't use middleware!")
	}
	r.middleware = append(r.middleware, middleware...)
}
