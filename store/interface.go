package store

import "eps-backend/model"

type KpiStore interface {
	FindAll(startDt, endDt string, page int, view int, mdn string, status int, shift string) (*[]model.VKpis, int64, int64, error)
	Test() (*[]model.VKpis, error)
}

type DepositStore interface {
	GetBalanceToday() ([]model.CurrentDeposit, error)
}

type SalesStore interface {
	GetSalesToday() ([]model.SalesReport, error)
}
