package goat

import (
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

var emptyHandler httprouter.Handle

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

	out := r.IndexGet()
	expected := map[string]string{
		"bar_url":     "/bar",
		"foo_bar_url": "/foo/bar",
	}
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("IndexGet should regurn %v, but did return %v", expected, out)
	}
}
