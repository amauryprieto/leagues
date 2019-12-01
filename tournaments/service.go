package tournaments

import (
	"github.com/go-kit/kit/log"
)

// Service is the users's service interface
type Service interface {
	NewTournament(Tournament) (*Tournament, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService creates a new users's service
func NewService(r Repository, l log.Logger) Service {
	return &service{
		repository: r,
		logger:     l,
	}
}

func (s *service) NewTournament(data Tournament) (t *Tournament, err error) {
	return
}
