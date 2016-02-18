package goat

import "net/http"

// Middleware reprents a default middleware function
type Middleware func(http.Handler) http.Handler

// chain calls all middlewares and returns the final handler
func (r *Router) chain() http.Handler {
	var final http.Handler

	final = r.router
	mw := r.allMiddleware()
	for i := len(mw) - 1; i >= 0; i-- {
		final = mw[i](final)
	}

	return final
}

// allMiddleware returns the middleware from this router and all parents
func (r *Router) allMiddleware() []Middleware {
	mw := r.middleware

	if r.parent != nil {
		mw = append(mw, r.parent.allMiddleware()...)
	}

	return mw
}

// Use adds middleware to the router
func (r *Router) Use(middleware ...Middleware) {
	if r.parent != nil {
		panic("subrouters can't use middleware!")
	}
	r.middleware = append(r.middleware, middleware...)
}
