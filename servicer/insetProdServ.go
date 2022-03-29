package servicer

import (
	"apiEchor/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InsertProdServ(e echo.Context) error {

	prodServ := models.ProdServ{}
	err := e.Bind(&prodServ)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Erro inesperado")
	}

	ProdServs = append(ProdServs, prodServ)

	return e.JSON(http.StatusCreated, "msg :Produto inserido com sucesso")
}
