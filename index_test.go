package goat

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	r := New()

	r.Get("/foo/bar", "foo_bar_url", emptyHandler)
	r.Get("/bar", "bar_url", emptyHandler)
	r.Post("/foo", "foo_url", emptyHandler)

	sr := r.Subrouter("sub")
	sr.Get("/", "sub_url", emptyHandler)

	out := r.Index()
	expected := map[string]string{
		"foo_url":     "/foo",
		"bar_url":     "/bar",
		"foo_bar_url": "/foo/bar",
		"sub_url":     "/sub/",
	}
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Index should return %v, but did return %v", expected, out)
	}
}

func TestIndexHandler(t *testing.T) {
	r := New()
	r.Get("/foo/bar", "foo_bar_url", emptyHandler)
	r.Post("/foo", "foo_url", emptyHandler)
	r.Get("/bar", "bar_url", emptyHandler)

	w := httptest.NewRecorder()
	p := Params{}
	r.IndexHandler(w, nil, p)

	expected := `{
  "bar_url": "/bar",
  "foo_bar_url": "/foo/bar",
  "foo_url": "/foo"
}`
	body := string(w.Body.Bytes())

	if body != expected {
		t.Errorf("indexHandler should return %s, but did return %s", expected, body)
	}
}
