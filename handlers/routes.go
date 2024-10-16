package handlers

import "net/http"

func NewServer() http.Handler {
	r := http.NewServeMux()

	r.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	r.HandleFunc("GET /", Wrapper(Auth(Index)))

	return r
}
