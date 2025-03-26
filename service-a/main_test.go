package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHelloHandler verifies that the helloHandler returns the correct message.
func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	expected := "Hello from Service A from octo!\n"
	if strings.TrimSpace(string(body)) != strings.TrimSpace(expected) {
		t.Errorf("expected %q, got %q", expected, string(body))
	}
}
