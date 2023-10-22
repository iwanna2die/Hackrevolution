package server

import (
	_ "HackRevolution/docs"
	"HackRevolution/internal/auth"
	"HackRevolution/internal/lenta"
	"HackRevolution/internal/middleware"
	"HackRevolution/internal/profile"
	"HackRevolution/internal/test"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Add_handlers(s *Server) {
	apiV1 := s.Group("/api/v1/")
	auths := apiV1.Group("/auth/")
	auths.POST("/login", auth.LoginPage)
	auths.GET("/logout", auth.Logout)
	auths.GET("/refresh", auth.Refresh)
	adm := apiV1.Group("adm").Use(middleware.Authenticate("Преподаватель"))
	adm.GET("/lentacours/:id", lenta.GetCourseOtchet)
	tovalidate := apiV1.Group("/validate").Use(middleware.Authenticate(""))
	tovalidate.GET("/getprofile", profile.GetProfile)
	tovalidate.GET("/lenta", lenta.LentaGet)
	tovalidate.GET("/note/:id", lenta.GetNote)
	tovalidate.GET("/lentacours/:id", lenta.GetLentaCourse)
	s.GET("/test", test.Test)
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
