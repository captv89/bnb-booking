package test

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		t.Log("Handler is http.Handler")
	default:
		t.Error("Handler is not http.Handler", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		t.Log("Handler is http.Handler")
	default:
		t.Error("Handler is not http.Handler", v)
	}
}

