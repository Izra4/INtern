package model

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Regist struct {
	Nama        string `json:"nama" binding:"required"`
	Email       string `json:"email" binding:"email,required"`
	Password    string `json:"password" binding:"required"`
	Passconfirm string `json:"passconfirm" binding:"required"`
	Number      string `json:"number" binding:"required"`
}

type LogIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserClaims struct {
	ID uint `json:"id"` // bebas dahh mau pake 'Subject' di registered claims juga bebas
	jwt.RegisteredClaims
}

// exp example: (time.Hour * 1), etc
func NewUserClaims(id uint, exp time.Duration) UserClaims {
	return UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
