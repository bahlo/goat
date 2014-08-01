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
	} else if bytes.Equal(w.Body.Bytes(), buf.Bytes()) {
		t.Errorf("WriteError should set Body to %v, but did set it to %v", buf, w.Body)
	}
}
