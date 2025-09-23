package main

import (
	"fmt"
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

	fmt.Printf("Começando o servidor na porta: %d",3000)
	err := http.ListenAndServe(":3000", rot)
	if err != nil {
			panic(err)
	}
}

func LivrosRotas() chi.Router {
	r := chi.NewRouter()
	armazem := livroapi.NewLivroArmazem()
	livroHandler := livroapi.LivroHandler{
		Armazem: armazem,
	}
	r.Get("/", livroHandler.ListaLivros)
	r.Post("/", livroHandler.CreateLivro)
	r.Get("/{id}", livroHandler.GetLivros)
	r.Put("/{id}", livroHandler.UpdateLivro)
	r.Delete("/{id}", livroHandler.DeleteLivro)
	return r
}
