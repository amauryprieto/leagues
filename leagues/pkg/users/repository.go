package users

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

func (r *repository) SaveUser(data User) (t *User, err error) {
	return
}
