package main

import (
	"net/http"
	"simple-serv/internal/handlers"
)

func main() {

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  
  http.HandleFunc("/", handlers.NewHolaHandler().Serve)
  http.HandleFunc("/login", handlers.NewLoginHandler().Serve)

  http.ListenAndServe(":8080", nil)
}
