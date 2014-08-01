package goat

import (
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
)

// Router represents a router
type Router struct {
	prefix string
	router *httprouter.Router
	index  map[string]string
}

// Handle describes the function that should be used by handlers
type Handle func(http.ResponseWriter, *http.Request, Params)

// notFoundHandler handles (as you already know) the 404 error
func (r *Router) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	WriteError(w, "404 Not Found")
}

// ServeHTTP calls the same method on the router
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

// subPath returns the prefix of the router + the given path and eliminates
// duplicate slashes
func (r *Router) subPath(p string) string {
	pre := r.prefix

	if (pre == "/" || pre[:len(pre)-1] == "/") && p[:1] == "/" {
		pre = pre[:len(pre)-1]
	}

	return pre + p
}

// addRoute adds a route to the index and passes it over to the httprouter
func (r *Router) addRoute(m, p, t string, fn Handle) {
	path := r.subPath(p)

	// Add to index
	// TODO: Only GET?
	if len(t) > 0 && m == "GET" {
		// TODO: Display total path including host
		r.index[t] = path
	}

	// Wrapper function to bypass the parameter problem
	wf := func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		fn(w, req, paramsFromHTTPRouter(p))
	}

	r.router.Handle(m, path, wf)
}

// Get adds a GET route
func (r *Router) Get(p, t string, fn Handle) {
	r.addRoute("GET", p, t, fn)
}

// Get adds a POST route
func (r *Router) Post(p, t string, fn Handle) {
	r.addRoute("POST", p, t, fn)
}

// Get adds a DELETE route
func (r *Router) Delete(p, t string, fn Handle) {
	r.addRoute("DELETE", p, t, fn)
}

// Get adds a PUT route
func (r *Router) Put(p, t string, fn Handle) {
	r.addRoute("PUT", p, t, fn)
}

// TODO: Add PATCH, OPTIONS, HEAD?

// Index returns a string map with the titles and urls of all GET routes
func (r *Router) IndexGet() map[string]string {
	// Sort
	sorted := make(map[string]string)
	var keys []string
	for k := range r.index {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted[k] = r.index[k]
	}

	return sorted
}
