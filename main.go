package main

import (
	"github.com/gin-gonic/gin"
	"go-api-rest/apirestgolang/controller"
	"go-api-rest/apirestgolang/db"
	"go-api-rest/apirestgolang/repository"
	"go-api-rest/apirestgolang/useCase"
	"log"
)

func main() {
	server := gin.Default()

	// Conectar ao banco de dados
	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Criar o repositório
	productRepository := repository.NewProductRepository(dbConnection)

	// Criar o caso de uso
	productUseCase := useCase.NewProductUseCase(productRepository)

	// Criar o controlador
	productController := controller.NewProductController(productUseCase)

	// Rotas
	server.GET("/product", productController.GetProducts)

	// Endpoint de teste
	server.PATCH("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Iniciar o servidor na porta 8000
	if err := server.Run(":8000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
