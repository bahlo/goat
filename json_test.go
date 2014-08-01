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

	// Test code
	if w.Code != code {
		t.Errorf("WriteError should set Code to %i, but did set it to %i", code, w.Code)
	}

	// Test body
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

	if w.Body == nil {
		t.Errorf("WriteJSON should set the Body to %s, but didn't", json)
	} else if string(w.Body.Bytes()) == string(buf.Bytes()) {
		t.Errorf("WriteJSON set the Body to %v, but should set it to %v", buf, w.Body)
	}

	// Test error
	w = httptest.NewRecorder()
	if err := WriteJSON(w, nil); err != nil {
		t.Errorf("WriteJSON should return an error, but didn't")
	}
}
