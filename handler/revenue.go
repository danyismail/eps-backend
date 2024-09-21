package handler

import (
	"net/http"

	"eps-backend/structs"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRevenueHour(c echo.Context) error {
	h.e.Logger.Info("::GetRevenueHour Started::")
	dbConn := c.Param("e")
	result, err := h.revenueStore.GetRevenueByHour(dbConn)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       result,
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}
