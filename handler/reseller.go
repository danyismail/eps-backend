package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetLabaReseller(c echo.Context) error {
	h.e.Logger.Info("::GetLabaReseller Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	resellerID := c.QueryParam("id")
	result, err := h.resellerStore.GetLaba(c.Param("e"), startDt, endDt, resellerID)
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

func (h *Handler) GetSummaryReseller(c echo.Context) error {
	h.e.Logger.Info("::GetSummaryReseller Started::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	resellerID := c.QueryParam("id")
	result, err := h.resellerStore.GetSum(c.Param("e"), startDt, endDt, resellerID)
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

func (h *Handler) ListSupplier(c echo.Context) error {
	h.e.Logger.Info("::ListSupplier Started::")
	param := model.ResellerParam{}
	err := c.Bind(&param)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	result, err := h.resellerStore.GetList(c.Param("e"), param.Search)
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

func (h *Handler) GetLabaHourly(c echo.Context) error {
	h.e.Logger.Info("::GetLabaHourly Started::")
	result, err := h.resellerStore.GetLabaHourly(c.Param("e"))
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

func (h *Handler) GetLabaRugi(c echo.Context) error {
	h.e.Logger.Info("::GetLabaRugi Started::")
	from := c.QueryParam("startDt")
	to := c.QueryParam("endDt")
	result, err := h.resellerStore.GetLabaRugi(c.Param("e"), from, to)
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
