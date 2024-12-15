package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetValueList(c echo.Context) error {
	h.e.Logger.Info("::Get ValueList Started::")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	view, _ := strconv.Atoi(c.QueryParam("view"))
	shortCode := c.QueryParam("short_code")
	vl, total, err := h.valueListStore.GetAll(page, view, shortCode)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       vl,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		ResultCount: int64(len(vl)),
		Total:       int64(total),
		Data:        vl,
		StatusCode:  http.StatusOK,
		Message:     "success",
	})
}

func (h *Handler) CreateValueList(c echo.Context) error {
	h.e.Logger.Info("::Create ValueList Started::")
	vl := model.ValueList{}
	if err := c.Bind(&vl); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}
	err := h.valueListStore.Create(&vl)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
	})
}

func (h *Handler) UpdateValueList(c echo.Context) error {
	h.e.Logger.Info("::Update ValueList Started::")
	vl := model.ValueList{}
	if err := c.Bind(&vl); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.valueListStore.Update(id, &vl)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
	})
}

func (h *Handler) DeleteValueList(c echo.Context) error {
	h.e.Logger.Info("::Delete ValueList Started::")
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.valueListStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusOK,
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
	})
}

func (h *Handler) GetValueListById(c echo.Context) error {
	h.e.Logger.Info("::Get ValueList By Id Started::")
	id, _ := strconv.Atoi(c.Param("id"))
	vl, err := h.valueListStore.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
		Data:       vl,
	})
}
