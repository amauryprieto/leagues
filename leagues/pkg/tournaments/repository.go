package tournaments

import "database/sql"

type Repository interface {
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveTournament(data Tournament) (t *Tournament, err error) {
	return
}
