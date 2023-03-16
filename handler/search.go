package handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	kecamatan := c.Query("kecamatan")
	var gedungs []entity.Gedung
	if name != "" {
		if err := database.DB.Where("nama LIKE ?", "%"+name+"%").Preload("Links").Find(&gedungs).Error; err != nil {
			sdk.FailOrError(c, http.StatusNotFound, "Informasi gedung tidak ada", err)
			return
		}
		if len(gedungs) == 0 {
			sdk.Fail(c, http.StatusNotFound, "Data gedung tidak ditemukan")
			return
		}
	} else if kecamatan != "" {
		if err := database.DB.Where("kecamatan LIKE ?", "%"+kecamatan+"%").Preload("Links").Find(&gedungs).Error; err != nil {
			sdk.FailOrError(c, http.StatusNotFound, "Informasi gedung tidak ada", err)
			return
		}
		if len(gedungs) == 0 {
			sdk.Fail(c, http.StatusNotFound, "Data gedung tidak ditemukan")
			return
		}
	} else {
		if err := database.DB.Preload("Links").Find(&gedungs).Error; err != nil {
			sdk.FailOrError(c, http.StatusNotFound, "Informasi gedung tidak ada", err)
			return
		}
	}
	type req struct {
		Id        uint   `json:"id"`
		Nama      string `json:"nama"`
		Alamat    string `json:"alamat"`
		Kecamatan string `json:"kecamatan"`
		Harga     int    `json:"harga"`
		Link      string `json:"link"`
	}

	var result []req
	for _, value := range gedungs {
		var temp req
		temp.Id = value.ID
		temp.Nama = value.Nama
		temp.Alamat = value.Alamat
		temp.Kecamatan = value.Kecamatan
		temp.Harga = value.Harga
		temp.Link = value.Links[0].Link
		result = append(result, temp)
	}
	sdk.Success(c, 200, "Data found", result)
}
