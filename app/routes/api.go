package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/surajNirala/hotel_services/app/controllers/API"
)

func ApiRoutes(apiRouter *gin.Engine) {
	route := apiRouter.Group("/api")
	{
		route.GET("/hotels", api.HotelList)
		route.POST("/hotels/store", api.HotelStore)
		route.GET("/hotels/:hotel_id", api.HotelDetail)
		route.PUT("/hotels/:hotel_id", api.HotelUpdate)
		route.DELETE("/hotels/:hotel_id", api.HotelDelete)
	}
}
