package main

import "net/http"

func hello() string {
	return "Hello world!"
}

func handleLogin(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/login", handleLogin)
}
