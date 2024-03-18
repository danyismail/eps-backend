package handler

import (
	"eps-backend/db"
	"eps-backend/store"
)

type Handler struct {
	kpiStore         store.KpiStore
	depositStore     store.DepositStore
	salesStore       store.SalesStore
	supplierStore    store.SupplierStore
	depositNoteStore store.DepositNote
}

func NewHandler(db db.DBConnection) *Handler {
	return &Handler{
		kpiStore:         store.NewKpiStore(db),
		depositStore:     store.NewDepositStore(db),
		salesStore:       store.NewSalesStore(db),
		supplierStore:    store.NewSupplierStore(db),
		depositNoteStore: store.NewDepositNoteStore(db),
	}
}
