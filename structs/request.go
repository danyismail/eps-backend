package structs

import "github.com/labstack/echo/v4"

type RequestPaginate struct {
	Page int `json:"page" example:"1"`
	View int `json:"view" example:"10"`
}

func (r *RequestPaginate) Binding(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type PageView struct {
	StartDt string `json:"startDt" example:"2024-05-11"`
	EndDt   string `json:"endDt" example:"2024-05-11"`
	Page    int    `json:"page" example:"1"`
	View    int    `json:"view" example:"10"`
	Mdn     string `json:"mdn" example:"081284088408"`
	Status  int    `json:"status" example:"20"`
	Shift   string `json:"shift" example:"sore"`
}

func (r *PageView) Binding(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type CreateUser struct {
	Username string `json:"username" required:"max:6" example:"John Doe"`
	Email    string `json:"email" example:"example@email.com"`
	Password string `json:"password" example:"password123"`
	Role     string `json:"role" example:"superadmin"`
}

func (u *CreateUser) Binding(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	return nil
}

type Login struct {
	Email    string `json:"email" example:"example@email.com"`
	Password string `json:"password" example:"password123"`
}

func (u *Login) Binding(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	return nil
}
