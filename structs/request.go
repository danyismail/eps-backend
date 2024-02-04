package structs

import "github.com/labstack/echo"

type PageView struct {
	StartDt string `json:"startDt"`
	EndDt   string `json:"endDt"`
	Page    int    `json:"page"`
	View    int    `json:"view"`
	Mdn     string `json:"mdn"`
	Status  int    `json:"status"`
	Shift   string `json:"shift"`
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
