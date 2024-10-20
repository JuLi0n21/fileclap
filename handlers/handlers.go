package handlers

import (
	"fmt"
	"net/http"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/JuLi0n21/fileclap/repository"
	"github.com/JuLi0n21/fileclap/web"
)

type Server struct {
	Repo *repository.Repository
}

// NewHandler creates a new handler with the given repository.
func NewHandler(repo *repository.Repository) *Server {
	return &Server{Repo: repo}
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) error {

	u := models.GetUser(r.Context())

	f, err := s.Repo.FileRepository.GetAllFilesForUser(u)
	if err != nil {
		f = []*models.File{}
	}

	d, err := s.Repo.FileRepository.GetAllFoldersForUser(u)
	if err != nil {
		d = []*models.Folder{}
	}

	cmp := web.Index(u.Name, f, d)

	cmp.Render(r.Context(), w)

	return nil
}

func (s *Server) Settings(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "hi")
	return nil
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) error {

	cmp := web.Login()
	cmp.Render(r.Context(), w)

	return nil
}
