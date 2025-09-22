package livroapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LivroHandler struct {
	Armazem LivroArmazenamento
}

func (ll LivroHandler) ListaLivros(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(ll.Armazem.Lista())
	if err != nil {
		http.Error(w, "Erro Interno", http.StatusInternalServerError)
		return
	}

}

func (ll LivroHandler) GetLivros(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	livro := ll.Armazem.Get(id)
	if livro == nil {
		http.Error(w, "Livro Não Encontrado", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(livro)
	if err != nil {
		http.Error(w, "Erro Interno", http.StatusInternalServerError)
		return
	}
}

func (ll LivroHandler) UpdateLivro(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var livro Livro
	err := json.NewDecoder(r.Body).Decode(&livro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	livroAt := ll.Armazem.Update(id, livro)
	if livroAt == nil {
		http.Error(w, "Livro não encontrado", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(livroAt)
	if err != nil {
		http.Error(w, "Erro Interno", http.StatusInternalServerError)
		return
	}

}

func (ll LivroHandler) DeleteLivro(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	livro := ll.Armazem.Delete(id)
	if livro == nil {
		http.Error(w, "Livro não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ll LivroHandler) CreateLivro(w http.ResponseWriter, r *http.Request) {
	var livro Livro
	err := json.NewDecoder(r.Body).Decode(&livro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ll.Armazem.Create(livro)
	err = json.NewEncoder(w).Encode(livro)
	if err != nil {
		http.Error(w, "Erro Interno", http.StatusInternalServerError)
		return
	}

}
