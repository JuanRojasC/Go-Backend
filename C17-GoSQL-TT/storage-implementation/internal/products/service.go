package products

type Service interface {
	GetOne(id int) (Product, error)
	GetAll() ([]Product, error)
	GetFullData(id int) (Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id int) (Product, error) {
	ps, err := s.repository.GetOne(id)
	if err != nil {
		return Product{}, err
	}

	return ps, nil
}

/* func (s *service) GetOne(id int) (Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ps, err := s.repository.GetOneWithContext(ctx, id)
	if err != nil {
		return Product{}, err
	}

	return ps, nil
} */

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) GetFullData(id int) (Product, error) {
	ps, err := s.repository.GetFullData(id)
	if err != nil {
		return Product{}, err
	}

	return ps, nil
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	producto, err := s.repository.Store(nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
