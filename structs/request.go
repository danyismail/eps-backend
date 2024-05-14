package structs

import "github.com/labstack/echo/v4"

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
