package domain ///TODO: прочитать про DAO
import (
	"fmt"
	"github.com/Nicrii/Advertisement/datasources/postgresql/ads_db"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/lib/pq"
	"net/http"
)

const (
	queryGetAd = "SELECT * FROM advertisement WHERE id=$1;"
)

func (ad *Ad) GetAd() *utils.ApplicationError {
	stmt, err := ads_db.Client.Prepare(queryGetAd)
	if err != nil {
		return &utils.ApplicationError{Message: err.Error(), StatusCode: http.StatusInternalServerError, Code: "server_error"}
	}
	defer stmt.Close()

	result := stmt.QueryRow(ad.Id)
	if err := result.Scan(&ad.Id, &ad.Name, &ad.Description, pq.Array(&ad.ImagesURLs), &ad.Price, &ad.Date); err != nil {
		return &utils.ApplicationError{Message: fmt.Sprintf("error when trying to get ad %d %s", ad.Id, err.Error()), StatusCode: http.StatusInternalServerError, Code: "server_error"}
	}

	return nil
}
