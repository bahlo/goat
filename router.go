package goat

import "github.com/julienschmidt/httprouter"

// Router represents a router
type Router struct {
	middleware []Middleware
	router     *httprouter.Router
}
