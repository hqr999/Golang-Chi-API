package livroapi

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

// NewLivroArmazem creates and returns a pointer to a new LivroArmazem instance.
func NewLivroArmazem() *LivroArmazem {
	return &LivroArmazem{
		livros: []*Livro{
			{
				ID:             "1",
				Titulo:         "A Beginners Guide to Mathematical Logic",
				Autor:          "Raymond Smullyan",
				DataPublicacao: "18/06/2014",
				Lingua:         "Inglês",
			},
		},
	}
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
