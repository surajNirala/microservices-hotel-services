package commons

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponseOrder struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponseOrder struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"` // Make error field optional
}

func ResponseSuccess(c *gin.Context, status int, message string, data interface{}) {
	res := SuccessResponseOrder{
		Status:  status,
		Message: message,
	}
	if data != nil {
		res.Data = data
	}
	c.JSON(status, res)
}

// ResponseError formats error API responses
func ResponseError(c *gin.Context, status int, message string, optionalErr ...error) {
	res := ErrorResponseOrder{
		Status:  status,
		Message: message,
	}
	if len(optionalErr) > 0 && optionalErr[0] != nil {
		res.Error = optionalErr[0] // Set error only if it's not nil
	}
	c.JSON(status, res)
	c.Abort()
}
