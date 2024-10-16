package handlers

import "net/http"

func NewServer() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /", Wrapper(Auth(Index)))

	return r
}
