package product

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
	Getproducts(params *getProductRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
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

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellers()
	if err != nil {
		panic(err)
	}
	totalVentas, err := s.repo.GetTotalVentas()
	if err != nil {
		panic(err)
	}
	return &ProductTopResponse{Data: products, TotalVentas: totalVentas}, err
}
