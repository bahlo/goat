package goat

import (
	"net/http/httptest"
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

func TestNotFoundHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := New()

	r.notFoundHandler(w, nil)

	expCode := 404
	if w.Code != expCode {
		t.Errorf("NotFoundHandler should set the status code to %i, but didn't", expCode)
	}

	expBody := `{
  "error": "404 Not Found"
}`
	body := string(w.Body.Bytes())
	if body != expBody {
		t.Errorf("NotFoundHandler set the Body to %s, but should set it to %s", body, expBody)
	}
}
