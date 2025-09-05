package main

import (
	"fmt"
	"net/http"

	"github.com/Happy-Fy/redirect-service/internal/config"
	"github.com/Happy-Fy/redirect-service/internal/controller"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("Failed to load .env: %s\n", err))
	}
}

func main() {
	cnf, err := config.Configs()
	if err != nil {
		panic(err)
	}

	rh := controller.NewRedirectHandler(cnf)

	http.HandleFunc("/", rh.Handle)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
