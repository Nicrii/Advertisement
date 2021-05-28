package domain ///TODO: прочитать про DAO
import (
	"github.com/Nicrii/Advertisement/datasources/postgresql/ads_db"
	"github.com/Nicrii/Advertisement/utils"
	"net/http"
)

var (
	//ads = map[int64]*Ad{
	//	123: &Ad{Id: 123, Name: "TestAdvertisement", Description: "This is test date", ImagesURLs: []string{"link1", "link2"}, Price: 100500, Date: date.Date{Year: 2021}},
	//}
	//AdDao adsDaoInterface
	adsDb = make(map[int64]*Ad)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING id;"
	queryGetUser    = "SELECT * FROM users WHERE id=$1;"
	queryUpdateUser = "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	queryDeleteUser = "DELETE FROM users WHERE id=$1;"
)

//func init() {
//	AdDao = &adDao{}
//}
//
//type adDao struct {
//}
//
//type adsDaoInterface interface {
//	GetAd(adId int64) (*Ad, *utils.ApplicationError)
//}

func (ad *Ad) GetAd() *utils.ApplicationError {
	if err := ads_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := adsDb[ad.Id]
	if result == nil {
		return &utils.ApplicationError{Message: "ad not found", StatusCode: http.StatusNotFound, Code: "not_found"}
	}
	ad.Name = result.Name
	ad.Description = result.Description
	ad.ImagesURLs = result.ImagesURLs
	ad.Price = result.Price
	ad.Date = result.Date

	return nil
}
