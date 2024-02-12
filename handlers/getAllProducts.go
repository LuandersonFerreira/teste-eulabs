package handlers

import (
	"encoding/json"
	"net/http"
	"teste-eulabs/models"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(e echo.Context) error {
	produtos, err := models.GetAllProducts()
	if err != nil {
		return err
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	e.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(e.Response()).Encode(produtos)
}
