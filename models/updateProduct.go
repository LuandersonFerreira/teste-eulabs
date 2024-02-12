package models

import "teste-eulabs/db"

func UpdateProduct(produto Produto) error {
	conn, err := db.ConnectDb()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `UPDATE produtos SET nome = ?, preco = ? WHERE id = ?`

	_, err = conn.Exec(query, produto.Nome, produto.Preco, produto.Id)
	if err != nil {
		return err
	}

	return nil
}
