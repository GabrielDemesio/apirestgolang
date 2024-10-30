package repository

import (
	"errors"
	"go-api-rest/apirestgolang/model"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	SaveProduct(product model.Product) error
	DeleteProduct(productID int) error
	GetProductById(productID int) (model.Product, error)
	EditProduct(product model.Product) error
}

type ProductRepositoryImpl struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{connection: connection}
}
func (pr *ProductRepositoryImpl) GetProducts() ([]model.Product, error) {
	var productList []model.Product

	if err := pr.connection.Table("product").Find(&productList).Error; err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return nil, err
	}

	return productList, nil
}
func (pr *ProductRepositoryImpl) GetProductById(productID int) (model.Product, error) {
	var product model.Product
	if err := pr.connection.Table("product").First(&product, "id = ?", productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Erro ao buscar produto: %v", err)
			return model.Product{}, err
		}
		log.Printf("Erro ao buscar produto: %v", err)
		return model.Product{}, err
	}
	return product, nil
}
func (pr *ProductRepositoryImpl) SaveProduct(product model.Product) error {
	if err := pr.connection.Table("product").Create(&product).Error; err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return err
	}
	return nil
}
func (pr *ProductRepositoryImpl) DeleteProduct(productID int) error {
	if err := pr.connection.Table("product").Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		log.Printf("Erro ao deletar produto com ID %d: %v", productID, err)
		return err
	}
	return nil
}
func (pr *ProductRepositoryImpl) EditProduct(product model.Product) error {
	var existingProduct model.Product
	if err := pr.connection.Table("product").First(&existingProduct, product.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Produto não encontrado: ID %d", product.ID)
			return gorm.ErrRecordNotFound
		}
		log.Printf("Erro ao buscar produto: %v", err)
		return err
	}
	if err := pr.connection.Table("product").Save(&product).Error; err != nil {
		log.Printf("Erro ao editar produto: %v", err)
		return err
	}

	return nil
}
