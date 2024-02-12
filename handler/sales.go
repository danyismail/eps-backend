package handler

import (
	"eps-backend/structs"
	"eps-backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetSales(c echo.Context) error {
	c.Logger().Info("::GetSales Started::")
	result, err := h.salesStore.GetSalesToday()
	if err != nil {
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

func (h *Handler) GetSalesProd(c echo.Context) error {
	c.Logger().Info("::GetSales Started::")
	result, err := h.salesStore.GetSalesTodayProd()
	if err != nil {
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
	c.Logger().Info("::GetSalesPeriode Started::")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	if startDate == "" || endDate == "" {
		startDate = utils.StartDate
		endDate = utils.EndDate
	}
	result, err := h.salesStore.Sales(startDate, endDate)
	if err != nil {
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

func (h *Handler) GetSalesPeriodeProd(c echo.Context) error {
	c.Logger().Info("::GetSalesPeriode Started::")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	if startDate == "" || endDate == "" {
		startDate = utils.StartDate
		endDate = utils.EndDate
	}
	result, err := h.salesStore.SalesProd(startDate, endDate)
	if err != nil {
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
