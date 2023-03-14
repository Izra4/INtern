package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomId() string {
	// Membuat seed untuk generator angka acak
	rand.Seed(time.Now().UnixNano())

	// Membuat 2 karakter acak
	var letters strings.Builder
	for i := 0; i < 2; i++ {
		letters.WriteByte(byte(rand.Intn(26) + 65))
	}

	// Membuat 5 digit angka acak
	number := rand.Intn(100000)
	numberString := strconv.Itoa(number)
	for len(numberString) < 5 {
		numberString = "0" + numberString
	}

	// Menggabungkan karakter dan angka ke format "XX-12332"
	return letters.String() + "-" + numberString
}

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
	type nominal struct {
		Nominal int `json:"nominal"`
	}
	var price nominal
	if err := c.Bind(&price); err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to read", err)
	}
	var req = entity.Payment{
		ID:       randomId(),
		UserID:   claims.ID,
		GedungID: uint(GedungId),
		Link:     link,
		Nominal:  price.Nominal,
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
