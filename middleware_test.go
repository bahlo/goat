package goat

import (
	"net/http"
	"testing"
)

func TestUse(t *testing.T) {
	r := New()
	mw := func(h http.Handler) http.Handler { return nil }
	//exp := []Middleware{mw}

	r.Use(mw)

	// TODO: Test function equality
}
