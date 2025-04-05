package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Sub        string           `json:"sub"`
	Sid        string           `json:"sid"`
	Email      string           `json:"email"`
	Name       string           `json:"name"`
	VerifiedAt *jwt.NumericDate `json:"vyf"`
	jwt.RegisteredClaims
}
