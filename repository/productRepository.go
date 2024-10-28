package repository

import (
	"database/sql"
	"go-api-rest/apirestgolang/model"
	"log"
)

// Definição da interface ProductRepository
type ProductRepository interface {
	GetProducts() ([]model.Product, error)
}

// Implementação concreta do repositório
type ProductRepositoryImpl struct {
	connection *sql.DB
}

// Função para criar uma nova instância de ProductRepositoryImpl
func NewProductRepository(connection *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{connection}
}

// Implementação do método GetProducts
func (pr *ProductRepositoryImpl) GetProducts() ([]model.Product, error) {
	query := "SELECT id, productname, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Printf("Erro ao executar a consulta: %v", err)
		return nil, err
	}
	defer rows.Close() // Garantir que as linhas sejam fechadas

	var productList []model.Product

	for rows.Next() {
		var productObj model.Product
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			log.Printf("Erro ao escanear a linha: %v", err)
			return nil, err
		}
		productList = append(productList, productObj)
	}

	// Verificar se houve erros durante a iteração
	if err = rows.Err(); err != nil {
		log.Printf("Erro durante a iteração das linhas: %v", err)
		return nil, err
	}

	return productList, nil
}
