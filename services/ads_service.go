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
	result := &domain.Ad{Id: adId}
	if err := result.GetAd(); err != nil {
		return nil, err
	}
	return result, nil
}
