package listing

import "errors"

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	GetBeer(int) (*Beer, error)
	// GetBeers returns all beers saved in storage.
	GetBeers() ([]*Beer, error)
}

// Service provides beer and review listing operations.
type Service interface {
	GetBeer(int) (*Beer, error)
	GetBeers() ([]*Beer, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) GetBeers() ([]*Beer, error) {
	return s.r.GetBeers()
}

// GetBeer returns a beer
func (s *service) GetBeer(id int) (*Beer, error) {
	return s.r.GetBeer(id)
}
