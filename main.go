package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/julienschmidt/httprouter"
)

type Login struct {
	Email string
	Password string
	Remember bool
}

type LoginError struct {
	Errors struct{
		Email string `json:"email"`
		Password string `json:"password"`
	} `json:"errors"`
}

func handleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	m := Login{}
	_ = json.NewDecoder(r.Body).Decode(&m)
	hasEmail := m.Email != ""
	hasPassword := m.Password != ""
	// Input validation
	if !hasEmail || !hasPassword {
		loginError := LoginError{Errors: struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}{Email: "", Password: ""}}
		if !hasEmail {
			loginError.Errors.Email = "Please provide an email address"
		}
		if !hasPassword {
			loginError.Errors.Password = "Please provide a password"
		}
		log.Printf("handleLogin: %+v", loginError)
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(loginError)
		return
	}
	// Email validation
	_, err := mail.ParseAddress(m.Email)
	if err != nil {
		loginError := LoginError{Errors: struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}{Email: "Not a valid email address", Password: ""}}
		log.Printf("handleLogin: %+v", loginError)
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(loginError)
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, _ = fmt.Fprint(writer, "Welcome!\n")
	})
	router.POST("/login", handleLogin)

	log.Fatal(http.ListenAndServe(":8080", router))
}
