package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"marketplace/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (p *AuthPostgres) CreateUser(user models.User) (models.User, error) {
	var output models.User
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) VALUES ($1, $2) RETURNING login, password_hash", usersTable)
	row := p.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&output.Login, &output.Password); err != nil {
		return models.User{}, err
	}
	return output, nil
}

func (p *AuthPostgres) GetUser(login string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT login FROM %s WHERE login=$1", usersTable)
	err := p.db.Get(&user, query, login)
	return user, err
}
