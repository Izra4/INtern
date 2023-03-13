package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func FindAllGedung(c *gin.Context) {
	var get []entity.Gedung

	//var result []disp
	if err := database.DB.Debug().
		Preload("Links", func(db *gorm.DB) *gorm.DB {
			return db.Limit(1) // set limit to 1 for Model2
		}).
		Find(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to load data", err)
		return
	}

	//if err := database.DB.Preload("Links").Select("nama, alamat, kecamatan, harga, Links").Find(&get).Error; err != nil {
	//	sdk.FailOrError(c, http.StatusInternalServerError, "Error to get Data", err)
	//}
	sdk.Success(c, 200, "Data found", get)
}

func GetGedungByID(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	var get entity.Gedung
	if err := database.DB.Preload("Links").First(&get, id).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not Found", err)
		return
	}
	splitFasil := strings.Split(get.Fasilitas, ";")
	splitTag := strings.Split(get.Tag, ";")
	splitAturan := strings.Split(get.Aturan, ";")

	type fasil struct {
		Nama      string
		Alamat    string
		Kecamatan string
		Harga     int
		Kapasitas string
		Luas      string
		Fasilitas []string
		Tag       []string
		Links     []entity.Link
		Aturan    []string
	}

	var result = fasil{
		Nama:      get.Nama,
		Alamat:    get.Alamat,
		Kecamatan: get.Kecamatan,
		Harga:     get.Harga,
		Kapasitas: get.Kapasitas,
		Luas:      get.Luas,
		Tag:       splitTag,
		Links:     get.Links,
		Fasilitas: splitFasil,
		Aturan:    splitAturan,
	}

	sdk.Success(c, http.StatusOK, "Data Found", result)
}
