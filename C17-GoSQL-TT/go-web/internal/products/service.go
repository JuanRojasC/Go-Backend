package products

import (
	"context"
)

type Service interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	GetByName(name string) ([]Product, error)
	Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	UpdateField(id int, fields ...ProductField) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

// Get all data from repositorys
func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Get one resource from repository with id match
func (s *service) GetOne(id int) (Product, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// Get one resource from repository with name match
func (s *service) GetByName(name string) ([]Product, error) {
	products, err := s.repository.GetByName(name)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

// Save a new resource in repository
func (s *service) Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	product, err := s.repository.Save(name, color, price, stock, code, published)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// Update a resource in repository
func (s *service) Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	product, err := s.repository.Update(ctx, id, name, color, price, stock, code, published)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// Update field of resource
func (s *service) UpdateField(id int, fields ...ProductField) (Product, error) {
	product, err := s.repository.Patch(id, fields...)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// Delete a resource in repository
func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Return a Service Interface
func NewService(r Repository) Service {
	return &service{r}
}
