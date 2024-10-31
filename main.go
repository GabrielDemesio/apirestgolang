package main

import (
	"apirestgo/controller"
	"apirestgo/db"
	"apirestgo/repository"
	"apirestgo/useCase"
	"github.com/gin-gonic/gin"
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
	router.GET("/product/name/:name", productController.GetProductByName)
	router.POST("/product", productController.SaveProduct)
	router.GET("/product/:id", productController.GetProductById)
	router.DELETE("/product/:id", productController.DeleteProduct)
	router.PUT("/product/:id", productController.EditProduct)

	err = router.Run(":8000")
	if err != nil {
		return
	}
}
