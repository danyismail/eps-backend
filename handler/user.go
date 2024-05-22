package handler

import (
	"eps-backend/structs"
	"eps-backend/utils"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Create User godoc
// @Summary Create New User
// @Description Create New User
// @Tags User
// @Accept json
// @Produce json
// @Param data body structs.CreateUser true "request body"
// @Success 200 {object} structs.CommonResponse
// @Router /api/user/create [post]
func (h *Handler) CreateUser(c echo.Context) error {
	h.e.Logger.Info("::CreateUser Started::")
	newUser := structs.CreateUser{}
	if err := newUser.Binding(c); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	user, err := h.userStore.Create(newUser)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Data:       user,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, structs.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}

// GetAllPaginated godoc
// @Summary Get All User
// @Description Get All User With Pagination
// @Tags User
// @Accept json
// @Produce json
// @Param data body structs.RequestPaginate true "request body"
// @Success 200 {object} structs.CommonResponse
// @Router /api/user/list [get]
func (h *Handler) GetAllPaginated(c echo.Context) error {
	h.e.Logger.Info("::GetAllPaginated Started::")
	req := structs.RequestPaginate{}
	if err := req.Binding(c); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
	}
	total := h.userStore.Count()
	users, err := h.userStore.GetAll(req.Page, req.View)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       users,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Total:       total,
		ResultCount: int64(req.View),
		Data:        utils.MappingUserResponse(users),
		StatusCode:  http.StatusOK,
		Message:     "success",
	})
}

// GetByID godoc
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id param string true "query param"
// @Success 200 {object} structs.CommonResponse
// @Router /api/user/list [get]
func (h *Handler) GetByID(c echo.Context) error {
	h.e.Logger.Info("::GetByID Started::")
	id := c.QueryParam("id")
	user, err := h.userStore.GetUser(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data: structs.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// Delete godoc
// @Summary Delete User By ID
// @Description Delete User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id param string true "query param"
// @Success 200 {object} structs.CommonResponse
// @Router /api/user/list [delete]
func (h *Handler) Delete(c echo.Context) error {
	h.e.Logger.Info("::Delete Started::")
	id := c.QueryParam("id")
	err := h.userStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.SimpleCommonResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       fmt.Sprintf("%s successfully deleted", id),
	})
}

// Login User godoc
// @Summary User Login
// @Description User Login
// @Tags User
// @Accept json
// @Produce json
// @Param data body structs.LoginRequest true "request body"
// @Success 200 {object} structs.CommonResponse
// @Router /api/user/create [post]
func (h *Handler) Login(c echo.Context) error {
	h.e.Logger.Info("::Login Started::")
	req := structs.Login{}
	if err := req.Binding(c); err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Data:       "error request validate",
		})
	}

	user, err := h.userStore.GetUser(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	if user == nil {
		userNotFound := errors.New("user not found")
		h.errorBot.SendMessage(userNotFound)
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    userNotFound.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			Message:    "email or password is wrong",
			StatusCode: http.StatusForbidden,
		})
	}

	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), *user, 15)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.SimpleCommonResponse{
			StatusCode: http.StatusForbidden,
			Data:       "failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, structs.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data: structs.LoginResponse{
			Username: user.Username,
			Role:     user.Role,
			Token:    token,
		},
	})
}
