package service

import (
	"marketplace/internal/models"
	"marketplace/internal/repository"
	"time"
)

type AdvertService struct {
	repo repository.Advertisement
}

func NewAdvertService(repo repository.Advertisement) *AdvertService {

	return &AdvertService{repo}
}

func (adv *AdvertService) Create(login string, input models.Advert) (models.Advert, error) {
	if err := input.Validate(); err != nil {
		return models.Advert{}, err
	}
	input.Owner = login
	input.PostingDate = time.Now()
	return adv.repo.Create(input)
}

func (adv *AdvertService) GetAll(login string, params models.AdvertParams) ([]models.AdvertOutput, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	adverts, err := adv.repo.GetAll(login, params)
	if err != nil {
		return nil, err
	}
	return adverts, nil
}
