package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"users-service/models"

	"users-service/utils"

	"golang.org/x/crypto/bcrypt"
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

type tokenResponse struct {
	Token string `json:"token"`
}

func (u *User) register(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		u.l.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// crypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 10)
	if err != nil {
		u.l.Println(err.Error())
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	usr.Password = string(hash)
	fmt.Println(usr)
	usr.Password = ""
	// Generate a jwt token
	var tokenResponse tokenResponse
	tokenResponse.Token, err = utils.CreateJWTToken(1)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	// token response
	json.NewEncoder(w).Encode(tokenResponse)
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		u.register(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
