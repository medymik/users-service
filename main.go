package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/medymik/configo/env"
)

func main() {
	// Load env from file
	env := env.NewEnv(".env")
	env.Load()
	// New mux
	r := mux.NewRouter()
	// Server
	s := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}
	s.ListenAndServe()
}
