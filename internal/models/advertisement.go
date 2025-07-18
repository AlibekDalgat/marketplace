package models

import (
	"errors"
	"fmt"
	"marketplace/pkg/validation"
	"time"
)

type Advert struct {
	Title       string    `json:"title" binding:"required"`
	Text        *string   `json:"text"`
	Image       *string   `json:"image"`
	Price       *float32  `json:"price" binding:"required"`
	PostingDate time.Time `json:"posting_date"`
	Owner       string    `json:"owner"`
}

func (advert Advert) Validate() error {
	err := ""
	if len(advert.Title) == 0 || advert.Price == nil {
		err += "Объявление должно иметь название и цену\n"
	}
	if len(advert.Title) > 100 {
		err += "Длина названия не должна превышать 100\n"
	}
	if advert.Text != nil && len(*advert.Text) > 1000 {
		err += "Длина текста не должна превышать 1000\n"
	}
	if advert.Text != nil && len(*advert.Text) > 1000 {
		err += "Длина текста не должна превышать 1000\n"
	}
	if advert.Image != nil {
		if errImage := validation.ValidImage(*advert.Image); errImage != nil {
			err += errImage.Error() + "\n"
		}
	}
	if *advert.Price < 0 {
		err += "Цена должна быть положительной\n"
	}
	if err == "" {
		return nil
	}
	return errors.New(err)
}

type AdvertParams struct {
	Sort      string
	Direction string
	Limit     int
	Page      int
	PriceMin  int
	PriceMax  int
}

func (params AdvertParams) Validate() error {
	err := ""
	if params.Sort != "posting_date" && params.Sort != "price" {
		err += fmt.Sprintf("Некорректный тип сортировки: %s\n", params.Sort)
	}
	if params.Direction != "desc" && params.Direction != "asc" {
		err += fmt.Sprintf("Некорректное направление сортировки: %s\n", params.Direction)
	}
	if params.Page < 1 {
		err += fmt.Sprintf("Несуществующая страница: %s\n", params.Page)
	}
	if params.Limit < 1 {
		err += fmt.Sprintf("Некорректный размер страницы: %s\n", params.Limit)
	}
	if params.PriceMin > params.PriceMax || params.PriceMin < 0 || params.PriceMax < 0 {
		err += fmt.Sprintf("Некорректные фильтры. priceMin: %d, priceMax: %d\n", params.PriceMin, params.PriceMax)
	}
	if err == "" {
		return nil
	}
	return errors.New(err)
}

type AdvertOutput struct {
	Title   string  `db:"title" json:"title"`
	Text    *string `db:"text" json:"text"`
	Image   *string `db:"image" json:"image"`
	Price   float32 `db:"price" json:"price"`
	Owner   string  `db:"owner" json:"owner"`
	IsOwner bool    `db:"is_owner" json:"isOwner"`
}
