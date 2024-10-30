package main

import (
	"github.com/gin-gonic/gin"
	"go-api-rest/apirestgolang/controller"
	"go-api-rest/apirestgolang/db"
	"go-api-rest/apirestgolang/repository"
	"go-api-rest/apirestgolang/useCase"
)

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepository(dbConnection)
	productUseCase := useCase.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUseCase)

	router := gin.Default()
	router.GET("/product", productController.GetProducts)
	router.POST("/product", productController.SaveProduct)
	router.GET("/product/:id", productController.GetProductById)
	router.DELETE("/product/:id", productController.DeleteProduct)
	router.PUT("/product/:id", productController.EditProduct)

	err = router.Run(":8000")
	if err != nil {
		return
	}
}
