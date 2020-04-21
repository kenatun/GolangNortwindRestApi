package order

import (
	"context"

	"github.com/GolangNortwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getOrderByIdRequest struct {
	orderId int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}
type addOrderRequest struct {
	OrderDate    string
	CustomerId   int
	OrderDetails []addOrderDetailRequest
}

type addOrderDetailRequest struct {
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
}

func makeGetOrderByIdEndPoint(s Service) endpoint.Endpoint {
	getOrderByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)
		result, err := s.GetOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrderByIdEndPoint
}

func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	getOrdersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrdersEndPoint
}

func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	addOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.InsertOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return addOrderEndpoint
}
