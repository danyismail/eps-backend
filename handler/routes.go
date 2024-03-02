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

	f := v1.Group("/finance")
	f.GET("/eps/supplier", h.GetSuppliersEps)
	f.GET("/eps/supplier/:id", h.GetSupplierByIdEps)
	f.POST("/eps/supplier/create", h.CreateSupplierEps)
	f.POST("/eps/supplier/update", h.UpdateSupplierEps)
	f.DELETE("/eps/supplier/delete/:id", h.DeleteSupplierEps)

	f.GET("/amz/supplier", h.GetSuppliersAmz)
	f.GET("/amz/supplier/:id", h.GetSupplierByIdAmz)
	f.POST("/amz/supplier/create", h.CreateSupplierAmz)
	f.POST("/amz/supplier/update", h.UpdateSupplierAmz)
	f.DELETE("/amz/supplier/delete/:id", h.DeleteSupplierAmz)

	f.POST("/:e/deposit", h.CreateDeposit)
	f.GET("/:e/deposit/:id", h.GetDeposit)
	f.GET("/:e/deposit/created", h.GetDepositCreated)
	f.GET("/:e/deposit/uploaded", h.GetDepositUploaded)
	f.GET("/:e/deposit/done", h.GetDepositDone)
	f.GET("/:e/image/:id", h.GetImage)
	f.POST("/:e/deposit/:id", h.UpdateDeposit)
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
