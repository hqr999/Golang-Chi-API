package livroapi

type Livro struct {
	ID             string `json:"id"`
	Titulo         string `json:"title"`
	Autor          string `json:"autor"`
	DataPublicacao string `json:"data_publicacao"`
	Lingua         string `json:"lingua_original"`
}

var livros = []*Livro{
	{
		ID:             "1",
		Titulo:         "A Beginners Guide to Mathematical Logic",
		Autor:          "Raymond Smullyan",
		DataPublicacao: "18/06/2014",
		Lingua:         "Inglês",
	},
}

func listLivros() []*Livro {
	return livros
}

func getLivro(id string) *Livro {
	for _, value := range livros {
		if value.ID == id {
			return value
		}
	}
	return nil
}

func armazenaLivro(livro Livro) {
	livros = append(livros, &livro)
}

func deletaLivro(id string) *Livro {
	for i, liv := range livros {
		if liv.ID == id {
			livros = append(livros[:i], (livros)[i+1:]...)
			return &Livro{}
		}
	}
	return nil
}

func atualizaLivro(id string, livroAtualizado Livro) *Livro {
	for i, livro := range livros {
		if livro.ID == id {
			livros[i] = &livroAtualizado
			return livro
		}
	}
	return nil
}
