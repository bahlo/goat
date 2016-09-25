package goat

import (
	"net/http"
	"reflect"
)

// IndexHandler writes the index of all GET methods to the ResponseWriter
func (r *Router) IndexHandler(w http.ResponseWriter, _ *http.Request, _ Params) {
	WriteJSON(w, r.Index())
}

// Index returns a string map with the titles and urls of all routes, grouped by method
func (r *Router) Index() map[string]map[string]string {
	index := r.methodIndex

	// Recursion
	for _, sr := range r.children {
		si := sr.Index()

		for method, routes := range si {
			for title, url := range routes {
				if reflect.DeepEqual(index, reflect.Zero(reflect.TypeOf(index)).Interface()) {
					index = make(map[string]map[string]string)
				}
				if reflect.DeepEqual(index[method], reflect.Zero(reflect.TypeOf(index[method])).Interface()) {
					index[method] = make(map[string]string)
				}
				index[method][title] = url
			}
		}
	}

	return index
}
