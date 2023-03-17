package handler

import (
	"InternBCC/database"
	"InternBCC/repository"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type gedungHandler struct {
	repo repository.GedungRepository
}

func NewGedungHandler() *gedungHandler {
	return &gedungHandler{repo: repository.NewGedungRepository(database.DB)}
}

func (h *gedungHandler) FindAllGedung(c *gin.Context) {
	get, err := h.repo.FindAll()
	if err != nil {
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
}

func (h *gedungHandler) GetGedungByID(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	get, err := h.repo.FindByID(uint(id))
	if err != nil {
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
		Links     []string
		Aturan    []string
	}
	var links []string
	for _, link := range get.Links {
		links = append(links, link.Link)
	}
	var result = fasil{
		Nama:      get.Nama,
		Alamat:    get.Alamat,
		Kecamatan: get.Kecamatan,
		Harga:     get.Harga,
		Kapasitas: get.Kapasitas,
		Luas:      get.Luas,
		Tag:       splitTag,
		Links:     links,
		Fasilitas: splitFasil,
		Aturan:    splitAturan,
	}

	sdk.Success(c, http.StatusOK, "Data Found", result)
}
