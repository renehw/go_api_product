package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renehw/go_api_product/controller"
	"github.com/renehw/go_api_product/db"
	"github.com/renehw/go_api_product/repository"
	"github.com/renehw/go_api_product/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.Connect()

	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.DELETE("/product/:productId", ProductController.DeleteProductById)

	server.Run(":8000")

}
