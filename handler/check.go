package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       "PONG",
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}
