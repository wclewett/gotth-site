package handlers

import (
  "simple-serv/internal/templates"
  "net/http"
)

type LoginHandler struct{}

func NewGetLoginHandler() *LoginHandler {
  return &LoginHandler{}
}

func (h *LoginHandler) Serve(w http.ResponseWriter, r *http.Request) {
	c := templates.Login()
	err := templates.App(c, "Acceder").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
