package tournaments

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errBadRoute = errors.New("bad route")
var errUnknown = errors.New("unknown tournament")
var errInvalidArgument = errors.New("invalid argument")

// MakeHandler returns a handler for the booking service.
func MakeHandler(s Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	tournamentsAdd := kithttp.NewServer(
		makeAddTournamentEndpoint(s),
		decodeAddTournamentRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/tournaments/v1", tournamentsAdd).Methods("POST")

	return r
}

func decodeAddTournamentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return addTournamentRequest{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errUnknown:
		w.WriteHeader(http.StatusNotFound)
	case errInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
