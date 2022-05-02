package routes

import (
	"app_web_01/controllers"
	db "app_web_01/databases"
	"net/http"
)

func CarregaRotas() {
	db.ConexaoInicial()

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
}
