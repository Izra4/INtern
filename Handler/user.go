package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	//get name, email, number, password
	var get model.Regist
	if err := c.ShouldBindJSON(&get); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon lengkapi input Anda", err)
		return
	}
	if get.Password != get.Passconfirm {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password tidak sama",
		})
		return
	}
	//Hashing
	hash, err := sdk.Hashing(get.Password)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to Hash", err)
		return
	}
	//Create
	user := entity.User{
		Model:    gorm.Model{},
		Nama:     get.Nama,
		Email:    get.Email,
		Password: hash,
		Number:   get.Number,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to create", err)
		return
	}

	//Respond
	sdk.Success(c, http.StatusOK, "Success to Register", user)
}

func LogIn(c *gin.Context) {
	var body model.LogIn
	if err := c.ShouldBindJSON(&body); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Error to read", err)

		return
	}

	//cari data
	var req entity.User
	database.DB.First(&req, "email = ?", body.Email)
	if req.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Invalid Email / Password",
		})
		return
	}

	err := sdk.ValidateHash(req.Password, body.Password)
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Failed to compare the password", err)
		return
	}
	tokenString, err := sdk.Token(req)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}
	sdk.Success(c, http.StatusOK, "User berhasil log in", map[string]string{"token": tokenString})
}

func LogOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	sdk.Success(c, http.StatusOK, "Anda telah log out", "")
}

func Validate(c *gin.Context) {
	id, _ := c.Get("user")
	claims := id.(model.UserClaims)

	user := entity.User{}

	err := database.DB.First(&user, claims.ID)
	if err.Error != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err.Error)
		return
	}

	c.JSON(200, gin.H{
		"data":    user.Nama,
		"error":   nil,
		"message": "logged in",
	})
}
