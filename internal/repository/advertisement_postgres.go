package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"marketplace/internal/models"
)

type AdvertPostgres struct {
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (p *AdvertPostgres) Create(input models.Advert) (models.Advert, error) {
	var output models.Advert
	query := fmt.Sprintf("INSERT INTO %s (title, text, image, price, posting_date, owner) VALUES ($1, $2, $3, $4, $5, $6) "+
		"RETURNING title, text, image, price, posting_date, owner", advertisementsTable)
	row := p.db.QueryRow(query, input.Title, input.Text, input.Image, input.Price, input.PostingDate, input.Owner)
	if err := row.Scan(&output.Title, &output.Text, &output.Image, &output.Price, &output.PostingDate, &output.Owner); err != nil {
		return models.Advert{}, err
	}
	return output, nil
}

func (p *AdvertPostgres) GetAll(login string, params models.AdvertParams) ([]models.AdvertOutput, error) {
	var adverts []models.AdvertOutput
	query := fmt.Sprintf("SELECT title, text, image, price, owner, (owner = $1) as is_owner FROM %s "+
		"WHERE price >= $2 AND price <= $3 ORDER BY %s %s LIMIT $4 OFFSET $5", advertisementsTable, params.Sort, params.Direction)
	err := p.db.Select(&adverts, query, login, params.PriceMin, params.PriceMax, params.Limit, (params.Page-1)*params.Limit)
	if err != nil {
		return nil, err
	}
	return adverts, nil
}
