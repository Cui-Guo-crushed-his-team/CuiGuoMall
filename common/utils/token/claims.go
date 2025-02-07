package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const expirationTime = 2 //两小时

type UserClaims struct {
	UserID    string `json:"user_id"`
	UserTrait string `json:"user_trait"`
	UserRole  string `json:"user_role"`
	jwt.RegisteredClaims
}

func NewUserClaims(userID, userTrait, userRole string) *UserClaims {
	return &UserClaims{
		UserID:    userID,
		UserTrait: userTrait,
		UserRole:  userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Cui-Guo",
		},
	}
}
