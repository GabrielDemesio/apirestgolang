package useCase

import (
	"go-api-rest/apirestgolang/model"
	"go-api-rest/apirestgolang/repository"
)

type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
	SaveProduct(product model.Product) (model.Product, error)
}

type ProductUseCaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		repo: repo,
	}
}

func (uc *ProductUseCaseImpl) GetProducts() ([]model.Product, error) {
	return uc.repo.GetProducts()
}
func (uc *ProductUseCaseImpl) SaveProduct(product model.Product) (model.Product, error) {
	if err := uc.repo.SaveProduct(product); err != nil {
		return model.Product{}, err
	}
	return product, nil
}
