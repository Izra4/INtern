package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/model"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Booking(c *gin.Context) {
	id, _ := c.Get("user")
	claims := id.(model.UserClaims)
	GeIDStr := c.Param("id")
	GeID, _ := strconv.Atoi(GeIDStr)
	var req model.Booking
	if err := c.Bind(&req); err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Lengkapi isian Anda", err)
		return
	}
	var get = entity.Booking{
		Nama:      req.Nama,
		Tanggal:   req.Tanggal,
		Keperluan: req.Keperluan,
		Nomer:     req.Nomor,
		Alamat:    req.Alamat,
		UserID:    claims.ID,
		GedungID:  uint(GeID),
	}
	if err := database.DB.Create(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Error to create", err)
	}

}
