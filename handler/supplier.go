package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetSuppliers(c echo.Context) error {
	h.e.Logger.Info("::GetSuppliers Started::")
	result, err := h.supplierStore.GetSuppliers(c.Param("e"))
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

func (h *Handler) GetActiveSuppliers(c echo.Context) error {
	h.e.Logger.Info("::GetSuppliersEActive Started::")
	result, err := h.supplierStore.GetActiveSuppliers(c.Param("e"))
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

func (h *Handler) GetSupplierById(c echo.Context) error {
	h.e.Logger.Info("::GetSupplierById Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	result, err := h.supplierStore.GetSupplierById(c.Param("e"), id)
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

func (h *Handler) CreateSupplier(c echo.Context) error {
	h.e.Logger.Info("::CreateSupplier Started::")
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
	err = h.supplierStore.CreateSupplier(c.Param("e"), *supplier)
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

func (h *Handler) UpdateSupplier(c echo.Context) error {
	h.e.Logger.Info("::UpdateSupplier Started::")
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
	err = h.supplierStore.UpdateSupplier(c.Param("e"), *supplier)
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

func (h *Handler) DeleteSupplier(c echo.Context) error {
	h.e.Logger.Info("::DeleteSupplier Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	err := h.supplierStore.DeleteSupplier(c.Param("e"), id)
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
