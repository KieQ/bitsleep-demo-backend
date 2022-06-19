package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnSuccess(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, gin.H{"status_code":0, "status_message":"","data": data})
}

func OnFailure(c *gin.Context, statusCode int64, statusMessage string){
	c.JSON(http.StatusOK, gin.H{"status_code": statusCode, "status_message":statusMessage, "data":nil})
}