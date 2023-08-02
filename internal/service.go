package internal

import "github.com/GauravMakhijani/ecommerce-products/internal/domain"

type ProductService interface {
	GetProducts() []domain.Product
	GetProductsOfCategory(category string) []domain.Product
	AddProduct(product domain.Product) domain.Product
	GetProductById(id uint64) (domain.Product, error)
	UpdateProductQTY(id uint64, quantity uint64) error
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProducts() []domain.Product {
	return s.repo.GetProducts()
}

func (s *productService) GetProductsOfCategory(category string) []domain.Product {
	return s.repo.GetProductsOfCategory(category)
}

func (s *productService) AddProduct(product domain.Product) domain.Product {
	return s.repo.AddProduct(product)
}

func (s *productService) GetProductById(id uint64) (domain.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) UpdateProductQTY(id uint64, quantity uint64) error {
	err := s.repo.UpdateProductQTY(id, quantity)
	if err != nil {
		return err
	}

	return nil
}
