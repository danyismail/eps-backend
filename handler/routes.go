package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func (h *Handler) Register(v1 *echo.Group) {
	sf := v1.Group("/eps")
	sf.GET("/ping", h.HealthCheck)
	sf.GET("/mockKpis", h.MockKPI)
	sf.POST("/getKpis", h.GetKPI)
	sf.POST("/getKpisProd", h.GetKPIProd)
	sf.GET("/deposit", h.GetBalance)
	sf.GET("/depositProd", h.GetBalanceProd)
	sf.GET("/sales", h.GetSales)
	sf.GET("/salesProd", h.GetSalesProd)
	sf.GET("/salesPeriode", h.GetSalesPeriode)
	sf.GET("/salesPeriodeProd", h.GetSalesPeriodeProd)
}

func (h *Handler) HttpErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var code = report.Code
		if report.Code > 88000 {
			code = http.StatusInternalServerError
		}
		rid := c.Response().Header().Get(echo.HeaderXRequestID)
		report.SetInternal(echo.NewHTTPError(0, "Request ID : "+rid))

		c.Logger().Error(report)

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}
			}
		}

		c.JSON(code, report)
	}
}
