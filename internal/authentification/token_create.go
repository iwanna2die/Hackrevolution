package authentification

import (
	"HackRevolution/models"
	"HackRevolution/pkg/postgres"
	"HackRevolution/pkg/postgres/requests/token_req"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"log"
	"os"
	"strconv"
	"time"
)

func CreateToken(worker models.User, duration time.Duration, secret string) (string, error) {
	claims := models.Claims{
		Student: worker,
		StandardClaims: jwt.StandardClaims{
			Subject:   worker.Student.IdStudent.String(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
		return "", err
	}
	return token, nil
}

func ValidateToken(tokenString string, secret string) (claims *models.Claims, msg string) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		msg = "not valid token"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token expired"
		return
	}
	return claims, msg
}

func UpdateToken(rt string, id uuid.UUID) any {
	_, err := postgres.Changes(1, token_req.CreateToken, "", id, rt, time.Now())
	if err != "" {
		return err
	}
	return ""
}

func RefreshToken(rt string) (any, string) {
	claims, err := ValidateToken(rt, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != "" {
		return err, "err"
	}
	dur, errs := strconv.ParseInt(os.Getenv("ACCESS_TIME"), 10, 64)
	if errs != nil {
		return errs.Error(), "err"
	}
	tm := time.Minute * time.Duration(dur)
	token, errs := CreateToken(claims.Student, tm, os.Getenv("ACCESS_TOKEN_SECRET"))
	if errs != nil {
		return errs.Error(), "err"
	}
	errrs := UpdateToken(rt, claims.Student.Student.IdStudent)
	if errrs != "" {
		return errrs, "err"
	}
	return "", token
}
