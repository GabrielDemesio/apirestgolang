package mapper

import (
	"apirestgo/dto"
	"apirestgo/model"
)

func ToProductResponse(product model.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          int(product.ID),
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}
}

func ToProduct(productRequest dto.ProductRequest) model.Product {
	return model.Product{
		Name:        productRequest.Name,
		Price:       productRequest.Price,
		Description: productRequest.Description,
	}
}
