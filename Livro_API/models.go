package livroapi

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type LivroArmazenamento interface {
	Lista() []*Livro
	Get(string) *Livro
	Update(string, Livro) *Livro
	Create(Livro)
	Delete(string) *Livro
}

type LivroArmazem struct {
	livros []*Livro
}

type Livro struct {
	ID             string `json:"id"`
	Titulo         string `json:"title"`
	Autor          string `json:"autor"`
	DataPublicacao string `json:"data_publicacao"`
	Lingua         string `json:"lingua_original"`
}

//caminhoAbsoluto resolve o caminho de um arquivo de forma segura.
func caminhoAbsoluto(caminho string) string {
	//Pega o caminho do arquivo que está sendo executado	
	_,nome_arq,_,ok := runtime.Caller(1)
		
		if !ok {
				panic("Falha ao obter o caminho do arquivo atual")
		}
		//Junta o diretório do arquivo atual com o caminho do arquivo desejado 
	return filepath.Join(filepath.Dir(nome_arq),caminho)

}

// NewLivroArmazem cria e retorna um ponteiro para uma mova instância de LivroArmazem.
func NewLivroArmazem() *LivroArmazem {
	//1. Primeiro vamos abrir o arquivo JSON
	arquivo, err := os.Open(caminhoAbsoluto("../livros.json"))
	if err != nil {
		log.Fatalf("Falha ao abrir o arquivo JSON: %v", err)
	}
	defer arquivo.Close()

	//2. Decodifique o conteúdo do arquivo no tipo []*Livro
	var livros []*Livro
	decoder := json.NewDecoder(arquivo)
	err = decoder.Decode(&livros)
	if err != nil {
		log.Fatalf("Falha ao decodificar o JSON: %v", err)
	}

	//3. Retorna o armazém com os livros carregados
	return &LivroArmazem{livros: livros}
}

func (ll *LivroArmazem) Lista() []*Livro {
	return ll.livros
}

func (ll *LivroArmazem) Get(id string) *Livro {
	for _, value := range ll.livros {
		if value.ID == id {
			return value
		}
	}
	return nil
}

func (ll *LivroArmazem) Create(livro Livro) {
	ll.livros = append(ll.livros, &livro)
}

func (ll *LivroArmazem) Delete(id string) *Livro {
	for i, liv := range ll.livros {
		if liv.ID == id {
			ll.livros = append(ll.livros[:i], (ll.livros)[i+1:]...)
			return &Livro{}
		}
	}
	return nil
}

func (ll *LivroArmazem) Update(id string, livroAtualizado Livro) *Livro {
	for i, livro := range ll.livros {
		if livro.ID == id {
			ll.livros[i] = &livroAtualizado
			return livro
		}
	}
	return nil
}
