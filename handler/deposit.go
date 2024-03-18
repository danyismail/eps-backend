package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetSupplierBalance(c echo.Context) error {
	c.Logger().Info("::GetSupplierBalance Started::")
	result, err := h.depositStore.GetBalance(c.Param("e"))
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

func (h *Handler) GetBalance(c echo.Context) error {
	c.Logger().Info("::GetBalance Started::")
	result, err := h.depositStore.GetBalanceToday()
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

func (h *Handler) GetBalanceProd(c echo.Context) error {
	c.Logger().Info("::GetBalance Started::")
	result, err := h.depositStore.GetBalanceTodayProd()
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
