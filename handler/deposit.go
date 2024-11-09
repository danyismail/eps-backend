package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetSupplierBalance(c echo.Context) error {
	h.e.Logger.Info("::GetSupplierBalance Started::")
	result, err := h.depositStore.GetBalance(c.Param("e"), c.QueryParam("p"))
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
