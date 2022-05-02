package models

import (
	db "app_web_01/databases"
	"app_web_01/services"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.Conexao()
	p := Produto{}
	produtos := []Produto{}

	var (
		id, amount        int
		name, description string
		price             float64
	)

	rows, err := db.Query("SELECT * from produtos;")
	services.TrataErro(err)

	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = name
		p.Descricao = description
		p.Preco = price
		p.Quantidade = amount

		produtos = append(produtos, p)

		fmt.Printf("Data row = (%d, %s, %s, %f, %d)\n", id, name, description, price, amount)
	}
	err = rows.Err()
	services.TrataErro(err)
	fmt.Println("Done.")

	defer db.Close()
	return produtos
}

func CriarNovoProduto(name, description string, precoConvertido float64, quantidadeConvertida int) {
	db := db.Conexao()
	sqlStatement, err := db.Prepare("INSERT INTO produtos (name , description , price , amount) VALUES (?, ?, ?, ?);")
	services.TrataErro(err)

	res, err := sqlStatement.Exec(&name, &description, &precoConvertido, &quantidadeConvertida)
	services.TrataErro(err)

	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
	services.TrataErro(err)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.Conexao()
	// Modify some data in table.
	sqlStatement, err := db.Prepare("DELETE FROM produtos WHERE id = ?;")
	services.TrataErro(err)

	rows, err := sqlStatement.Exec(id)
	services.TrataErro(err)

	rowCount, err := rows.RowsAffected()
	services.TrataErro(err)

	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
	defer db.Close()
}

func EditaProduto(id string) Produto {

	db := db.Conexao()
	// Read some data from the table.
	row, err := db.Query("SELECT * from produtos WHERE id=?;", id)
	services.TrataErro(err)

	produtoParaAtualizar := Produto{}
	fmt.Println("Reading data:")
	for row.Next() {
		var (
			id, amount        int
			name, description string
			price             float64
		)
		err := row.Scan(&id, &name, &description, &price, &amount)
		services.TrataErro(err)

		fmt.Printf("Data row = (%d, %s, %s, %f, %d)\n", id, name, description, price, amount)
		produtoParaAtualizar.Nome = name
		produtoParaAtualizar.Descricao = description
		produtoParaAtualizar.Preco = price
		produtoParaAtualizar.Quantidade = amount
	}
	fmt.Println("Done.")
	defer db.Close()
	return produtoParaAtualizar
}
