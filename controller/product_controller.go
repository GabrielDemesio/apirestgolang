package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-api-rest/apirestgolang/model"
	"go-api-rest/apirestgolang/useCase"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (p *ProductController) GetProductById(ctx *gin.Context) {
	// Pega o productID como string e tenta convertê-lo para int
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto é obrigatório"})
		return
	}

	// Converte o productID para int
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto deve ser um número válido"})
		return
	}

	// Chama o use case para buscar o produto usando o ID inteiro
	product, err := p.productUseCase.GetProductById(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o produto"})
		return
	}

	// Retorna o produto encontrado com o status 200 OK
	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
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

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	// Obtém o productID como string e tenta convertê-lo para int
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto é obrigatório"})
		return
	}

	// Converte o productID para int
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto deve ser um número válido"})
		return
	}

	// Chama o use case para deletar o produto usando o ID inteiro
	err = p.productUseCase.DeleteProduct(productID)
	if err != nil {
		// Verifica se o erro foi de produto não encontrado
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		// Se houver outro erro, retorna um erro interno
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o produto"})
		return
	}

	// Retorna o status 204 No Content se o produto foi excluído com sucesso
	ctx.Status(http.StatusNoContent)
}
