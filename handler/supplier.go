package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetSuppliersEps(c echo.Context) error {
	c.Logger().Info("::GetSuppliersEps Started::")
	result, err := h.supplierStore.GetSuppliersEps()
	if err != nil {
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

func (h *Handler) GetSuppliersEActive(c echo.Context) error {
	c.Logger().Info("::GetSuppliersEActive Started::")
	result, err := h.supplierStore.GetSuppliersEActive()
	if err != nil {
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

func (h *Handler) GetSupplierByIdEps(c echo.Context) error {
	c.Logger().Info("::GetSupplierByIdEps Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	result, err := h.supplierStore.GetSupplierByIdEps(id)
	if err != nil {
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

func (h *Handler) CreateSupplierEps(c echo.Context) error {
	c.Logger().Info("::CreateSupplierEps Started::")
	supplier := new(model.Supplier)
	err := c.Bind(supplier)
	if err != nil {
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusBadRequest, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	err = h.supplierStore.CreateSupplierEps(*supplier)
	if err != nil {
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

func (h *Handler) UpdateSupplierEps(c echo.Context) error {
	c.Logger().Info("::UpdateSupplierEps Started::")
	supplier := new(model.Supplier)
	err := c.Bind(supplier)
	if err != nil {
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusBadRequest, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	err = h.supplierStore.UpdateSuppliersEps(*supplier)
	if err != nil {
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

func (h *Handler) DeleteSupplierEps(c echo.Context) error {
	c.Logger().Info("::GetSupplierByIdEps Started::")
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)
	err := h.supplierStore.DeleteSupplierEps(id)
	if err != nil {
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
