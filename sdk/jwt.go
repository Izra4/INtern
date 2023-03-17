package sdk

import (
	"InternBCC/entity"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func Token(payload entity.User) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}
	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.NewUserClaims(payload.ID, exp))

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, KEY string) (string, error) {

	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // method used to sign the token
		if !ok {
			// wrong signing method
			return "", errors.New("wrong signing method")
		}
		return []byte(KEY), nil
	})

	if err != nil {
		// parse failed
		return "", fmt.Errorf("token has been tampered with")
	}

	if !token.Valid {
		// token is not valid
		return "", fmt.Errorf("invalid token")
	}

	return signedToken, nil
}
