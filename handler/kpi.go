package handler

import (
	"eps-backend/structs"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetAll(c echo.Context) error {
	h.e.Logger.Info("::GetAll KPI Started::")
	req := structs.PageView{}
	if err := req.Binding(c); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
	}
	path := c.Param("e")
	fmt.Println("cek param ", path)
	result, attribute, err := h.kpiStore.GetAll(path, req.StartDt, req.EndDt, req.Page, req.View, req.Mdn, req.Status, req.Shift)
	if err != nil {
		h.e.Logger.Error(err)
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
