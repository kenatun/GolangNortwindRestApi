package order

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
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
