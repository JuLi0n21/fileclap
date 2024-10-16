package handlers

import (
	"net/http"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/JuLi0n21/fileclap/web"
)

func Index(w http.ResponseWriter, r *http.Request) error {

	cmp := web.Index(models.GetUser(r.Context()).Name)
	cmp.Render(r.Context(), w)
	return nil
}
