package models

import (
	"teste-eulabs/db"
)

func DeleteProduct(id string) error {
	conn, err := db.ConnectDb()
	if err != nil {
		return err
	}

	query := `DELETE FROM produtos WHERE id = ?;`
	_, err = conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
