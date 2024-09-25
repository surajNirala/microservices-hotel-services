package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/surajNirala/hotel_services/app/controllers/API"
)

func ApiRoutes(apiRouter *gin.Engine) {
	route := apiRouter.Group("/api")
	{
		route.GET("/hotels", api.HotelList)
		route.POST("/hotel/store", api.HotelStore)
		route.GET("/hotel/:hotel_id", api.HotelDetail)
		route.PUT("/hotel/:hotel_id", api.HotelUpdate)
		route.DELETE("/hotel/:hotel_id", api.HotelDelete)
	}
}
