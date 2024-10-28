package controller

import (
	"github.com/gin-gonic/gin"
	"go-api-rest/apirestgolang/model"
	"go-api-rest/apirestgolang/useCase"
	"net/http"
)

type ProductController struct {
	productUseCase useCase.ProductUseCase
}

func NewProductController(useCase useCase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: useCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		// Retorna uma mensagem de erro mais clara
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar produtos",
			"message": err.Error(),
		})
		return
	}
	// Retorna a lista de produtos em caso de sucesso
	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (p *ProductController) SaveProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := p.productUseCase.SaveProduct(product)
	if err != nil {
		// Retorna uma mensagem de erro mais clara se falhar ao salvar
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar produto",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": createdProduct,
	})
}
