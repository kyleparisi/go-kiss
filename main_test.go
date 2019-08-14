package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestHandleLogin(t *testing.T)  {
	req, _ := http.NewRequest("POST", "/login", nil)
	res := httptest.NewRecorder()
	handleLogin(res, req, httprouter.Params{})
	assert.Assert(t, strings.Contains(res.Body.String(), "Please provide an email"))
	assert.Assert(t, strings.Contains(res.Body.String(), "Please provide a password"))
}
