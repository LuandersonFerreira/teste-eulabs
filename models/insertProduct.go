package models

import "teste-eulabs/db"

func InsertProduct(produto Produto) error {
	conn, err := db.ConnectDb()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `INSERT INTO produtos (id, nome, preco) VALUES (?, ?, ?)`

	_, err = conn.Exec(sql, produto.Id, produto.Nome, produto.Preco)
	if err != nil {
		return err
	}

	return nil
}
