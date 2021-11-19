package products

import (
	entities "products-api/entities/product"
	"time"
)

type Service interface {
	GetAllProducts(query string) ([]entities.Product, error)
	CreateProduct(input entities.ProductInput) (entities.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllProducts(query string) ([]entities.Product, error) {
	products, err := s.repo.FindAll(query)

	if err != nil {
		return products, err
	}

	return products, err
}

func (s *service) CreateProduct(input entities.ProductInput) (entities.Product, error) {
	time := time.Now()

	var product = entities.Product{
		CreatedAt:   time,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		Quantity:    input.Quantity,
	}

	createProduct, err := s.repo.CreateProduct(product)

	if err != nil {
		return createProduct, err
	}

	return createProduct, nil
}
