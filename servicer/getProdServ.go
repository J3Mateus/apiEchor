package servicer

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetProdServ(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	serv := ProdServs[id]
	return e.JSON(http.StatusOK, serv)
}
