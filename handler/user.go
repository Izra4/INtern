package handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	//get name, email, number, password
	var get model.Regist
	if err := c.ShouldBindJSON(&get); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon lengkapi input Anda", err)
		return
	}
	if !strings.HasSuffix(get.Email, "@gmail.com") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Email harus berakhiran @gmail.com",
		})
		return
	}
	if (!upperCase(get.Password)) || (!hasNum(get.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password harus mengandung 1 huruf kapital dan 1 angka",
		})
		return
	}
	if (get.Password != get.Passconfirm) || (len(get.Password) < 8) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password tidak sama / kurang dari 8 karakter",
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

func Validate(c *gin.Context) {
	id, _ := c.Get("user")
	claims := id.(model.UserClaims)
	user := entity.User{}

	err := database.DB.First(&user, claims.ID)
	if err.Error != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err.Error)
		return
	}

	sdk.Success(c, http.StatusOK, "Logged In", user)
}

func ChangeNameNumber(c *gin.Context) {
	type changes struct {
		Email string `json:"email"`
		Nama  string `json:"nama"`
		Nomor string `json:"nomor"`
	}
	userId := c.MustGet("user")
	claims := userId.(model.UserClaims)

	var user entity.User
	if err := database.DB.First(&user, claims.ID).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "User not found", err)
		return
	}
	var req changes
	if err := c.BindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Error to Read", err)
		return
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}

	if req.Nomor != "" {
		user.Number = req.Nomor
	}
	if err := database.DB.Save(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}
	sdk.Success(c, http.StatusOK, "Profil telah diperbarui", user)
}

func ChangePass(c *gin.Context) {
	id, _ := c.Get("user")
	claims := id.(model.UserClaims)

	var user entity.User
	if err := database.DB.First(&user, claims.ID).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "User not found", err)
		return
	}
	var req struct {
		OldPass string `json:"old_pass" binding:"required"`
		NewPass string `json:"new_pass" binding:"required"`
		Confirm string `json:"confirm" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon lengkapi datanya", err)
		return
	}
	if err := sdk.ValidateHash(user.Password, req.OldPass); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Failed to compare the password", err)
		return
	}
	if (!upperCase(req.NewPass)) || (!hasNum(req.NewPass)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password harus mengandung 1 huruf kapital dan 1 angka",
		})
		return
	}
	if (req.NewPass != req.Confirm) || (len(req.NewPass) < 8) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Password tidak sama / kurang dari 8 karakter",
		})
		return
	}

	hash, err := sdk.Hashing(req.NewPass)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to Hash", err)
		return
	}
	if err := database.DB.Model(&user).Update("password", hash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}
	sdk.Success(c, http.StatusOK, "Password berhasil dirubah", user)
}

func DeleteAccount(c *gin.Context) {
	type pass struct {
		Password string `json:"password" binding:"required"`
	}
	id := c.MustGet("user")
	claims := id.(model.UserClaims)

	var user entity.User
	if err := database.DB.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "User not found", err)
		return
	}
	var req pass
	if err := c.BindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon mengisi password Anda", err)
		return
	}

	err := sdk.ValidateHash(user.Password, req.Password)
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Password Anda Salah", err)
		return
	}
	user.Email = randomString(10)
	if err := database.DB.Save(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed", err)
		return
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Gagal menghapus akun Anda", err)
		return
	}
	sdk.Success(c, http.StatusOK, "Akun Anda telah terhapus", user)

}

func ForgotPassword(c *gin.Context) {
	type email struct {
		Email string `json:"email"`
	}
	var req email
	if err := c.BindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "error to read", err)
		return
	}

	var user entity.User
	if err := database.DB.Where("email = ?", req.Email).Find(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
		return
	}
	if user.ID == 0 {
		sdk.Fail(c, http.StatusNotFound, "User not found")
		return
	}
	newPass := randomString(8)
	m := gomail.NewMessage()
	m.SetHeader("From", "grahagrent@gmail.com")
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "Reset Password")
	m.SetBody("text/plain", fmt.Sprintf("Password baru Anda: %s", newPass))

	d := gomail.NewDialer("smtp.gmail.com", 587, "grahagrent@gmail.com", os.Getenv("PASS"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	hash, err := sdk.Hashing(newPass)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to Hash", err)
		return
	}
	if err := database.DB.Model(&user).Where("email = ?", req.Email).Update("password", hash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}
	sdk.Success(c, http.StatusOK, "Password telah direset", err)
}

func randomString(length int) string {
	var random = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = random[rand.Intn(len(random))]
	}

	return string(b)
}

func upperCase(pass string) bool {
	for _, char := range pass {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasNum(pass string) bool {
	for _, char := range pass {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
