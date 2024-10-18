package handlers

import (
	"net/http"

	"github.com/JuLi0n21/fileclap/repository"
)

func NewServer(repo *repository.Repository) http.Handler {
	h := NewHandler(repo)
	r := http.NewServeMux()

	r.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	r.HandleFunc("GET /{$}", Wrapper(Auth(h.Index)))
	r.HandleFunc("GET /u/{useruuid}/settings", Wrapper(Auth(h.Settings)))

	return r
}
