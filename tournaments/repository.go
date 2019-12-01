package tournaments

import (
	"database/sql"

	"github.com/go-kit/kit/log"
)

// Repository _
type Repository interface {
}

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository _
func NewRepository(db *sql.DB, l log.Logger) Repository {
	return &repository{
		db:     db,
		logger: l,
	}
}

func (r *repository) SaveTournament(data Tournament) (t *Tournament, err error) {
	return
}
