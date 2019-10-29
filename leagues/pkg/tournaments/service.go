package tournaments

// Service is the users's service interface
type Service interface {
	NewTournament(Tournament) (*Tournament, error)
}

type service struct {
	repository Repository
}

// NewService creates a new users's service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) NewTournament(data Tournament) (t *Tournament, err error) {
	return
}
