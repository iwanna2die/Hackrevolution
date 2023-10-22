package profile

import (
	"HackRevolution/models"
	"HackRevolution/pkg/postgres"
	"HackRevolution/pkg/postgres/requests/user"
	"HackRevolution/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Tags Profile GetAll
// @Description getmyprofile
// @ID getprofile
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/validate/getprofile [get]
func GetProfile(c *gin.Context) {
	value, errs := c.Get("id")
	if errs == false {
		utils.MakeGinResponse(c, 500, errs)
		return
	}
	rows, err := postgres.DB().SDB.Queryx(user.UserAll, value)
	status, param := postgres.ReadDbAll(rows, err, &models.UserInfo{})
	utils.MakeGinResponse(c, status, param)
}
