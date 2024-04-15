package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetLabaReseller(c echo.Context) error {
	h.e.Logger.Info("::GetLabaReseller Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	resellerID := c.QueryParam("id")
	result, err := h.resellerStore.GetLaba(c.Param("e"), startDt, endDt, resellerID)
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

func (h *Handler) GetSummaryReseller(c echo.Context) error {
	h.e.Logger.Info("::GetSummaryReseller Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	resellerID := c.QueryParam("id")
	result, err := h.resellerStore.GetSum(c.Param("e"), startDt, endDt, resellerID)
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
