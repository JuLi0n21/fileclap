package main

import (
	"log"
	"net/http"

	"github.com/JuLi0n21/fileclap/handlers"
	"github.com/JuLi0n21/fileclap/repository"
	"github.com/joho/godotenv"
)

func main() {
	run()
}

func run() error {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	repo, err := repository.NewSQLiteRepository("database.db")
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	log.Fatal(http.ListenAndServe(":8080", handlers.NewServer(repo)))
	return nil

}
