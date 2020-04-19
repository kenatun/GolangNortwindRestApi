package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIdRequest struct {
	ProductID int
}

type getProductRequest struct {
	Limit  int
	Offset int
}

func makeProductByIdEndpoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductRequest)
		result, err := s.Getproducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}
