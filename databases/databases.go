package databases

import (
	env "app_web_01/enviroment"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	host     = "localhost"
// 	database = "app_web_go"
// 	user     = "root"
// 	password = "admin"
// )

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConectaComBancoDeDados() {

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", env.Use("DB_USER"), env.Use("PASSWORD"), env.Use("HOST"), env.Use("DB_PORT"), env.Use("DATABASE"))

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Create table.
	_, err = db.Exec("CREATE TABLE if not exists produtos (id serial PRIMARY KEY, name VARCHAR(50), description VARCHAR(50), price VARCHAR(50), amount INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")
}

func TesteEnv() {
	fmt.Println(env.Use("USER"))
	fmt.Println(env.Use("PASSWORD"))
	fmt.Println(env.Use("HOST"))
	fmt.Println(env.Use("DATABASE"))
}
