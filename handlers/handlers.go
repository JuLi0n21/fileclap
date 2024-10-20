package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/JuLi0n21/fileclap/repository"
	"github.com/JuLi0n21/fileclap/web"
)

type Server struct {
	Repo *repository.Repository
}

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
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return err
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if u, err := s.Repo.UserRepository.LoginUser(username, password); err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			_ = u
			fmt.Println("Login successful")
		}

	} else {
		cmp := web.Login()
		cmp.Render(r.Context(), w)
	}
	return nil
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return err
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		confirmpassword := r.FormValue("confirmpassword")

		if confirmpassword != password {
			return errors.New("Password mismatch")
		}

		email := r.FormValue("email")

		if u, err := s.Repo.UserRepository.RegisterUser(username, email, password); err != nil {

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			_ = u
			return nil
		} else {
			//failed login
		}

		fmt.Fprintf(w, "Username: %s\n", username)
		fmt.Fprintf(w, "Password: %s\n", password)

	} else {

		cmp := web.Register()
		cmp.Render(r.Context(), w)
	}

	return nil
}
