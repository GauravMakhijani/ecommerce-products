package api

import (
	"net/http"

	"github.com/GauravMakhijani/ecommerce-products/internal"
	"github.com/gorilla/mux"
)

func Setup(service internal.ProductService) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/products", getProductsController(service)).Methods(http.MethodGet)
	router.HandleFunc("/products", AddProductsController(service)).Methods(http.MethodPost)

	return
}
