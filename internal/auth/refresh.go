package auth

import (
	"HackRevolution/internal/authentification"
	"HackRevolution/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// @Summary Refresh
// @Tags Auth.Refresh
// @Description Refresh
// @ID Refresh
// @Accept  json
// @Produce  json
// @Success 200
// @Posted 201
// @Failure 400,404 {object} utils.errorResponse
// @Failure 500 {object} utils.errorResponse
// @Failure default {object} utils.errorResponse
// @Router /api/v1/auth/refresh [get]
func Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh")
	if err != nil {
		utils.MakeGinResponse(c, 401, "not authorized")
		return
	}
	errs, token := authentification.RefreshToken(rt)
	if errs != "" {
		utils.MakeGinResponse(c, 401, errs)
		return
	}

	claims, errs := authentification.ValidateToken(token, os.Getenv("ACCESS_TOKEN_SECRET"))
	if errs != "" {
		utils.MakeGinResponse(c, http.StatusInternalServerError, errs)
		c.Abort()
		return
	}
	c.SetCookie("id_usr", claims.Student.Student.IdStudent.String(), int(claims.ExpiresAt), "/", "10.200.164.12", false, true)
	c.SetCookie("token", token, int(claims.ExpiresAt), "/", "10.200.164.12", false, true)
	utils.MakeGinResponse(c, 200, "refreshed")
	return
}
