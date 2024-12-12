package main

import (
	"apirestgo/controller"
	"apirestgo/db"
	_ "apirestgo/docs"
	"apirestgo/repository"
	"apirestgo/useCase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Product API
// @version 1.0
// @description API for managing products.
// @host localhost:8000
// @BasePath /
func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepository(dbConnection)
	productUseCase := useCase.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUseCase)

	router := gin.Default()

	// Rotas da API
	router.GET("/product", productController.GetProducts)
	router.GET("/product/name/:name", productController.GetProductByName)
	router.POST("/product", productController.SaveProduct)
	router.GET("/product/:id", productController.GetProductById)
	router.DELETE("/product/:id", productController.DeleteProduct)
	router.PUT("/product/:id", productController.EditProduct)

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = router.Run(":8000")
	if err != nil {
		return
	}
}
