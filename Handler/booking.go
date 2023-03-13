package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Booking(c *gin.Context) {
	id, _ := c.Get("user")
	//status unautho
	//if err != true {
	//	sdk.FailOrError(c, http.StatusUnauthorized, "Unauthorized", err)
	//}
	claims := id.(model.UserClaims)
	GeIDStr := c.Param("id")
	GeID, _ := strconv.Atoi(GeIDStr)
	//statusbad
	var req model.Booking
	if err := c.BindJSON(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Lengkapi isian Anda", err)
		return
	}
	tglStr := req.Tanggal
	template := "2006-01-02"
	parsed, err := time.Parse(template, tglStr)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to parse", err)
	}
	var get = entity.Booking{
		Nama:      req.Nama,
		Tanggal:   parsed,
		Keperluan: req.Keperluan,
		Nomer:     req.Nomor,
		Alamat:    req.Alamat,
		Fasilitas: req.Fasilitas,
		UserID:    claims.ID,
		GedungID:  uint(GeID),
	}
	if err := database.DB.Create(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Error to create", err)
		return
	}
	sdk.Success(c, 200, "Booking telah dibuat", get)

}
