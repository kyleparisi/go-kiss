package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	expectedStr := "Hello world!"
	result := hello()
	if result != expectedStr {
		t.Fatalf("Expected %s, got %s", expectedStr, result)
	}
}

// Reference: https://www.calhoun.io/why-cant-i-pass-this-function-as-an-http-handler/
// Reference: https://blog.questionable.services/article/testing-http-handlers-go/
func TestHandleLogin(t *testing.T)  {
	r, _ := http.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	// func -> HandlerFunc -> Handler
	handler := http.HandlerFunc(handleLogin)
	handler.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		log.Fatalf("Did not receive a 200")
	}
}
