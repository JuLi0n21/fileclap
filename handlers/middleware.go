package handlers

import (
	"context"
	"net/http"
	"time"

	"log"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/JuLi0n21/fileclap/utils"
)

var Users = map[string]*models.User{}

var AuthCookie string = "fileclap_session_cookie"

func Wrapper(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Auth(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		var ctx context.Context
		var u *models.User
		var cookie http.Cookie

		if cookie, err := r.Cookie(AuthCookie); err == nil {

			u = Users[cookie.Value]
		} else if err == http.ErrNoCookie {

			value, err := utils.GenValue(128)
			if err != nil {
				return err
			}

			cookie = &http.Cookie{
				Name:  AuthCookie,
				Value: value, Path: "/",
				Expires: time.Now().Add(24 * time.Hour),
				Secure:  false, HttpOnly: true,
				SameSite: http.SameSiteStrictMode}

			http.SetCookie(w, cookie)
		} else {
			return err
		}

		if u == nil {
			u = models.NewUser("Default User")
			Users[cookie.Value] = u

		}

		uuid := r.PathValue("useruuid")
		if uuid != "" && uuid != u.ID.String() {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return nil
		}

		ctx = context.WithValue(r.Context(), models.UserContext, u)
		r = r.WithContext(ctx)

		if err := next(w, r); err != nil {
			return err
		}

		return nil
	}
}
