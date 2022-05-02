package db

import (
	env "app_web_01/enviroment"
	"app_web_01/services"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conexao() *sql.DB {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", env.Use("DB_USER"), env.Use("PASSWORD"), env.Use("HOST"), env.Use("DB_PORT"), env.Use("DATABASE"))

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	services.TrataErro(err)

	fmt.Println("Successfully created connection to database.")
	return db
}

func ConexaoInicial() {

	db := Conexao()
	defer db.Close()

	err := db.Ping()
	services.TrataErro(err)
	// Create table.
	_, err = db.Exec("CREATE TABLE if not exists produtos (id serial PRIMARY KEY, name VARCHAR(50), description VARCHAR(50), price VARCHAR(50), amount INTEGER);")
	services.TrataErro(err)
	fmt.Println("Finished creating table.")
}

func TesteEnv() {
	fmt.Println(env.Use("USER"))
	fmt.Println(env.Use("PASSWORD"))
	fmt.Println(env.Use("HOST"))
	fmt.Println(env.Use("DATABASE"))
}
