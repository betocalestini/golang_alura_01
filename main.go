package main

import (
	env "app_web_01/enviroment"
	"app_web_01/routes"
	"fmt"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	fmt.Println("Servidor rodando na porta", env.Use("PORT"))
	err := http.ListenAndServe(env.Use("PORT"), nil)

	if err != nil {
		fmt.Println(err)
	}
}
