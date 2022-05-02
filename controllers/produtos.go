package controllers

import (
	"app_web_01/models"
	"app_web_01/services"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)

}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(price, 64)
		services.TrataErro(err)

		quantidadeConvertida, err := strconv.Atoi(amount)
		services.TrataErro(err)

		models.CriarNovoProduto(name, description, precoConvertido, quantidadeConvertida)
		// Insert some data into table.
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
