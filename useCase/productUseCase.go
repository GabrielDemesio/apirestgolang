package useCase

import (
	"go-api-rest/apirestgolang/model"
	"go-api-rest/apirestgolang/repository"
)

// Interface que define os métodos do caso de uso
type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
}

// Implementação concreta do caso de uso
type ProductUseCaseImpl struct {
	repository repository.ProductRepository
}

func NewProductUseCase(rep repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		repository: rep,
	}
}

func (p *ProductUseCaseImpl) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}
