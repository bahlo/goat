package goat

import (
	"encoding/json"
	"net/http"
)

// WriteError writes a string as JSON encoded error
func WriteError(w http.ResponseWriter, code int, err string) {
	w.WriteHeader(code)

	WriteJSON(w, map[string]string{
		"error": err,
	})
}

// WriteJSON writes the given interface as JSON or returns an error
func WriteJSON(w http.ResponseWriter, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
	return nil
}

// WriteJSONWithStatus writes the given statuscode into the header and the
// given interface as JSON  or returns an error
func WriteJSONWithStatus(w http.ResponseWriter, code int, v interface{}) error {
	w.WriteHeader(code)

	return WriteJSON(w, v)
}
