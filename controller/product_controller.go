package controller

import (
	"apirestgo/model"
	"apirestgo/useCase"
	"errors"
	"fmt"
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

// @Summary Get all products
// @Description Get the list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Router /product [get]
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

// @Summary Get product by id
// @Description Get a specific product
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Router /product/{id} [get]
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

// @Summary Create a new product
// @Description Create a new product in the system (ID will be auto-generated and cannot be provided)
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data (ID will be ignored)"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /product [post]
func (p *ProductController) SaveProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be provided when creating a new product"})
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

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Delete a product
// @Description Delete the product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} string "Product Deleted"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /product/{id} [delete]
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

// @Summary Update an existing product
// @Description Update an existing product by its ID (ID in body must match ID in URL)
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Updated product data (ID cannot be changed)"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /product/{id} [put]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	if product.ID != 0 && product.ID != productID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID in body does not match ID in URL"})
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

// @Summary Get products by name
// @Description Get the list of a specific product
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Router /product/name [get]
func (p *ProductController) GetProductByName(ctx *gin.Context) {
	productName := ctx.Param("name")
	if productName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product name is required"})
		return
	}

	product, err := p.productUseCase.GetProductByName(productName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("product not found: %s", productName)})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}
