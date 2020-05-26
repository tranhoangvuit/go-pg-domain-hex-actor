package adding

import "errors"

var ErrDuplicate = errors.New("beer already exists")

// Repository provides access to beer repository.
type Repository interface {
}

// Service provides beer adding operations.
type Service interface {
}

type service struct {
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
