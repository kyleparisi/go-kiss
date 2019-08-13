package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
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
// Reference: http://blog.ralch.com/tutorial/design-patterns/golang-adapter/
func TestHandleLogin(t *testing.T)  {
	// Create the Request "Object"
	// Object being a pointer to a value in a memory address (*x => &y => object)
	// And values are either base types (string, int, etc), structs with properties
	// that are allocated via make() if of type slices, maps, and channels
	r, _ := http.NewRequest("POST", "/login", nil)
	// Convert the interface "http.ResponseWriter" into "Object"
	// Object being a pointer to a value in a memory address (*x => &y => object)
	w := httptest.NewRecorder()
	handleLogin(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("Not 200 status")
	}
}

// If you wanted to make the objects by hand, this is equivalent to the above
// test.
func TestHandleLoginDetailed(t *testing.T)  {
	u,_ := url.Parse("/login")
	req := &http.Request{
		Method:     "POST",
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       nil,
		Host:       u.Host,
	}
	// ResponseRecorder's body is a memory buffer
	// An actual http.ResponseWriter body is an io.ReadCloser
	res := &httptest.ResponseRecorder{
		Body:      new(bytes.Buffer),
		Code:      200,
	}
	handleLogin(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("Not 200 status")
	}
}
