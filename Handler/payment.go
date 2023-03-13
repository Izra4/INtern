package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Payment(c *gin.Context) {
	userId := c.MustGet("user")
	claims := userId.(model.UserClaims)
	GedungIdStr := c.Param("id")
	GedungId, _ := strconv.Atoi(GedungIdStr)
	supClient := supabasestorageuploader.NewSupabaseClient(
		"https://ontvftbxgsmzxwlqhsdn.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im9udHZmdGJ4Z3Ntenh3bHFoc2RuIiwicm9sZSI6ImFub24iLCJpYXQiOjE2Nzg0MDgxMzQsImV4cCI6MTk5Mzk4NDEzNH0.7yypIF1_gkHACVRxolU2KjhLpdUumKw3OdaRtHSnB9Q",
		"gambar-gedung",
		"",
	)
	file, err := c.FormFile("bukti")
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Error to upload file", err)
		return
	}
	link, err := supClient.Upload(file)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to upload file", err)
		return
	}

	var req = entity.Payment{
		Model:    gorm.Model{},
		UserID:   claims.ID,
		GedungID: uint(GedungId),
		Link:     link,
	}
	if err := database.DB.Create(&req).Error; err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon upload file bukti pembayaran", err)
	}
	sdk.Success(c, http.StatusOK, "Pembayaran Berhasil", req)
}

func GetHistory(c *gin.Context) {
	userId := c.MustGet("user")
	claims := userId.(model.UserClaims)

	var get []entity.Payment
	if err := database.DB.Where("user_id = ?", claims.ID).Find(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
	}
	sdk.Success(c, http.StatusOK, "Data found", get)
}
