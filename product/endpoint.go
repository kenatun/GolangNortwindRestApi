package product

import (
	"context"

	"github.com/GolangNortwindRestApi/helper"
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

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct{}

// @Summary Producto by Id
// @Tags Producto
// @Accept json
// @Produce json
// @Param id path int true "Producto Id"
// @Success 200 {object} product.Product "ok"
// @Router /products/{id} [get]
func makeProductByIdEndpoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		product, err := s.GetProductById(&req)
		helper.Catch(err)
		return product, nil
	}
	return getProductByIdEndpoint
}

// @Summary Lista de Productos
// @Tags Producto
// @Accept json
// @Produce json
// @Param request body product.getProductRequest true "User Data"
// @Success 200 {object} product.ProductList "ok"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductRequest)
		result, err := s.Getproducts(&req)
		helper.Catch(err)
		return result, nil
	}
	return getProductsEndPoint
}

// @Summary Insertar Productos
// @Tags Producto
// @Accept json
// @Produce json
// @Param request body product.getAddProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [post]
func MakeAddproductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		helper.Catch(err)
		return productId, nil
	}
	return addProductEndPoint
}

// @Summary Update Productos
// @Tags Producto
// @Accept json
// @Produce json
// @Param request body product.updateProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [put]
func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return r, nil
	}
	return updateProductEndPoint
}

// @Summary Eliminar Productos
// @Tags Producto
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {integer} int "ok"
// @Router /products/{id} [delete]
func makeDeleteProductendPoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteProductEndPoint
}

// @Summary Best Sellers
// @Tags Producto
// @Accept json
// @Produce json
// @Success 200 {object} product.ProductTopResponse "ok"
// @Router /products/bestSellers [get]
func makeBestSellersEndPoint(s Service) endpoint.Endpoint {
	getBestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		helper.Catch(err)
		return result, nil
	}
	return getBestSellersEndpoint
}
