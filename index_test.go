package goat

import (
	"reflect"
	"testing"
)

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
