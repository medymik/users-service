package main

import (
	"log"
	"net/http"
	"os"
	"users-service/handlers"

	"github.com/gorilla/mux"
	"github.com/medymik/configo/env"
)

func registerHandlers(r *mux.Router, l *log.Logger) {
	r.Handle("/api/users", handlers.NewUser(l))
}

func main() {
	// New logger
	l := log.New(os.Stdout, "users-service", log.LstdFlags)
	// Load env from file
	env := env.NewEnv(".env")
	env.Load()
	// New mux
	r := mux.NewRouter()
	// register handlers
	registerHandlers(r, l)
	// Server
	s := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}
