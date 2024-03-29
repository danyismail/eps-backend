package handler

import (
	"eps-backend/structs"
	"eps-backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetSales(c echo.Context) error {
	h.e.Logger.Info("::GetSales Started::")
	result, err := h.salesStore.GetSales(c.Param("e"))
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

func (h *Handler) GetSalesPeriode(c echo.Context) error {
	h.e.Logger.Info("::GetSalesPeriode Started::")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	if startDate == "" || endDate == "" {
		startDate = utils.StartDate
		endDate = utils.EndDate
	}
	result, err := h.salesStore.GetSalesPeriode(c.Param("e"), startDate, endDate)
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
