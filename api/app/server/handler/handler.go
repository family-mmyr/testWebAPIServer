package handler

import (
	"net/http"

	"github.com/family-mmyr/testWebAPIServer/api/domain/model"
	"github.com/labstack/echo"
)

func Students() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := []model.Student{
			{
				Name: "taro",
				Age:  16,
				Sex:  "male",
			},
			{
				Name: "hana",
				Age:  17,
				Sex:  "female",
			},
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"students": resp,
		})
	}
}
