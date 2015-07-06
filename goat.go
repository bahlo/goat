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
	r.router.NotFound = http.HandlerFunc(r.notFoundHandler)

	return r
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h := r.chain()
	h.ServeHTTP(w, req)
}

// Run starts the server
func (r *Router) Run(address string) error {
	return http.ListenAndServe(address, r.chain())
}

// RunTLS starts the server, but expects HTTPS connections
func (r *Router) RunTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, r.chain())
}
