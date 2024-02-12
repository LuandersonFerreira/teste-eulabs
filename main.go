package main

import (
	"teste-eulabs/configs"
	"teste-eulabs/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	port := configs.GetServerPort()

	e := echo.New()
	e.POST("/products", handlers.InsertProduct)
	e.GET("/products", handlers.GetAllProducts)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)
	e.Logger.Fatal(e.Start(port))
}
