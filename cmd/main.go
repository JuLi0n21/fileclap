package main

import (
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
		return err
	}

	repo, err := repository.NewSQLiteRepository("database.db")
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	http.ListenAndServe(":8080", handlers.NewServer(repo))
	return nil

}
