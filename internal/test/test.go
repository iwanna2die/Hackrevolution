package test

import (
	"HackRevolution/utils"
	"github.com/gin-gonic/gin"
)

// import (
//
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"inventory/utils"
//	"io"
//	"net/http"
//
// )
//
// @Summary TEST
// @Tags  TEST
// @Description TEST
// @ID TEST
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /test [get]
func Test(c *gin.Context) {
	utils.MakeGinResponse(c, 200, "hi there")
}
