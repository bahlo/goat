package goat

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	// In
	code := 500
	err := "foo"

	// Expected
	json := `{
  "error": "` + err + `"
}
`
	buf := bytes.NewBufferString(json)

	w := httptest.NewRecorder()
	WriteError(w, code, err)

	// Test Code
	if w.Code != code {
		t.Errorf("WriteError should set Code to %i, but did set it to %i", code, w.Code)
	}

	// Test Header
	expectedContentTypeHeader := "application/json; charset=utf-8"
	contentTypeHeader := w.Header().Get("Content-Type")
	if contentTypeHeader != expectedContentTypeHeader {
		t.Errorf("WriteError should set Content-Type header to %s, but did set it to %s", expectedContentTypeHeader, contentTypeHeader)
	}

	// Test Body
	if w.Body == nil {
		t.Errorf("WriteError should set Body to %s, but didn't", json)
	} else if string(w.Body.Bytes()) == string(buf.Bytes()) {
		t.Errorf("WriteError should set Body to %v, but did set it to %v", buf, w.Body)
	}
}

func TestWriteJSON(t *testing.T) {
	in := map[string]string{
		"foo":   "bar",
		"knock": "knock",
	}
	json := `{
  "foo": "bar",
  "knock": "knock"
}
`
	buf := bytes.NewBufferString(json)

	w := httptest.NewRecorder()
	WriteJSON(w, in)

	// Test Body
	if w.Body == nil {
		t.Errorf("WriteJSON should set the Body to %s, but didn't", json)
	} else if string(w.Body.Bytes()) == string(buf.Bytes()) {
		t.Errorf("WriteJSON set the Body to %v, but should set it to %v", buf, w.Body)
	}

	// Test Header
	expectedContentTypeHeader := "application/json; charset=utf-8"
	contentTypeHeader := w.Header().Get("Content-Type")
	if contentTypeHeader != expectedContentTypeHeader {
		t.Errorf("WriteJSON should set Content-Type header to %s, but did set it to %s", expectedContentTypeHeader, contentTypeHeader)
	}

	// Test Error
	w = httptest.NewRecorder()
	if err := WriteJSON(w, WriteJSON); err == nil {
		t.Errorf("WriteJSON should return an error, but didn't")
	}
}

func TestWriteJSONWithStatus(t *testing.T) {
	// in
	code := 201
	in := map[string]interface{}{
		"foo": "bar",
		"bar": "foo",
	}
	json := `{
  "foo": "bar",
  "bar": "foo"
}
`
	buf := bytes.NewBufferString(json)

	w := httptest.NewRecorder()
	WriteJSONWithStatus(w, code, in)

	// Test Code
	if w.Code != code {
		t.Errorf("WriteJSONWithStatus should set Code to %i, but did set it to %i", code, w.Code)
	}

	// Test Header
	expectedContentTypeHeader := "application/json; charset=utf-8"
	contentTypeHeader := w.Header().Get("Content-Type")
	if contentTypeHeader != expectedContentTypeHeader {
		t.Errorf("WriteJSONWithStatus should set Content-Type header to %s, but did set it to %s", expectedContentTypeHeader, contentTypeHeader)
	}

	// Test Body
	if w.Body == nil {
		t.Errorf("WriteJSONWithStatus should set the Body to %s, but didn't", json)
	} else if string(w.Body.Bytes()) == string(buf.Bytes()) {
		t.Errorf("WriteJSONWithStatus set the Body to %v, but should set it to %v", buf, w.Body)
	}
}
