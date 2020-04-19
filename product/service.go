package product

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
	Getproducts(params *getProductRequest) (*ProductList, error)
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

func (s *service) Getproducts(params *getProductRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}
