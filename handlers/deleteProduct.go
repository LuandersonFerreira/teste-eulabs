package handlers

import (
	"fmt"
	"net/http"
	"teste-eulabs/models"

	"github.com/labstack/echo/v4"
)

func DeleteProduct(e echo.Context) error {
	productId := e.Param("id")

	err := models.DeleteProduct(productId)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Produto deletado com sucesso! ID: %s", productId)
	return e.String(http.StatusOK, message)
}
