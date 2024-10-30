package controller

import (
	"apirestgo/model"
	"apirestgo/useCase"
	"errors"
	"github.com/gin-gonic/gin"
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
			"error":   "Error to find product",
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}
	product, err := p.productUseCase.GetProductById(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error to find the product"})
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
			"error":   "Error to create product",
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}
	err = p.productUseCase.DeleteProduct(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error to delete product"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (p *ProductController) EditProduct(ctx *gin.Context) {
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product is required"})
		return
	}

	product.ID = productID

	updatedProduct, err := p.productUseCase.EditProduct(product)
	if err != nil {
		if err.Error() == "Product not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error to update the product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated with success", "data": updatedProduct})
}
