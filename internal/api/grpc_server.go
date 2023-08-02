package api

import (
	"context"

	"github.com/GauravMakhijani/ecommerce-products/internal"
	"github.com/GauravMakhijani/ecommerce-utils/productspb"
)

type server struct {
	productspb.UnimplementedProductServiceServer
	service internal.ProductService
}

// GetProductById(context.Context, *GetProductByIdRequest) (*GetProductByIdResponse, error)
// UpdateProductQTY(context.Context, *UpdateProductQTYRequest) (*UpdateProductQTYResponse, error)
func NewGrpcServer(service internal.ProductService) productspb.ProductServiceServer {
	return &server{
		service: service,
	}
}

func (s *server) GetProductById(ctx context.Context, req *productspb.GetProductByIdRequest) (*productspb.GetProductByIdResponse, error) {
	product, err := s.service.GetProductById(req.GetProductId())
	if err != nil {
		return nil, err
	}
	return &productspb.GetProductByIdResponse{
		Product: &productspb.Product{
			Id:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category,
			Quantity: product.Quantity,
		},
	}, nil
}

func (s *server) UpdateProductQTY(ctx context.Context, req *productspb.UpdateProductQTYRequest) (*productspb.UpdateProductQTYResponse, error) {
	err := s.service.UpdateProductQTY(req.GetProductId(), req.GetQuantity())
	if err != nil {
		return &productspb.UpdateProductQTYResponse{
			Success: false,
		}, err
	}
	return &productspb.UpdateProductQTYResponse{
		Success: true,
	}, nil
}
