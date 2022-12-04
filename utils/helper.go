package utils

import (
	"github.com/gin-gonic/gin"
)

// SendJSONResponse formats response to JSON
func SendJSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// SendJSONError formats response to JSON
func SendJSONError(c *gin.Context, statusCode int, err error) {
	SendJSONResponse(c, statusCode, gin.H{
		"error": err.Error(),
	})
}
