package goat

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router represents a router
type Router struct {
	// Tree structure
	parent   *Router
	children []*Router

	// The prefix, default: /
	prefix string

	// The index, maps titles to urls
	index map[string]string

	// The router
	router *httprouter.Router

	// Middleware
	middleware []Middleware
}

// Handle describes the function that should be used by handlers
type Handle func(http.ResponseWriter, *http.Request, Params)

// Subrouter creates and returns a subrouter
func (r *Router) Subrouter(path string) *Router {
	sr := &Router{
		index:  make(map[string]string),
		prefix: r.subPath(path),
		router: r.router,
	}

	// Init relationships
	r.children = append(r.children, sr)
	sr.parent = r

	return sr
}

// addRoute adds a route to the index and passes it over to the httprouter
func (r *Router) addRoute(m, p, t string, fn Handle) {

	path := r.subPath(p)

	// Add to index
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
func (r *Router) Get(path, title string, fn Handle) {
	r.addRoute("GET", path, title, fn)
}

// Post adds a POST route
func (r *Router) Post(path, title string, fn Handle) {
	r.addRoute("POST", path, title, fn)
}

// Delete adds a DELETE route
func (r *Router) Delete(path, title string, fn Handle) {
	r.addRoute("DELETE", path, title, fn)
}

// Put adds a PUT route
func (r *Router) Put(path, title string, fn Handle) {
	r.addRoute("PUT", path, title, fn)
}

// Options adds a OPTIONS route
func (r *Router) Options(path, title string, fn Handle) {
	r.addRoute("OPTIONS", path, title, fn)
}

// notFoundHandler handles (as you already know) the 404 error
func (r *Router) notFoundHandler(w http.ResponseWriter, req *http.Request) {
	WriteError(w, 404, "404 Not Found")
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
