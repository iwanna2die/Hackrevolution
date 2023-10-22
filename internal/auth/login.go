package auth

import (
	"HackRevolution/internal/authentification"
	models "HackRevolution/models"
	"HackRevolution/pkg/postgres"
	"HackRevolution/pkg/postgres/requests/user_req"
	"HackRevolution/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// @Summary logout
// @Tags Auth.Login
// @Description logout
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/auth/logout [get]
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "10.200.164.12", false, true)
	c.SetCookie("refresh", "", -1, "/api/v1/auth/refresh", "10.200.164.12", false, true)
	c.SetCookie("id_usr", "", -1, "/", "10.200.164.12", false, true)
	utils.MakeGinResponse(c, 200, "logged out")
}

// @Summary Login
// @Tags Auth.Login
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Param logpwd body models.LogPwd true "Login: ""\n password: ""\n"
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/auth/login [post]
func LoginPage(c *gin.Context) {
	var logpwd models.LogPwd
	var worker []models.Student
	var roles []models.RoleStudent
	err := c.BindJSON(&logpwd)
	if err != nil {
		e := fmt.Sprintf("received invalid param which is not json: %v", c.Param("id"))
		log.Println(e)
		utils.MakeGinResponse(c, http.StatusNotFound, e)
		return
	}
	if utils.CheckExist(user_req.ExistLoginPwd, logpwd.Login, logpwd.Password) {
		err := postgres.DB().SDB.Select(&worker, user_req.TakeUserInfo, logpwd.Login, logpwd.Password)
		if err == nil {
			err = postgres.DB().SDB.Select(&roles, user_req.TakeUserRole, worker[0].IdStudent)
			fmt.Println(err)
			if err == nil {
				user := models.MakeUser(worker[0], roles)

				tma, errs := utils.MakeDuration(os.Getenv("ACCESS_TIME"))
				at, e := authentification.CreateToken(user, tma, os.Getenv("ACCESS_TOKEN_SECRET"))

				if errs != nil || e != nil {
					utils.MakeGinResponse(c, http.StatusInternalServerError, "could not generate access token")
					return
				}
				fmt.Println(user)
				tmr, errs := utils.MakeDuration(os.Getenv("REFRESH_TIME"))
				rt, e := authentification.CreateToken(user, tmr, os.Getenv("REFRESH_TOKEN_SECRET"))
				err := authentification.UpdateToken(rt, user.Student.IdStudent)
				fmt.Println("UpdateToken:", err, " CreateToken:", e, "MakeDuration:", errs)
				if err != "" || e != nil || errs != nil {
					utils.MakeGinResponse(c, http.StatusInternalServerError, "could not generate refresh token")
					return
				}

				c.SetCookie("token", at, int(tma), "/", "10.200.164.12", false, true)
				c.SetCookie("refresh", rt, int(tmr), "/api/v1/auth/refresh", "10.200.164.12", false, true)
				c.SetCookie("id_usr", user.Student.IdStudent.String(), int(tma), "/", "10.200.164.12", false, true)
				utils.MakeGinResponse(c, 200, "access")
				return

			}
		} else {
			utils.MakeGinResponse(c, 500, "some database error")
			return
		}

	}
	utils.MakeGinResponse(c, http.StatusBadRequest, "Login or Password is incorrect")

}
