package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

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

//func ChangeNameNumber(c *gin.Context) {
//	type changes struct {
//		Email string `json:"email"`
//		Nama  string `json:"nama"`
//		Nomor string `json:"nomor"`
//	}
//	userId := c.MustGet("user")
//	claims := userId.(model.UserClaims)
//
//	var req changes
//	if err := c.BindJSON(&req); err != nil {
//		sdk.FailOrError(c, http.StatusInternalServerError, "Error to Read", err)
//		return
//	}
//	if err := database.DB.Model(&users).Where("id = ?", claims.ID).Updates(req).Error; err != nil {
//		sdk.FailOrError(c, http.StatusInternalServerError, "error", err)
//		return
//	}
//
//	sdk.Success(c, http.StatusOK, "Profil telah diperbarui", users)
//}

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

	var req pass
	if err := c.BindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon mengisi password Anda", err)
		return
	}

	var user entity.User
	if err := database.DB.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "User not found", err)
		return
	}

	err := sdk.ValidateHash(user.Password, req.Password)
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Password Anda Salah", err)
		return
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Gagal menghapus akun Anda", err)
		return
	}
	sdk.Success(c, http.StatusOK, "Akun Anda telah terhapus", user)
}
