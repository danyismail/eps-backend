package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetKpi(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	result, err := h.kpiStore.FindAll(1, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)

}

func (h *Handler) KpiTest(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	result, err := h.kpiStore.Test()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)

}
