package order

import (
	"context"

	"github.com/GolangNortwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getOrderByIdRequest struct {
	orderId int64
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
