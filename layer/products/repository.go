package products

import (
	entities "products-api/entities/product"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(query string) ([]entities.Product, error)
	CreateProduct(input entities.Product) (entities.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(query string) ([]entities.Product, error) {
	var product []entities.Product

	if err := r.db.Order(query).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) CreateProduct(input entities.Product) (entities.Product, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}
