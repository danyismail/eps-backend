package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBrandRevenue(c echo.Context) error {
	h.e.Logger.Info("::GetBrandRevenue Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	result, err := h.hubStore.GetBrandRevenue(startDt, endDt)
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
