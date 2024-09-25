package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/hotel_services/app/commons"
	"github.com/surajNirala/hotel_services/app/config"
	"github.com/surajNirala/hotel_services/app/models"
	"github.com/surajNirala/hotel_services/app/validation"
)

var DB = config.DB

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type userResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HotelList(c *gin.Context) {

	var hotels []models.Hotel
	DB.Select("id", "name", "user_id").Order("created_at DESC").Find(&hotels)
	commons.ResponseSuccess(c, 200, "Get all hotel list.", hotels)
	return
}

func HotelStore(c *gin.Context) {
	var request models.Hotel
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Request binding error: %v", err)
		customErrors := validation.TranslateValidationErrors(err)
		res := gin.H{
			"status":  400,
			"message": "Invalid Request",
			"errors":  customErrors,
		}
		c.JSON(400, res)
		return
	}

	hoteldata := models.Hotel{
		Name:   request.Name,
		UserID: request.UserID,
	}

	if err := DB.Create(&hoteldata).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Hotel not created successfully",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}
	res := gin.H{
		"status":  201,
		"message": "Hotel Created Successfully",
		"data":    hoteldata,
	}
	c.JSON(201, res)
}

func HotelDetail(c *gin.Context) {
	var hotel models.Hotel
	hotel_id := c.Param("hotel_id")
	result := DB.Select("id", "name", "user_id").Where("id = ?", hotel_id).Find(&hotel)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Hotel not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}
	res := Response{
		Status:  200,
		Message: "Fetch Hotel Detail",
		Data:    hotel,
	}
	c.JSON(200, res)
}

func HotelUpdate(c *gin.Context) {
	var hotel models.Hotel
	hotel_id := c.Param("hotel_id")
	result := DB.Select("id", "name", "user_id").Where("id = ?", hotel_id).Find(&hotel)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Hotel not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}
	var request models.Hotel
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Request binding error: %v", err)
		customErrors := validation.TranslateValidationErrors(err)
		res := gin.H{
			"status":  400,
			"message": "Invalid Request",
			"errors":  customErrors,
		}
		c.JSON(400, res)
		return
	}
	hoteldata := models.Hotel{
		Name:   request.Name,
		UserID: request.UserID,
	}

	if err := DB.Where("id = ?", hotel_id).Updates(&hoteldata).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Hotel is not updated.",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}

	res := Response{
		Status:  200,
		Message: "Hotel Detail Updated Successfully.",
		Data:    hoteldata,
	}
	c.JSON(200, res)
}

func HotelDelete(c *gin.Context) {
	var hotel models.Hotel
	hotel_id := c.Param("hotel_id")
	result := DB.Select("id", "name", "user_id").Where("id = ?", hotel_id).Find(&hotel)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Hotel not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}

	if err := DB.Where("id = ?", hotel_id).Delete(&hotel).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Hotel is not deleted.",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}

	res := Response{
		Status:  200,
		Message: "Hotel Deleted Successfully.",
		Data:    nil,
	}
	c.JSON(200, res)
}
