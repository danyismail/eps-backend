package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	h.e.Logger.Info("::HealthCheck Started::")
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       "pong",
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}
