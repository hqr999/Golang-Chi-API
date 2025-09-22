package livroapi

import (
	"encoding/json"
	"net/http"
)

type LivroHandler struct{}

func (ll LivroHandler) ListaLivros(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listLivros())
	if err != nil {
		http.Error(w, "Erro Interno", http.StatusInternalServerError)
		return
	}

}

func (ll LivroHandler) GetLivros(w http.ResponseWriter, r *http.Request)   {}
func (ll LivroHandler) UpdateLivro(w http.ResponseWriter, r *http.Request) {}
func (ll LivroHandler) DeleteLivro(w http.ResponseWriter, r *http.Request) {}
func (ll LivroHandler) CreateLivro(w http.ResponseWriter, r *http.Request) {}
