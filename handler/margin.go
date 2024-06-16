package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetMarginReseller(c echo.Context) error {
	h.e.Logger.Info("::GetMarginReseller Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	result, err := h.marginStore.ByReseller(c.Param("e"), startDt, endDt)
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

func (h *Handler) GetMarginSupplier(c echo.Context) error {
	h.e.Logger.Info("::GetMarginSupplier Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	result, err := h.marginStore.BySupplier(c.Param("e"), startDt, endDt)
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

func (h *Handler) GetMarginProvider(c echo.Context) error {
	h.e.Logger.Info("::GetMarginProvider Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	result, err := h.marginStore.ByProvider(c.Param("e"), startDt, endDt)
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
