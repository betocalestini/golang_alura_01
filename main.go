package main

import (
	"app_web_01/databases"
	env "app_web_01/enviroment"
	"fmt"
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "camiseta",
			Descricao:  "azul",
			Preco:      39.50,
			Quantidade: 5},
		{
			Nome:       "tenis",
			Descricao:  "Confortavel",
			Preco:      259.00,
			Quantidade: 40,
		},
		{"Fone", "esportivo", 199.0, 5},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}

func main() {
	databases.ConectaComBancoDeDados()

	http.HandleFunc("/", index)
	fmt.Println("Servidor rodando na porta", env.Use("PORT"))
	err := http.ListenAndServe(env.Use("PORT"), nil)

	if err != nil {
		fmt.Println(err)
	}
}
