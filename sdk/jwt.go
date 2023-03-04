package sdk

import (
	"InternBCC/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func Token(payload entity.User, c *gin.Context) (string, error) {
	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Email or Password",
		})
		return "", nil
	}
	return tokenString, nil
}
