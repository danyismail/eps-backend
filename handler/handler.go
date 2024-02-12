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

func NewHandler(dev *gorm.DB, prod *gorm.DB) *Handler {
	return &Handler{
		kpiStore:     store.NewKpiStore(dev, prod),
		depositStore: store.NewDepositStore(dev, prod),
		salesStore:   store.NewSalesStore(dev, prod),
	}
}
