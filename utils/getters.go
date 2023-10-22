package utils

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type errorResponse struct {
	Message any `json:"message"`
	Status  int `json:"status"`
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

func MakeGinResponse(c *gin.Context, statusCode int, value any) {
	if statusCode == 200 && reflect.TypeOf(value).String() != "string" {
		c.JSON(200, value)
		return
	}
	c.JSON(statusCode, errorResponse{Status: statusCode, Message: value})
	return
}
