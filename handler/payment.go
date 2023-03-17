package handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomId() string {
	rand.Seed(time.Now().UnixNano())

	var letters strings.Builder
	for i := 0; i < 2; i++ {
		letters.WriteByte(byte(rand.Intn(26) + 65))
	}
	number := rand.Intn(100000)
	numberString := strconv.Itoa(number)
	for len(numberString) < 5 {
		numberString = "0" + numberString
	}

	return letters.String() + "-" + numberString
}

func Payment(c *gin.Context) {
	userId := c.MustGet("user")
	claims := userId.(entity.UserClaims)
	GedungIdStr := c.Param("id")
	GedungId, err := strconv.Atoi(GedungIdStr)
	if err != nil {
		sdk.Fail(c, http.StatusInternalServerError, "Failed to parse")
	}
	var booking entity.Booking
	if err := database.DB.Where("user_id", claims.ID).Last(&booking).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Booking data not found", err)
		return
	}

	supClient := supabasestorageuploader.NewSupabaseClient(
		os.Getenv("PROJECT_URL"),
		os.Getenv("PROJECT_API"),
		os.Getenv("PROJECT_STORAGE"),
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

	nominal := c.PostForm("nominal")
	harga, err := strconv.Atoi(nominal)
	if err != nil {
		sdk.Fail(c, http.StatusInternalServerError, "Failed to parse")
	}
	status := c.PostForm("status")
	var req = entity.Payment{
		ID:        randomId(),
		UserID:    claims.ID,
		GedungID:  uint(GedungId),
		Link:      link,
		Nominal:   harga,
		Status:    status,
		BookingID: booking.ID,
	}
	if err := database.DB.Create(&req).Error; err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Mohon upload file bukti pembayaran", err)
		return
	}
	sdk.Success(c, http.StatusOK, "Pembayaran Berhasil", req)
}

func GetHistory(c *gin.Context) {
	userId := c.MustGet("user")
	claims := userId.(entity.UserClaims)

	var get []entity.Payment
	if err := database.DB.Where("user_id = ?", claims.ID).Find(&get).
		Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
	}

	type body struct {
		ID             string `json:"id"`
		Tanggal        string `json:"tanggal"`
		UserID         uint   `json:"user_id"`
		GedungID       uint   `json:"gedung_id"`
		NamaGedung     string `json:"nama_gedung"`
		LinkGedung     string `json:"link_gedung"`
		Link           string `json:"link"`
		Nominal        int    `json:"nominal"`
		Status         string `json:"status"`
		BookingID      uint   `json:"booking_id"`
		NamaUser       string `json:"nama_user"`
		TanggalBooking string `json:"tanggal_booking"`
		Fasilitas      string `json:"fasilitas"`
	}

	var result []body
	for _, value := range get {
		var temp body
		waktu := value.CreatedAt
		format := waktu.Format("2006-01-02")
		temp.ID = value.ID
		temp.Tanggal = format
		temp.UserID = value.UserID
		temp.GedungID = value.GedungID
		var gedung entity.Gedung
		if err := database.DB.Debug().Preload("Links").Where("id = ?", value.GedungID).Find(&gedung).Error; err != nil {
			sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
			return
		}

		temp.NamaGedung = gedung.Nama
		temp.LinkGedung = gedung.Links[0].Link
		temp.Link = value.Link
		temp.Nominal = value.Nominal
		temp.Status = value.Status
		temp.BookingID = value.BookingID
		var booking entity.Booking
		if err := database.DB.Where("id = ?", value.BookingID).Find(&booking).Error; err != nil {
			sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
			return
		}
		waktuBooking := booking.Tanggal
		wbFormat := waktuBooking.Format("2006-01-02")
		temp.NamaUser = booking.Nama
		temp.TanggalBooking = wbFormat
		temp.Fasilitas = booking.Fasilitas
		result = append(result, temp)
	}
	if len(result) == 0 {
		sdk.Fail(c, http.StatusNotFound, "Anda belum melakukan pembaaran")
		return
	}
	sdk.Success(c, http.StatusOK, "Data found", result)
}
