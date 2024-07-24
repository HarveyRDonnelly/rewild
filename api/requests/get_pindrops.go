package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type GetPindropsRequest struct {
	Delta           float64 `json:"delta"`
	CentreLongitude float64 `json:"centre_longitude"`
	CentreLatitude  float64 `json:"centre_latitude"`
}

type GetPindropsResponse []entities.Pindrop

func getPindropsRoute(r *gin.Engine) *gin.Engine {
	r.GET("/pindrop/", func(c *gin.Context) {

		var requestBody GetPindropsRequest
		var pindrops []entities.Pindrop

		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err)
		}

		// Retrieve project info
		pindropsDBResponse := db.GetPindrops(
			DB,
			db.GetPindropsDBRequest{
				Delta:           requestBody.Delta,
				CentreLongitude: requestBody.CentreLongitude,
				CentreLatitude:  requestBody.CentreLatitude,
			},
		)

		for _, pindropDBResponse := range pindropsDBResponse.Pindrops {
			pindrops = append(pindrops, db.ConstructPindrop(
				DB,
				pindropDBResponse,
			))
		}

		c.JSON(
			http.StatusOK,
			GetPindropsResponse(pindrops),
		)

	})

	return r
}
