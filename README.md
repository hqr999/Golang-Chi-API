# API Livros com Go

Esse projeto é uma implementação de uma API que devolve livros. Esse projeto foi feito usando a ótima biblioteca HTTP chi em Golang.

O programa faz as 4 operações básicas de uma API Rest, também chamado de **CRUD**. Nossa API portanto é capaz de:
    
* CREATE=Cria uma nova instância do livro.
* READ=Lê as saídas dos nossos livros que foram cadastrados.
* UPDATE=Atualiza algum livro se passarmos uma id válida
* Delete=Deleta algum livro se passarmos uma id válida

Nossa aplicação começa já lendo um arquivo .json já com alguns livros.Vocês os verá quando fizer o seu comando **GET**.

# Como Rodar 
Abra dois terminais, um para abrir e deixar rodando o servidor e o outro para.

## No primeiro terminal rode:

    > go run main.go

Cada fez que você fizer um comando http, ele aparecerá enquanto o servidor estiver rodando.

## No segundo terminal rode:

Cada um dos comandos usando o pacote curl faz um dos comandos http, você também pode usar o POSTMAN ou outro visualizador de API.Você pode rodar os seguintes comandos com o curl:

* **GET** 
        
        > curl -s localhost:3000/livros | jq .  
* **POST**

        > curl -X POST \        
        -H "Content-Type: application/json" \
        -d '{"id": "10","title": "Irmãos Karamázov","autor": "Fiodor Dostoievski","data_publicacao": "1883", "lingua_original": "Russo"}' \
        localhost:3000/livros

* **DELETE**

        > curl -X DELETE localhost:3000/livros/{id}

* **PUT** 

        > curl -X POST \        
        -H "Content-Type: application/json" \
        -d '{"id": "número","title": "título","autor": "autor","data_publicacao": "ano", "lingua_original": "língua"}' \
        localhost:3000/livros/{id}


# Referências 
https://go-chi.io/#/

https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm
