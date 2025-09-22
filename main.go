package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	livroapi "github.com/hqr999/Golang-Chi-Api/Livro_API"
)

func main() {
	rot := chi.NewRouter()
	rot.Use(middleware.Logger)
	rot.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("OK"))
	})

	rot.Mount("/livros", LivrosRotas())

	http.ListenAndServe(":3000", rot)
}

func LivrosRotas() chi.Router {
	r := chi.NewRouter()
	livroHandler := livroapi.LivroHandler{}
	r.Get("/", livroHandler.GetLivros)
	r.Post("/", livroHandler.CreateLivro)
	r.Get("/{id}", livroHandler.GetLivros)
	r.Put("/{id}", livroHandler.UpdateLivro)
	r.Delete("/{id}", livroHandler.DeleteLivro)
	return r
}
