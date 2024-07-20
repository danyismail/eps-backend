package handler

import (
	"eps-backend/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func (h *Handler) Register(v1 *echo.Group) {
	customJWTMiddleware := utils.CreateEchoMiddleware(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	v1.GET("/ping", h.HealthCheck)

	auth := v1.Group("/auth")
	auth.POST("/login", h.Login)

	user := v1.Group("/user", customJWTMiddleware)
	user.POST("/create", h.CreateUser)
	user.GET("/list", h.GetAllPaginated)
	user.GET("/search", h.GetByID)
	user.DELETE("/delete", h.Delete)

	kpi := v1.Group("/kpi", customJWTMiddleware)
	kpi.POST("/:e/list", h.GetAll)

	supplier := v1.Group("/supplier", customJWTMiddleware)
	supplier.GET("/:e", h.GetSuppliers)
	supplier.GET("/:e/active", h.GetActiveSuppliers)
	supplier.GET("/:e/balance", h.GetSupplierBalance)
	supplier.GET("/:e/:id", h.GetSupplierById)
	supplier.POST("/:e/create", h.CreateSupplier)
	supplier.POST("/:e/update", h.UpdateSupplier)
	supplier.DELETE("/:e/delete/:id", h.DeleteSupplier)

	sales := v1.Group("/sales", customJWTMiddleware)
	sales.GET("/:e", h.GetSales)
	sales.GET("/:e/pph", h.GetPPH)
	sales.GET("/:e/periode", h.GetSalesPeriode)

	deposit := v1.Group("/deposit", customJWTMiddleware)
	deposit.POST("/:e", h.CreateDeposit)
	deposit.GET("/:e/:id", h.GetDeposit)
	deposit.GET("/:e/created", h.GetDepositCreated)
	deposit.GET("/:e/uploaded", h.GetDepositUploaded)
	deposit.GET("/:e/done", h.GetDepositDone)
	deposit.GET("/:e/all", h.GetAllDeposit)
	deposit.GET("/:e/image/:id", h.GetImage)
	deposit.POST("/:e/update/:id", h.UpdateDeposit)
	deposit.GET("/:e/delete/:id", h.CancelDeposit)

	sn := v1.Group("/sn", customJWTMiddleware)
	sn.GET("/null/:target", h.CheckSN)
	sn.GET("/duplicate/:target", h.DuplicateSN)

	reseller := v1.Group("/reseller", customJWTMiddleware)
	reseller.GET("/:e/laba", h.GetLabaReseller)
	reseller.GET("/:e/sum", h.GetSummaryReseller)
	reseller.GET("/:e/list", h.ListSupplier)
	reseller.GET("/:e/laba/hourly", h.GetLabaHourly)
	reseller.GET("/:e/labarugi", h.GetLabaRugi)

	margin := v1.Group("/margin", customJWTMiddleware)
	margin.GET("/:e/reseller", h.GetMarginReseller)
	margin.GET("/:e/supplier", h.GetMarginSupplier)
	margin.GET("/:e/provider", h.GetMarginProvider)

	hub := v1.Group("/hub", customJWTMiddleware)
	hub.GET("/brand-revenue", h.GetBrandRevenue)
	hub.GET("/:cnx/brand-category-revenue", h.GetBrandCategoryRevenue)

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
