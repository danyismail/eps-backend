package handler

import (
	"eps-backend/bot"
	"eps-backend/db"
	"eps-backend/store"

	"github.com/labstack/echo"
)

type Handler struct {
	e                *echo.Echo
	errorBot         bot.MyBot
	kpiStore         store.KpiStore
	depositStore     store.DepositStore
	salesStore       store.SalesStore
	supplierStore    store.SupplierStore
	depositNoteStore store.DepositNote
}

func NewHandler(db db.DBConnection, e *echo.Echo) *Handler {
	return &Handler{
		e:                e,
		errorBot:         bot.BotInit(),
		kpiStore:         store.NewKpiStore(db),
		depositStore:     store.NewDepositStore(db),
		salesStore:       store.NewSalesStore(db),
		supplierStore:    store.NewSupplierStore(db),
		depositNoteStore: store.NewDepositNoteStore(db),
	}
}
