package goat

import (
	"net/http"
	"reflect"
	"testing"
)

var emptyHandler Handle

func TestSubPath(t *testing.T) {
	r := New()

	out := r.subPath("/foo")
	exp := "/foo"

	if out != exp {
		t.Errorf("subPath should return %s, but did return %s", exp, out)
	}
}

func TestAddRoute(t *testing.T) {
	r := New()

	r.addRoute("GET", "/test", "", emptyHandler)
	expected := make(map[string]string)
	if !reflect.DeepEqual(r.index, expected) {
		t.Errorf("addRoute with empty title should modify Router.index to %v, but it's %v", expected, r.index)
	}

	r.addRoute("GET", "/foo/bar", "foo_bar_url", emptyHandler)
	expected = map[string]string{
		"foo_bar_url": "/foo/bar",
	}
	if !reflect.DeepEqual(r.index, expected) {
		t.Errorf("addRoute should modify Router.index to %v, but did modify it to %v", expected, r.index)
	}
}

func TestIndex(t *testing.T) {
	r := New()

	r.Get("/foo/bar", "foo_bar_url", emptyHandler)
	r.Get("/bar", "bar_url", emptyHandler)
	r.Post("/foo", "foo_url", emptyHandler)

	out := r.Index()
	expected := map[string]string{
		"bar_url":     "/bar",
		"foo_bar_url": "/foo/bar",
	}
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Index should regurn %v, but did return %v", expected, out)
	}
}

func TestSubrouter(t *testing.T) {
	pre := "/user"
	r := New()
	sr := r.Subrouter(pre)

	if sr.prefix != pre {
		t.Errorf("Subrouter should set the prefix %s, but did set %s", pre, sr.prefix)
	}

	if sr.parent != r {
		t.Errorf("Subrouter should set %v as parent router, but did set %v", r, sr.parent)
	}

	if r.children[len(r.children)-1] != sr {
		t.Errorf("Subrouter should add %v to children of %v, but didn't", sr, r)
	}
}

func TestUse(t *testing.T) {
	r := New()
	mw := func(h http.Handler) http.Handler { return nil }
	//exp := []Middleware{mw}

	r.Use(mw)

	// TODO: Test function equality
}
