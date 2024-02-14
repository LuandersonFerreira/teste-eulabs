package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teste-eulabs/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func InsertProduct(e echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(e.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}
	//json_map has the JSON Payload decoded into a map
	nome, ok := json_map["nome"].(string)
	if ok {
		return fmt.Errorf("O campo Nome é obrigatório! ")
	}
	preco, ok := json_map["preco"].(float64)
	if ok {
		return fmt.Errorf("O campo Preço é obrigatório! ")
	}

	produto := models.Produto{
		Id:    uuid.New().String(),
		Nome:  nome,
		Preco: preco,
	}

	err = models.InsertProduct(produto)
	if err != nil {
		fmt.Println("erro: ", err)
		return err
	}
	message := fmt.Sprintf("Produto inserido com sucesso! ID: %s", produto.Id)
	return e.String(http.StatusOK, message)
}
