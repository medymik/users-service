package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"users-service/models"

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

func (u *User) register(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		u.l.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(usr)
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		u.register(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
