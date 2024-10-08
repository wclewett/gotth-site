package handlers

import (
	"simple-serv/internal/templates"
	"net/http"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) Serve(w http.ResponseWriter, r *http.Request) {
	c := templates.NotFound()
	err := templates.App(c, "Not Found").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

