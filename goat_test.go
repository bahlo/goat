package goat

import "testing"

func TestNew(t *testing.T) {
	r := New()

	if r == nil {
		t.Errorf("New returned nil, but should return a *Router")
	}
}
