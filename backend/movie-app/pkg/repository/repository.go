package repository

import (
	"github.com/DaukaSDU/movie-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user movie.User) (int, error)
	GetUser(username, password string) (movie.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
