package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	h.e.Logger.Info("::HealthCheck Started::")
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       "api is alive v1.0.1",
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}
