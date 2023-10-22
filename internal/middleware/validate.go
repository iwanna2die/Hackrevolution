package middleware

import (
	"HackRevolution/internal/authentification"
	"HackRevolution/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Authenticate(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		fmt.Println("token:" + token)
		if err != nil {
			utils.MakeGinResponse(c, 401, "unauthorized")
			c.Abort()
			return
		}

		claims, errs := authentification.ValidateToken(token, os.Getenv("ACCESS_TOKEN_SECRET"))

		if errs != "" {
			utils.MakeGinResponse(c, 401, "not authorized")
			c.Abort()
			return

		}
		if role == "" || role == claims.Student.Role[0].NameTypeStudent {
			c.Set("role", claims.Student.Role)
			c.Set("name", claims.Student.Student.FullName)
			c.Set("id", claims.Student.Student.IdStudent)
			c.Next()
		} else {
			utils.MakeGinResponse(c, 401, "access denied")
			c.Abort()
			return
		}

	}
}
