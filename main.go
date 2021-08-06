package main

import (
	"log"
	"net/http"
	"os"
	"users-service/handlers"
	"users-service/models"

	"github.com/gorilla/mux"
	"github.com/medymik/configo/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	// Connect Db
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error %v", err.Error())
	}
	// Auto Migrate
	db.AutoMigrate(&models.User{})
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
