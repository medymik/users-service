package handlers

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	l  *log.Logger
	db *gorm.DB
}

func NewUser(l *log.Logger, db *gorm.DB) *User {
	return &User{
		l:  l,
		db: db,
	}
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
