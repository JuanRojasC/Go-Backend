package products

// b
type Service interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error)
}

// c
type service struct {
	repository Repository
}

// e
func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// e
func (s *service) GetOne(id int) (Product, error) {
	product, err := s.repository.GetOne(id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// e
func (s *service) Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	product, err := s.repository.Save(name, color, price, stock, code, published)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// d
func NewService(r Repository) Service {
	return &service{r}
}
