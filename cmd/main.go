package main

import (
	"net/http"

	"github.com/JuLi0n21/fileclap/handlers"
	"github.com/joho/godotenv"
)

func main() {
	run()
}

func run() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	http.ListenAndServe(":8080", handlers.NewServer())
	return nil

}
