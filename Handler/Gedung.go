package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func FindAllGedung(c *gin.Context) {
	var get []entity.Gedung
	type disp struct {
		ID        uint
		Nama      string
		Kecamatan string
		Harga     string
		Alamat    string
	}

	var result []disp
	if err := database.DB.Model(&get).Find(&result).Error; err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to load data", err)
		return
	}
	//var links []entity.Link
	//if err := database.DB.Find(&links); err != nil {
	//	for _, link := range links {
	//		fmt.Println(link.ID)
	//	}
	//}
	sdk.Success(c, 200, "Data found", result)
}

func GetGedungByID(c *gin.Context) {
	id := c.Param("id")
	var get entity.Gedung
	if err := database.DB.First(&get, id).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not Found", err)
		return
	}
	splitFasil := strings.Split(get.Fasilitas, ";")
	splitTag := strings.Split(get.Tag, ";")
	type fasil struct {
		Nama      string
		Alamat    string
		Kecamatan string
		Harga     int
		Kapasitas string
		Luas      string
		Fasilitas []string
		Tag       []string
		Aturan    string
	}

	var result = fasil{
		Nama:      get.Nama,
		Alamat:    get.Alamat,
		Kecamatan: get.Kecamatan,
		Harga:     get.Harga,
		Kapasitas: get.Kapasitas,
		Luas:      get.Luas,
		Tag:       splitTag,
		Fasilitas: splitFasil,
		Aturan:    get.Aturan,
	}

	sdk.Success(c, http.StatusOK, "Data Found", result)
}
