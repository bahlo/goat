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
