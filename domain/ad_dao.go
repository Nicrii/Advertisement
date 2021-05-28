package domain ///TODO: прочитать про DAO
import (
	"fmt"
	"github.com/Nicrii/Advertisement/utils"
	"google.golang.org/genproto/googleapis/type/date"
	"net/http"
)

var (
	ads = map[int64]*Ad{
		123: &Ad{Id: 123, Name: "TestAdvertisement", Description: "This is test date", ImgURL: []string{"link1", "link2"}, Price: 100500, Date: date.Date{Year: 2021}},
	}
	AdDao adsDaoInterface
) //ToDo: добавить БД
func init() {
	AdDao = &adDao{}
}

type adDao struct {
}

type adsDaoInterface interface {
	GetAd(adId int64) (*Ad, *utils.ApplicationError)
}

func (a *adDao) GetAd(adId int64) (*Ad, *utils.ApplicationError) {
	if ad := ads[adId]; ad != nil {
		return ad, nil

	}
	return nil, &utils.ApplicationError{Message: fmt.Sprintf("ad %v not found", adId), StatusCode: http.StatusNotFound, Code: "not_found"}
}
