package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
)

type adsService struct {
}

var (
	AdsService adsService
)

func (a *adsService) GetAd(adId int64) (*domain.Ad, *utils.ApplicationError) {

	ad, err := domain.AdDao.GetAd(adId)
	if err != nil {
		return nil, err
	}

	return ad, nil
}
