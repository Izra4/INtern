package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func FindAllGedung(c *gin.Context) {
	var get []entity.Gedung

	//var result []disp
	if err := database.DB.Debug().Preload("Links").Find(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to load data", err)
		return
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
	for _, value := range get {
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
	//if err := database.DB.Preload("Links").Select("nama, alamat, kecamatan, harga, Links").Find(&get).Error; err != nil {
	//	sdk.FailOrError(c, http.StatusInternalServerError, "Error to get Data", err)
	//}
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
