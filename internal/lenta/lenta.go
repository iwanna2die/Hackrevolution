package lenta

import (
	"HackRevolution/models"
	"HackRevolution/pkg/postgres"
	"HackRevolution/pkg/postgres/requests/lnt"
	"HackRevolution/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get Lenta
// @Tags Lenta GetAll
// @Description getmylenta
// @ID getmylenta
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/validate/lenta [get]
func LentaGet(c *gin.Context) {
	value, errs := c.Get("id")
	if errs == false {
		utils.MakeGinResponse(c, 500, errs)
		return
	}
	rows, err := postgres.DB().SDB.Queryx(lnt.GetCurrentLnt, value)
	status, param := postgres.ReadDbAll(rows, err, &models.LentaUser{})
	utils.MakeGinResponse(c, status, param)
}

// @Summary Get Note by id klass
// @Tags Note Getid
// @Description getmynote
// @ID getmynote
// @Accept  json
// @Produce  json
// @Param id path string true "uuid"
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/validate/note/{id} [get]
func GetNote(c *gin.Context) {
	id, errs := utils.CheckUUIDParameter(c.Param("id"))
	if errs != "" {
		utils.MakeGinResponse(c, 500, errs)
		return
	}

	value, errrs := c.Get("id")
	if errrs == false {
		utils.MakeGinResponse(c, 500, errs)
		return
	}
	rows, err := postgres.DB().SDB.Queryx(lnt.GetCurrentNode, value, id)
	status, param := postgres.ReadDbAll(rows, err, &models.Notes{})
	utils.MakeGinResponse(c, status, param)
}

// @Summary Get Lenta by id klass
// @Tags Lenta Getid
// @Description getmyLenta
// @ID getmyLenta
// @Accept  json
// @Produce  json
// @Param id path string true "uuid"
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/validate/lentacours/{id} [get]
func GetLentaCourse(c *gin.Context) {
	id, errs := utils.CheckUUIDParameter(c.Param("id"))
	if errs != "" {
		utils.MakeGinResponse(c, 500, errs)
		return
	}

	value, errrs := c.Get("id")
	if errrs == false {
		utils.MakeGinResponse(c, 500, errs)
		return
	}
	if utils.CheckExist(lnt.CheckExistUserToCourse, value, id) {
		rows, err := postgres.DB().SDB.Queryx(lnt.GetAllZapis, id)
		status, param := postgres.ReadDbAll(rows, err, &models.Zapis{})
		utils.MakeGinResponse(c, status, param)
	} else {
		utils.MakeGinResponse(c, 404, "Вы не в этом курсе!")
	}

}

// @Summary Get Otchet by id klass
// @Tags Otchet Getid
// @Description getmyOtchet
// @ID Otchet
// @Accept  json
// @Produce  json
// @Param id path string true "uuid"
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/adm/lentacours/{id} [get]
func GetCourseOtchet(c *gin.Context) {
	id, errs := utils.CheckUUIDParameter(c.Param("id"))
	if errs != "" {
		utils.MakeGinResponse(c, 500, errs)
		return
	}
	rows, err := postgres.DB().SDB.Queryx(lnt.SendOtchet, id)
	status, param := postgres.ReadDbAll(rows, err, &models.Ocenka{})
	utils.MakeGinResponse(c, status, param)
}
