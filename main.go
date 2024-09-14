package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/proyectogolang/cmd/server/handlers/login"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/",login.Login)
	r.Post("/pepe", login.Register)
	http.ListenAndServe(":3000", r)
}