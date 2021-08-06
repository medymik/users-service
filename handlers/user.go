package main

import (
	"log"
	"net/http"
)

type User struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *User {
	return &User{
		l: l,
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
