package structs

import "github.com/labstack/echo"

type PageView struct {
	Page int    `json:"page"`
	View int    `json:"view"`
	Mdn  string `json:"mdn"`
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
