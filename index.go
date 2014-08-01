package goat

import (
	"net/http"
	"sort"
)

// IndexHandler writes the index of all GET methods to the ResponseWriter
func (r *Router) IndexHandler(w http.ResponseWriter, _ *http.Request, _ Params) {
	WriteJSON(w, r.Index())
}

// Index returns a string map with the titles and urls of all GET routes
func (r *Router) Index() map[string]string {
	index := r.index

	// Recursion
	for _, sr := range r.children {
		si := sr.Index()

		for k, v := range si {
			index[k] = v
		}
	}

	// Sort
	sorted := make(map[string]string)
	var keys []string
	for k := range index {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted[k] = index[k]
	}

	return sorted
}
