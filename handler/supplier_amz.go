package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetSuppliersAmz(c echo.Context) error {
	h.e.Logger.Info("::GetSuppliersAmz Started::")
	result, err := h.supplierStore.GetSuppliersAmz()
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

func (h *Handler) GetSuppliersAActive(c echo.Context) error {
	h.e.Logger.Info("::GetSuppliersEActive Started::")
	result, err := h.supplierStore.GetSuppliersAActive()
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

func (h *Handler) GetSupplierByIdAmz(c echo.Context) error {
	h.e.Logger.Info("::GetSupplierByIdAmz Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	result, err := h.supplierStore.GetSupplierByIdAmz(id)
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

func (h *Handler) CreateSupplierAmz(c echo.Context) error {
	h.e.Logger.Info("::CreateSupplierEps Started::")
	supplier := new(model.Supplier)
	err := c.Bind(supplier)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusBadRequest, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	err = h.supplierStore.CreateSupplierAmz(*supplier)
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
		Data:       supplier,
		StatusCode: http.StatusOK,
		Message:    "success",
	})

}

func (h *Handler) UpdateSupplierAmz(c echo.Context) error {
	h.e.Logger.Info("::UpdateSupplierEps Started::")
	supplier := new(model.Supplier)
	err := c.Bind(supplier)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusBadRequest, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	err = h.supplierStore.UpdateSuppliersAmz(*supplier)
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
		Data:       supplier,
		StatusCode: http.StatusOK,
		Message:    "success",
	})

}

func (h *Handler) DeleteSupplierAmz(c echo.Context) error {
	h.e.Logger.Info("::GetSupplierByIdEps Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	err := h.supplierStore.DeleteSupplierAmz(id)
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
		Data:       strId,
		StatusCode: http.StatusOK,
		Message:    "success",
	})

}
