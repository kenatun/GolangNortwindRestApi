package product

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIdRequest) (*Product, error) {
	//bussiness logic
	return s.repo.GetProductById(param.ProductID)
}
