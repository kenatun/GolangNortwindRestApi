package order

import "github.com/GolangNortwindRestApi/helper"

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(param *getOrdersRequest) (*OrderList, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOrderById(params *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(params)
}

func (s *service) GetOrders(param *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(param)
	helper.Catch(err)
	totalRecords, err := s.repo.GetTotalOrders(param)
	helper.Catch(err)
	return &OrderList{Data: orders, TotalRecords: totalRecords}, nil
}
