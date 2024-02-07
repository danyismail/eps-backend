package handler

import (
	"eps-backend/store"

	"gorm.io/gorm"
)

type Handler struct {
	kpiStore     store.KpiStore
	depositStore store.DepositStore
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		kpiStore:     store.NewKpiStore(db),
		depositStore: store.NewDepositStore(db),
	}
}
