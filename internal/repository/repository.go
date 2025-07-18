package repository

import (
	"github.com/jmoiron/sqlx"
	"marketplace/internal/models"
)

type Authorization interface {
	CreateUser(login models.User) (models.User, error)
	GetUser(login, password string) (models.User, error)
}

type Advertisement interface {
	Create(input models.Advert) (models.Advert, error)
	GetAll(login string, params models.AdvertParams) ([]models.AdvertOutput, error)
}

type Repository struct {
	Authorization
	Advertisement
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Advertisement: NewAdvertPostgres(db),
	}
}
