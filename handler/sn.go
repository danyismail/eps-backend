package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) CheckSN(c echo.Context) error {
	h.e.Logger.Info("::CheckSN Started::")
	path := c.Param("target")
	result, err := h.snStore.GetNullableSN(path)
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

func (h *Handler) DuplicateSN(c echo.Context) error {
	h.e.Logger.Info("::DuplicateSN Started::")
	path := c.Param("target")
	result, err := h.snStore.GetDuplicateSN(path)
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
