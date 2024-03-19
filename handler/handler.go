package handler

import (
	"eps-backend/bot"
	"eps-backend/db"
	"eps-backend/store"
)

type Handler struct {
	errorBot         bot.MyBot
	kpiStore         store.KpiStore
	depositStore     store.DepositStore
	salesStore       store.SalesStore
	supplierStore    store.SupplierStore
	depositNoteStore store.DepositNote
}

func NewHandler(db db.DBConnection) *Handler {
	return &Handler{
		errorBot:         bot.BotInit(),
		kpiStore:         store.NewKpiStore(db),
		depositStore:     store.NewDepositStore(db),
		salesStore:       store.NewSalesStore(db),
		supplierStore:    store.NewSupplierStore(db),
		depositNoteStore: store.NewDepositNoteStore(db),
	}
}
