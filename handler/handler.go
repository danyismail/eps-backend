package handler

import (
	"eps-backend/store"

	"gorm.io/gorm"
)

type Handler struct {
	kpiStore     store.KpiStore
	depositStore store.DepositStore
	salesStore   store.SalesStore
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		kpiStore:     store.NewKpiStore(db),
		depositStore: store.NewDepositStore(db),
		salesStore:   store.NewSalesStore(db),
	}
}
