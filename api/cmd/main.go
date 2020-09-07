package main

import (
	"github.com/family-mmyr/testWebAPIServer/api/app/server/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/students", handler.Students())

	e.Start(":8080")
}
