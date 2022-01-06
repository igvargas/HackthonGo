package internal

import "github.com/igvargas/HackthonGo/internal/models"

type ServiceSQL interface {
	Store(customer models.Customer) (models.Customer, error)
	// GetOne(id int) (models.Customer, error)
	// GetAll() ([]models.Customer, error)
}
type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(customer models.Customer) (models.Customer, error) {
	customerCreado, err := s.repository.Store(customer)
	if err != nil {
		return models.Customer{}, err
	}
	return customerCreado, nil
}
