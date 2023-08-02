package api

import (
	"encoding/json"
	"net/http"

	"github.com/GauravMakhijani/ecommerce-products/internal"
	"github.com/GauravMakhijani/ecommerce-products/internal/domain"
)

func getProductsController(service internal.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		var products []domain.Product

		if category == "" {
			products = service.GetProducts()
		} else {
			products = service.GetProductsOfCategory(category)
		}

		json.NewEncoder(w).Encode(products)
	}
}

func AddProductsController(service internal.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.Product

		json.NewDecoder(r.Body).Decode(&product)

		product = service.AddProduct(product)
		json.NewEncoder(w).Encode(product)
		w.WriteHeader(http.StatusCreated)
	}
}
