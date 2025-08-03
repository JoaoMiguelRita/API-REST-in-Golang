package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repositories
	ProductRepository := repository.NewProductRepository(dbConection)

	// Camada de services
	ProductService := service.NewProductService(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
