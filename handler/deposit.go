package handler

import (
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

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