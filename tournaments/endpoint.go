package tournaments

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addTournamentRequest struct {
	Name string `json:"name,omitempty"`
}

type addTournamentResponse struct {
	Err error `json:"error,omitempty"`
}

func (r addTournamentResponse) error() error { return r.Err }

func makeAddTournamentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addTournamentRequest)
		data := Tournament{Name: req.Name}
		_, err := s.NewTournament(data)
		return addTournamentResponse{Err: err}, nil
	}
}
