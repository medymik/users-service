package main

import (
	"fmt"
	"os"

	"github.com/medymik/configo/env"
)

func main() {
	env := env.NewEnv(".env")
	env.Load()
	fmt.Println(os.Getenv("DB_NAME"))
}
