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
	result, attribute, err := h.kpiStore.FindAll(req.StartDt, req.EndDt, req.Page, req.View, req.Mdn, req.Status, req.Shift)
	if err != nil {
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Total:       attribute.Total,
			ResultCount: attribute.View,
			Success:     attribute.Success,
			Failed:      attribute.Failed,
			Data:        nil,
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Total:       attribute.Total,
		ResultCount: attribute.View,
		Success:     attribute.Success,
		Failed:      attribute.Failed,
		Data:        result,
		StatusCode:  http.StatusOK,
		Message:     "success",
	})

}

func (h *Handler) GetKPIProd(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	req := structs.PageView{}
	if err := req.Binding(c); err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
	}
	result, attribute, err := h.kpiStore.FindAllProd(req.StartDt, req.EndDt, req.Page, req.View, req.Mdn, req.Status, req.Shift)
	if err != nil {
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Total:       attribute.Total,
			ResultCount: attribute.View,
			Success:     attribute.Success,
			Failed:      attribute.Failed,
			Data:        nil,
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Total:       attribute.Total,
		ResultCount: attribute.View,
		Success:     attribute.Success,
		Failed:      attribute.Failed,
		Data:        result,
		StatusCode:  http.StatusOK,
		Message:     "success",
	})

}

func (h *Handler) MockKPI(c echo.Context) error {
	c.Logger().Info("::GetKpi Started::")
	result, err := h.kpiStore.Test()
	if err != nil {
		h.errorBot.SendMessage(err)
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
