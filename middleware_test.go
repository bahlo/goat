package goat

import (
	"net/http"
	"testing"
)

func TestUse(t *testing.T) {
	r := New()
	mw := func(h http.Handler) http.Handler { return nil }
	r.Use(mw)

	if len(r.middleware) == 0 {
		t.Errorf("Use should add one item to middleware, but didn't")
	}

	// TODO: Check function equality
}

func TestUseSubrouter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when using middleware on subrouter")
		}
	}()

	mw := func(h http.Handler) http.Handler { return nil }

	r := New()
	sr := r.Subrouter("test")
	sr.Use(mw)
}

func TestAllMiddleware(t *testing.T) {
	mw := func(h http.Handler) http.Handler { return nil }

	r := New()
	r.Use(mw)

	sr := r.Subrouter("/test")
	ssr := sr.Subrouter("/test")

	if len(ssr.allMiddleware()) != 1 {
		t.Errorf("Expected one middleware from parents in ssr")
	}
}

func TestChain(t *testing.T) {
	c := 0
	mw := func(h http.Handler) http.Handler {
		c++
		return h
	}

	r := New()
	r.Use(mw, mw, mw)
	r.chain()

	if c != 3 {
		t.Errorf("Middleware not correctly chained")
	}
}
