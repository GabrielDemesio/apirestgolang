package repository

import (
	"go-api-rest/apirestgolang/model"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	SaveProduct(product model.Product) error
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

func (pr *ProductRepositoryImpl) SaveProduct(product model.Product) error {
	if err := pr.connection.Table("product").Create(&product).Error; err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return err
	}
	return nil
}
