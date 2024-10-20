package handlers

import (
	"net/http"

	"github.com/JuLi0n21/fileclap/repository"
)

func NewServer(repo *repository.Repository) http.Handler {
	h := NewHandler(repo)
	r := http.NewServeMux()

	r.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	r.HandleFunc("/{$}", Wrapper(Auth(h.Index)))
	r.HandleFunc("/login", Wrapper(Auth(h.Login)))
	r.HandleFunc("/register", Wrapper(Auth(h.Register)))
	r.HandleFunc("GET /u/{useruuid}/settings", Wrapper(Auth(h.Settings)))

	return r
}
