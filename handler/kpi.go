package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetKPI(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	req := structs.PageView{}
	if err := req.Binding(c); err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
	}
	result, total, err := h.kpiStore.FindAll(req.Page, req.View, req.Mdn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Total:      int(total),
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Total:      int(total),
		Data:       result,
		StatusCode: http.StatusOK,
		Message:    "success",
	})

}

func (h *Handler) MockKPI(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	result, err := h.kpiStore.Test()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Total:      10,
		Data:       result,
		StatusCode: http.StatusOK,
		Message:    "success",
	})

}