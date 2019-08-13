package main

import (
	"log"
	"net/http"
)

func hello() string {
	return "Hello world!"
}

func handleLogin(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
