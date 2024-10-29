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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar produtos",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto é obrigatório"})
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto deve ser um número válido"})
		return
	}
	product, err := p.productUseCase.GetProductById(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o produto"})
		return
	}
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
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto é obrigatório"})
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto deve ser um número válido"})
		return
	}
	err = p.productUseCase.DeleteProduct(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o produto"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (p *ProductController) EditProduct(ctx *gin.Context) {
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto é obrigatório"})
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto deve ser um número válido"})
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos para o produto"})
		return
	}

	product.ID = productID

	updatedProduct, err := p.productUseCase.EditProduct(product)
	if err != nil {
		if err.Error() == "Produto não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o produto"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso", "data": updatedProduct})
}
