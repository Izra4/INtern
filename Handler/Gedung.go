package Handler

import (
	"InternBCC/database"
	"InternBCC/entity"
	"InternBCC/sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllGedung(c *gin.Context) {
	var get []entity.Gedung
	if err := database.DB.Find(&get).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "failed to get data", err)
		return
	}
	sdk.Success(c, 200, "Data found", get)
}

func GetGedungByID(c *gin.Context) {
	id := c.Param("id")
	var get entity.Gedung
	if err := database.DB.First(&get, id).Error; err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not Found", err)
		return
	}
	sdk.Success(c, http.StatusOK, "Data Found", get)
}
