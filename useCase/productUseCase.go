package useCase

import (
	"errors"
	"go-api-rest/apirestgolang/model"
	"go-api-rest/apirestgolang/repository"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
	SaveProduct(product model.Product) (model.Product, error)
	GetProductById(id int) (model.Product, error)
	DeleteProduct(id int) error
	EditProduct(product model.Product) (model.Product, error)
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
func (uc *ProductUseCaseImpl) GetProductById(id int) (model.Product, error) {
	if id == 0 {
		return model.Product{}, errors.New("product ID is required")
	}
	product, err := uc.repo.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	return product, nil
}
func (uc *ProductUseCaseImpl) DeleteProduct(id int) error {
	if id == 0 {
		return errors.New("product ID is required")
	}
	if err := uc.repo.DeleteProduct(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}
func (uc *ProductUseCaseImpl) EditProduct(product model.Product) (model.Product, error) {
	if product.ID == 0 {
		return model.Product{}, errors.New("product id is required")
	}
	_, err := uc.repo.GetProductById(product.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	if err := uc.repo.EditProduct(product); err != nil {
		return model.Product{}, err
	}

	return product, nil
}
