package handlers

import (
  "simple-serv/internal/templates"
  "net/http"
)

type HolaHandler struct{}

func NewHolaHandler() *HolaHandler {
  return &HolaHandler{}
}

func (h *HolaHandler) Serve(w http.ResponseWriter, r *http.Request) {
	c := templates.Hola()
	err := templates.App(c, "Fabio Marsanz").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
