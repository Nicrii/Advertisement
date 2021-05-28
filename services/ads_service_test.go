package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
	"net/http"
	"testing"
)

var (
	adDaoMock     adsDaoMock
	getAdFunction func(adId int64) (*domain.Ad, *utils.ApplicationError)
)

type adsDaoMock struct{}

func (m *adsDaoMock) GetAd(adId int64) (*domain.Ad, *utils.ApplicationError) {
	return getAdFunction(adId)
}

func init() {
	domain.AdDao = &adDaoMock
}

func TestGetAdNotFoundInDatabase(t *testing.T) {
	getAdFunction = func(adId int64) (*domain.Ad, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "ad 0 not found",
		}
	}
	ad, err := AdsService.GetAd(0)
	assert.Nil(t, ad)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, "ad 0 not found", err.Message)
}

func TestGetAdNoError(t *testing.T) {
	getAdFunction = func(adId int64) (*domain.Ad, *utils.ApplicationError) {
		return &domain.Ad{
			Id:          123,
			Name:        "TestAdvertisement",
			Description: "This is test date",
			ImgURL:      []string{"link1", "link2"},
			Price:       100500,
			Date:        date.Date{Year: 2021},
		}, nil
	}
	ad, err := AdsService.GetAd(0)
	assert.Nil(t, err)
	assert.NotNil(t, ad)
	assert.EqualValues(t, ad.Id, uint64(123))
	assert.EqualValues(t, ad.Name, "TestAdvertisement")
	assert.EqualValues(t, ad.Description, "This is test date")
	assert.EqualValues(t, ad.ImgURL, []string{"link1", "link2"})
	assert.EqualValues(t, ad.Price, 100500)
	assert.EqualValues(t, ad.Date, date.Date{Year: 2021})
}
