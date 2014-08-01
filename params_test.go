package goat

import (
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestParamsFromHTTPRouter(t *testing.T) {
	in := httprouter.Params{
		httprouter.Param{"foo", "bar"},
		httprouter.Param{"k", "v"},
	}
	out := paramsFromHTTPRouter(in)
	exp := Params{
		"foo": "bar",
		"k":   "v",
	}

	if !reflect.DeepEqual(out, exp) {
		t.Errorf("paramsFromHTTPRouter returned %v, but should return %v", out, exp)
	}
}
