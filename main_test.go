package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestHandleLogin(t *testing.T)  {
	login := Login{}
	body, _ := json.Marshal(login)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handleLogin(res, req, httprouter.Params{})
	assert.Assert(t, strings.Contains(res.Body.String(), "Please provide an email"))
	assert.Assert(t, strings.Contains(res.Body.String(), "Please provide a password"))

	login.Email = "blah"
	login.Password = "blah"
	body, _ = json.Marshal(login)
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(body))
	res = httptest.NewRecorder()
	handleLogin(res, req, httprouter.Params{})
	assert.Assert(t, strings.Contains(res.Body.String(), "Not a valid email address"))

	login.Email = "kyle@example.com"
	login.Password = "blah"
	body, _ = json.Marshal(login)
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(body))
	res = httptest.NewRecorder()
	handleLogin(res, req, httprouter.Params{})
	fmt.Println(res.Body.String())
}
