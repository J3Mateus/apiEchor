package servicer

import (
	"apiEchor/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var ProdServs []models.ProdServ

var prodServ models.ProdServ

func GetallProdServ(e echo.Context) error {

	return e.JSON(http.StatusOK, ProdServs)
}
