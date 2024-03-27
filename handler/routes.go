package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func (h *Handler) Register(v1 *echo.Group) {

	v1.GET("/ping", h.HealthCheck)

	kpi := v1.Group("/kpi")
	kpi.POST("/:e/list", h.GetAll)

	supplier := v1.Group("/supplier")
	supplier.GET("/:e", h.GetSuppliers)
	supplier.GET("/:e/active", h.GetActiveSuppliers)
	supplier.GET("/:e/balance", h.GetSupplierBalance)
	supplier.GET("/:e/:id", h.GetSupplierById)
	supplier.POST("/:e/create", h.CreateSupplier)
	supplier.POST("/:e/update", h.UpdateSupplier)
	supplier.DELETE("/:e/delete/:id", h.DeleteSupplier)

	sales := v1.Group("/sales")
	sales.GET("/:e", h.GetSales)
	sales.GET("/:e/periode", h.GetSalesPeriode)

	deposit := v1.Group("/deposit")
	deposit.POST("/:e", h.CreateDeposit)
	deposit.GET("/:e/:id", h.GetDeposit)
	deposit.GET("/:e/created", h.GetDepositCreated)
	deposit.GET("/:e/uploaded", h.GetDepositUploaded)
	deposit.GET("/:e/done", h.GetDepositDone)
	deposit.GET("/:e/all", h.GetAllDeposit)
	deposit.GET("/:e/image/:id", h.GetImage)
	deposit.POST("/:e/update/:id", h.UpdateDeposit)
	deposit.DELETE("/:e/delete/:id", h.CancelDeposit)
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

		h.e.Logger.Error(report)

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
