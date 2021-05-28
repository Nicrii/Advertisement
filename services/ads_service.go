package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
	"net/http"
)

type adsService struct {
}

var (
	AdsService adsService
)

func (a *adsService) Get(adId int64, fields []string) (*domain.GetResponse, *utils.ApplicationError) {
	result, err := domain.Get(adId, fields)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adsService) Create(ad domain.Ad) (int64, int) {
	if status := ad.Validate(); status != http.StatusOK {
		return 0, status
	}

	return ad.Save()
}
