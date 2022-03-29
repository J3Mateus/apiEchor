package main

import (
	"apiEchor/servicer"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Print("Api rodando")
	e := echo.New()
	g := e.Group("/api")
	g.GET("/prodserv", servicer.GetallProdServ)
	g.GET("/prodserve/:id", servicer.GetProdServ)
	g.POST("/insetprodserv", servicer.InsertProdServ)
	e.Start(":8000")
}
