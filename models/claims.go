package models

import (
	"github.com/dgrijalva/jwt-go"
)

// change to user
type Claims struct {
	Student User
	jwt.StandardClaims
}
