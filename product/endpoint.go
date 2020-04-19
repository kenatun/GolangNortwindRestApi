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
type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
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

func MakeAddproductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return productId, nil
	}
	return addProductEndPoint
}
