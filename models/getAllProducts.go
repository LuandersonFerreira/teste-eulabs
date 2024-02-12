package models

import (
	"teste-eulabs/db"
)

func GetAllProducts() (produtos []Produto, err error) {
	conn, err := db.ConnectDb()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM produtos;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var produto Produto
		if err = rows.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Preco); err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}
