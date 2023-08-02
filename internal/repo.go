package internal

import (
	"errors"

	"github.com/GauravMakhijani/ecommerce-products/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() []domain.Product
	GetProductsOfCategory(category string) []domain.Product
	AddProduct(product domain.Product) domain.Product
	GetProductById(id uint64) (domain.Product, error)
	UpdateProductQTY(id uint64, quantity uint64) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) GetProducts() []domain.Product {
	var products []domain.Product
	r.db.Find(&products)
	return products
}

func (r *ProductRepositoryImpl) GetProductsOfCategory(category string) []domain.Product {
	var products []domain.Product
	r.db.Find(&products, "category = ?", category)
	return products
}

func (r *ProductRepositoryImpl) AddProduct(product domain.Product) domain.Product {
	r.db.Create(&product)
	return product
}

func (r *ProductRepositoryImpl) GetProductById(id uint64) (domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return product, errors.New("product not found")
	}
	return product, nil
}

func (r *ProductRepositoryImpl) UpdateProductQTY(id uint64, quantity uint64) error {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return errors.New("product not found")
	}
	product.Quantity = quantity
	r.db.Save(&product)
	return nil
}
